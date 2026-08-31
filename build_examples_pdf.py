#!/usr/bin/env python3
"""Run Welvet/cam examples, collect READMEs + live stdout, write one PDF.

Usage:
  python3 build_examples_pdf.py                  # welvet + cam
  python3 build_examples_pdf.py --suite welvet
  python3 build_examples_pdf.py --suite cam
  python3 build_examples_pdf.py --skip-run        # READMEs only (cached outputs unused)
  python3 build_examples_pdf.py -o out/book.pdf

Requires: fpdf2  (pip install --user fpdf2)
"""

from __future__ import annotations

import argparse
import json
import os
import re
import subprocess
import sys
import time
from dataclasses import dataclass, field
from datetime import datetime, timezone
from pathlib import Path

try:
    from fpdf import FPDF
    from fpdf.enums import XPos, YPos
except ImportError:
    print("Missing fpdf2. Install with:  pip install --user fpdf2", file=sys.stderr)
    sys.exit(1)

ROOT = Path(__file__).resolve().parent
WELVET = ROOT / "welvet"
CAM = ROOT / "cam"

LIBERATION = Path("/usr/share/fonts/liberation-sans-fonts")
SOURCE_CODE = Path("/usr/share/fonts/adobe-source-code-pro-fonts")


@dataclass
class Chapter:
    suite: str
    slug: str
    num: str
    title: str
    dir: Path
    readme: str = ""
    output: str = ""
    ok: bool | None = None
    elapsed_s: float = 0.0
    error: str = ""


@dataclass
class Report:
    chapters: list[Chapter] = field(default_factory=list)
    started: str = ""
    finished: str = ""


def load_welvet_chapters() -> list[Chapter]:
    meta = json.loads((WELVET / "_chapters.json").read_text(encoding="utf-8"))
    out: list[Chapter] = []
    for row in meta:
        if not row.get("has_main"):
            continue
        slug = row["slug"]
        d = WELVET / slug
        if not (d / "main.go").exists():
            continue
        out.append(
            Chapter(
                suite="welvet",
                slug=slug,
                num=str(row.get("num", "")),
                title=row.get("title") or slug,
                dir=d,
            )
        )
    return out


def load_cam_chapters() -> list[Chapter]:
    pkgs = [
        ("01", "01_modes", "Train modes"),
        ("02", "02_combine", "Combine strategies"),
        ("03", "03_camsync", "CamSync"),
        ("04", "04_kit", "CamKit"),
        ("05", "05_layers", "Layers as cams"),
        ("06", "06_recipes", "Recipes"),
    ]
    out: list[Chapter] = []
    for num, slug, title in pkgs:
        d = CAM / slug
        if not (d / "main.go").exists():
            continue
        out.append(Chapter(suite="cam", slug=slug, num=num, title=title, dir=d))
    return out


def read_readme(ch: Chapter) -> str:
    p = ch.dir / "README.md"
    if p.exists():
        return p.read_text(encoding="utf-8", errors="replace")
    root = ROOT / ch.suite / "README.md"
    if ch.suite == "cam" and (CAM / "README.md").exists() and ch.slug == "01_modes":
        pass
    if p.exists():
        return p.read_text(encoding="utf-8", errors="replace")
    return "_(no README.md)_"


def go_env_for(suite_root: Path) -> dict[str, str]:
    env = os.environ.copy()
    cache = suite_root / ".cache"
    gocache = cache / "gocache"
    gotmp = cache / "gotmp"
    tmp = cache / "tmp"
    for d in (gocache, gotmp, tmp):
        d.mkdir(parents=True, exist_ok=True)
    env["GOCACHE"] = str(gocache)
    env["GOTMPDIR"] = str(gotmp)
    env["TMPDIR"] = str(tmp)
    return env


