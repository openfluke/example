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

## WASM — npm + HTML (every chapter)

Each `welvet/<slug>/` and `cam/<slug>/` has **`npm/`** (Node) and **`html/`** (browser):

```bash
export WELVET_TS=/path/to/welvet/apps/w2a/typescript   # after npm run build:all
node welvet/11-dense/npm/run.mjs
node cam/01_modes/npm/run.mjs

# all chapters (73 OK + 8 native-only SKIP)
bash run-all-npm.sh
# or: bash run-all-wasm.sh

# HTML: sync assets, serve example root
bash _wasm/sync-assets.sh
npx --yes serve -l 4173 .
# → http://localhost:4173/welvet/11-dense/html/
```

Details: [`_wasm/README.md`](_wasm/README.md).

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

