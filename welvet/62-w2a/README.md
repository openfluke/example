# 62. w2a — validation harness

**Part:** IX · Validate  
**Package:** `github.com/openfluke/w2a`  
**Status:** ok — ✅ harness

## When

You need `github.com/openfluke/w2a`.

## Where

`import "github.com/openfluke/welvet/github.com/openfluke/w2a"`

```bash
cd 62-w2a && source ../env.sh && go run .
```

## Why

Engine packages must stay free of tests. w2a owns timed 34×20×3 matrices, gap census, honesty stamps, and the train-mode permutation smoke (Test49). See §63 for a live full-suite run.

## What

Interactive go run . ([0] Run ALL). Suites under suites/*. StampBackendNote / AffinePackable prevent fake ✅. Test49: AllNamedTrainModes × 1³/2³/3³ × Parallel/Bicameral/poly, origin-only — included in [0].

## Sample output (captured)

```
w2a is a separate module — engine packages never contain tests.

  cd w2a
  go run .                                      # interactive; [0] = ALL
  go test ./tests/dense -v
  go test ./tests/parallel -run Test49AllTrainModesCubes -count=1 -v
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/62-w2a`).