def run_chapter(ch: Chapter, timeout: float) -> None:
    env = go_env_for(ch.dir.parent if ch.suite == "welvet" else CAM)
    # welvet chapters: go run . in chapter dir
    # cam chapters: go run ./slug from cam root (matches cmd/runall)
    if ch.suite == "welvet":
        cmd = ["go", "run", "."]
        cwd = ch.dir
    else:
        cmd = ["go", "run", f"./{ch.slug}"]
        cwd = CAM
    t0 = time.perf_counter()
    try:
        proc = subprocess.run(
            cmd,
            cwd=cwd,
            env=env,
            capture_output=True,
            text=True,
            timeout=timeout,
        )
        ch.elapsed_s = time.perf_counter() - t0
        out = (proc.stdout or "") + (("\n" + proc.stderr) if proc.stderr else "")
        ch.output = out.strip()
        ch.ok = proc.returncode == 0
        if proc.returncode != 0:
            ch.error = f"exit {proc.returncode}"
    except subprocess.TimeoutExpired as e:
        ch.elapsed_s = time.perf_counter() - t0
        ch.ok = False
        ch.error = f"timeout after {timeout}s"
        out = ""
        if e.stdout:
            out += e.stdout if isinstance(e.stdout, str) else e.stdout.decode("utf-8", "replace")
        if e.stderr:
            out += "\n" + (e.stderr if isinstance(e.stderr, str) else e.stderr.decode("utf-8", "replace"))
        ch.output = out.strip()
    except Exception as e:  # noqa: BLE001 — surface any runner failure in PDF
        ch.elapsed_s = time.perf_counter() - t0
        ch.ok = False
        ch.error = str(e)
        ch.output = ""


def strip_md_noise(text: str) -> str:
    # keep content readable in PDF; light cleanup only
    text = text.replace("\r\n", "\n").replace("\r", "\n")
    text = text.replace("\u2014", "-").replace("\u2013", "-")
    text = text.replace("\u2018", "'").replace("\u2019", "'")
    text = text.replace("\u201c", '"').replace("\u201d", '"')
    text = text.replace("\u2192", "->").replace("\u00b7", "-")
    text = text.replace("\u2705", "[OK]").replace("\u274c", "[FAIL]")
    text = text.replace("\u2026", "...")
    text = text.replace("\u2299", "(o)").replace("\u22a5", "_|_").replace("\u2b1c", "[ ]")
    text = text.replace("\u226b", ">>")
    # drop remaining non-BMP / emoji that Liberation often lacks
    out = []
    for ch in text:
        o = ord(ch)
        if o < 32 and ch not in "\n\t":
            continue
        if o > 0xFFFF:
            out.append("?")
            continue
        out.append(ch)
    return "".join(out)


