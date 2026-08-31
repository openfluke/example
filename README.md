# openfluke/example

Runnable Welvet examples — **every feature** from the
[Welvet feature book](https://openfluke.github.io/welvet/), plus a deep **cameral** cookbook.

| Folder | What |
|--------|------|
| [`welvet/`](welvet/) | Book chapters 01–75: core → layers → runtime → systems → model → stubs — each with **When / Where / Why** + sample output |
| [`cam/`](cam/) | Parallel cameral deep-dive: Freeze, Shadow, CamSync, DNAReg, Rotate, Combine, every layer-as-cam — **asserted PASS/FAIL proofs** |

## Quick start

```bash
# Full Welvet book smoke (all mains)
cd welvet && source ./env.sh && go run ./cmd/runall

# One chapter
cd welvet/11-dense && source ../env.sh && go run .

# Cameral proofs
cd cam && go run ./cmd/runall
```

Local engine: `../welvet` (via `go.mod` replace). Docs source:
`../chaosglue/welvet/openfluke.github.io`.

## PDF book

Run every example, collect each README + live stdout, write one PDF:

```bash
cd /home/openfluke/git/example
pip install --user -r requirements-pdf.txt
python3 build_examples_pdf.py -o examples-book.pdf
# READMEs only (no go run):
python3 build_examples_pdf.py --skip-run -o examples-readmes.pdf
# one suite / filter:
python3 build_examples_pdf.py --suite welvet --only 20-cnn,71-dispatch
```

