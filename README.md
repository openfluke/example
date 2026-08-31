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

## When to open which

- **Learning a package** (Dense, MHA, Quant, DNA, …) → `welvet/NN-name/`
- **Multi-cam / BranchModes / CamSync / Freeze** → `cam/`
- **Why does this exist?** → each folder’s `README.md` (When / Where / Why / What)