class BookPDF(FPDF):
    def __init__(self) -> None:
        super().__init__(format="A4")
        self.set_auto_page_break(auto=True, margin=18)
        self._setup_fonts()

    def _setup_fonts(self) -> None:
        reg = LIBERATION / "LiberationSans-Regular.ttf"
        bold = LIBERATION / "LiberationSans-Bold.ttf"
        mono = SOURCE_CODE / "SourceCodePro-Regular.otf"
        mono_b = SOURCE_CODE / "SourceCodePro-Bold.otf"
        if not reg.exists():
            raise SystemExit(f"Need Liberation Sans at {reg}")
        self.add_font("Body", "", str(reg))
        self.add_font("Body", "B", str(bold if bold.exists() else reg))
        if mono.exists():
            self.add_font("Mono", "", str(mono))
            self.add_font("Mono", "B", str(mono_b if mono_b.exists() else mono))
            self.mono_ok = True
        else:
            self.mono_ok = False

    def header(self) -> None:
        if self.page_no() == 1:
            return
        self.set_font("Body", "", 8)
        self.set_text_color(100, 100, 100)
        self.set_x(self.l_margin)
        self.cell(0, 6, "openfluke/example - Welvet + Cam", new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        self.set_draw_color(200, 200, 200)
        self.line(self.l_margin, self.get_y(), self.w - self.r_margin, self.get_y())
        self.ln(4)
        self.set_x(self.l_margin)
        self.set_text_color(0, 0, 0)

    def footer(self) -> None:
        self.set_y(-14)
        self.set_font("Body", "", 8)
        self.set_text_color(120, 120, 120)
        self.cell(0, 8, f"page {self.page_no()}", align="C")
        self.set_text_color(0, 0, 0)

    def h1(self, text: str) -> None:
        self.set_x(self.l_margin)
        self.set_font("Body", "B", 18)
        self.multi_cell(0, 9, text, new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        self.ln(2)

    def h2(self, text: str) -> None:
        self.set_x(self.l_margin)
        self.set_font("Body", "B", 14)
        self.multi_cell(0, 7, text, new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        self.ln(1)

    def h3(self, text: str) -> None:
        self.set_x(self.l_margin)
        self.set_font("Body", "B", 11)
        self.multi_cell(0, 6, text, new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        self.ln(1)

    def body(self, text: str) -> None:
        self.set_x(self.l_margin)
        self.set_font("Body", "", 10)
        self.multi_cell(0, 5, text, new_x=XPos.LMARGIN, new_y=YPos.NEXT)

    def mono_block(self, text: str, max_chars: int) -> None:
        if len(text) > max_chars:
            text = (
                text[:max_chars]
                + f"\n... truncated ({len(text)} chars total; raise --max-output)"
            )
        self.set_fill_color(245, 245, 245)
        self.set_font("Mono" if self.mono_ok else "Body", "", 7.5)
        x = self.l_margin
        w = self.epw
        for line in text.splitlines() or [""]:
            while line:
                max_c = 110
                if len(line) > max_c:
                    chunk, line = line[:max_c], line[max_c:]
                else:
                    chunk, line = line, ""
                if self.get_y() > self.h - 22:
                    self.add_page()
                self.set_x(x)
                self.cell(w, 4, chunk, fill=True, new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        self.ln(2)

    def ensure_space(self, h: float = 20) -> None:
        if self.get_y() > self.h - h:
            self.add_page()


def render_markdown(pdf: BookPDF, md: str) -> None:
    md = strip_md_noise(md)
    lines = md.split("\n")
    i = 0
    in_code = False
    code_buf: list[str] = []
    while i < len(lines):
        line = lines[i]
        if line.strip().startswith("```"):
            if in_code:
                pdf.mono_block("\n".join(code_buf), max_chars=50_000)
                code_buf = []
                in_code = False
            else:
                in_code = True
            i += 1
            continue
        if in_code:
            code_buf.append(line)
            i += 1
            continue

        if not line.strip():
            pdf.ln(2)
            i += 1
            continue

        # tables → plain lines
        if line.strip().startswith("|"):
            row = " | ".join(c.strip() for c in line.strip().strip("|").split("|"))
            if re.match(r"^[\s|:-]+$", line):
                i += 1
                continue
            pdf.set_x(pdf.l_margin)
            pdf.set_font("Mono" if pdf.mono_ok else "Body", "", 8)
            pdf.multi_cell(0, 4, row, new_x=XPos.LMARGIN, new_y=YPos.NEXT)
            i += 1
            continue

        if line.startswith("# "):
            pdf.h1(line[2:].strip())
        elif line.startswith("## "):
            pdf.h2(line[3:].strip())
        elif line.startswith("### "):
            pdf.h3(line[4:].strip())
        elif line.startswith("- ") or line.startswith("* "):
            pdf.set_x(pdf.l_margin)
            pdf.set_font("Body", "", 10)
            pdf.multi_cell(0, 5, "- " + _inline(line[2:].strip()), new_x=XPos.LMARGIN, new_y=YPos.NEXT)
        else:
            pdf.body(_inline(line))
        i += 1
    if in_code and code_buf:
        pdf.mono_block("\n".join(code_buf), max_chars=50_000)


def _inline(s: str) -> str:
    s = re.sub(r"\*\*(.+?)\*\*", r"\1", s)
    s = re.sub(r"`([^`]+)`", r"\1", s)
    s = re.sub(r"\[([^\]]+)\]\([^)]+\)", r"\1", s)
    return s


def build_pdf(report: Report, out: Path, max_output: int) -> None:
    pdf = BookPDF()
    pdf.add_page()
    pdf.h1("Welvet examples book")
    pdf.body("Generated from openfluke/example - every chapter README plus live go run output.")
    pdf.ln(3)
    pdf.body(f"Started:  {report.started}")
    pdf.body(f"Finished: {report.finished}")
    ok_n = sum(1 for c in report.chapters if c.ok is True)
    fail_n = sum(1 for c in report.chapters if c.ok is False)
    skip_n = sum(1 for c in report.chapters if c.ok is None)
    pdf.body(f"Chapters: {len(report.chapters)}  ok={ok_n}  fail={fail_n}  skip-run={skip_n}")
    pdf.ln(4)

    pdf.h2("Contents")
    for ch in report.chapters:
        status = {True: "PASS", False: "FAIL", None: "SKIP"}[ch.ok]
        pdf.set_x(pdf.l_margin)
        pdf.set_font("Body", "", 9)
        label = f"{ch.suite}/{ch.slug}  -  {ch.title}  [{status}]"
        pdf.multi_cell(0, 4.5, label, new_x=XPos.LMARGIN, new_y=YPos.NEXT)

    for ch in report.chapters:
        pdf.add_page()
        status = {True: "PASS", False: "FAIL", None: "SKIP"}[ch.ok]
        pdf.h1(f"{ch.num}. {ch.title}" if ch.num else ch.title)
        pdf.set_x(pdf.l_margin)
        pdf.set_font("Body", "", 9)
        pdf.set_text_color(80, 80, 80)
        pdf.multi_cell(
            0,
            5,
            f"suite={ch.suite}  slug={ch.slug}  status={status}  "
            f"time={ch.elapsed_s:.2f}s" + (f"  error={ch.error}" if ch.error else ""),
            new_x=XPos.LMARGIN,
            new_y=YPos.NEXT,
        )
        pdf.set_text_color(0, 0, 0)
        pdf.ln(2)

        pdf.h2("README")
        render_markdown(pdf, ch.readme or "_(empty)_")

        pdf.ln(3)
        pdf.h2("Live output")
        if ch.ok is None:
            pdf.body("(run skipped)")
        elif not ch.output:
            pdf.body("(no output)")
        else:
            pdf.mono_block(strip_md_noise(ch.output), max_chars=max_output)

    out.parent.mkdir(parents=True, exist_ok=True)
    pdf.output(str(out))


def main() -> int:
    ap = argparse.ArgumentParser(description="Run examples + READMEs → PDF")
    ap.add_argument(
        "--suite",
        choices=("all", "welvet", "cam"),
        default="all",
        help="which example suite(s) to include",
    )
    ap.add_argument("--skip-run", action="store_true", help="only embed READMEs (no go run)")
    ap.add_argument("-o", "--output", type=Path, default=ROOT / "examples-book.pdf")
    ap.add_argument("--timeout", type=float, default=180.0, help="per-chapter go run timeout (s)")
    ap.add_argument("--max-output", type=int, default=40_000, help="max chars of live output per chapter")
    ap.add_argument("--only", default="", help="comma-separated slug substrings to include")
    args = ap.parse_args()

    chapters: list[Chapter] = []
    if args.suite in ("all", "welvet"):
        chapters.extend(load_welvet_chapters())
    if args.suite in ("all", "cam"):
        chapters.extend(load_cam_chapters())

    if args.only:
        keys = [k.strip() for k in args.only.split(",") if k.strip()]
        chapters = [c for c in chapters if any(k in c.slug for k in keys)]

    if not chapters:
        print("No chapters matched.", file=sys.stderr)
        return 1

    report = Report(started=datetime.now(timezone.utc).strftime("%Y-%m-%d %H:%M:%SZ"))
    for i, ch in enumerate(chapters, 1):
        ch.readme = read_readme(ch)
        print(f"[{i}/{len(chapters)}] {ch.suite}/{ch.slug} …", flush=True)
        if args.skip_run:
            ch.ok = None
            ch.output = ""
        else:
            run_chapter(ch, timeout=args.timeout)
            mark = "ok" if ch.ok else "FAIL"
            print(f"    {mark} ({ch.elapsed_s:.1f}s)", flush=True)

    report.chapters = chapters
    report.finished = datetime.now(timezone.utc).strftime("%Y-%m-%d %H:%M:%SZ")
    build_pdf(report, args.output, max_output=args.max_output)
    print(f"\nWrote {args.output} ({args.output.stat().st_size} bytes)")
    fails = [c for c in chapters if c.ok is False]
    if fails:
        print(f"{len(fails)} chapter(s) failed — still included in PDF:", file=sys.stderr)
        for c in fails:
            print(f"  - {c.suite}/{c.slug}: {c.error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
