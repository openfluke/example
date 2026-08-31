# 1. What Welvet is

**Part:** I · Orientation  
**Package:** `github.com/openfluke/welvet`  
**Status:** ok — ✅ engine

## When

You need `github.com/openfluke/welvet`.

## Where

`import "github.com/openfluke/welvet/github.com/openfluke/welvet"`

```bash
cd 01-welvet && source ../env.sh && go run .
```

## Why

Loom’s flat poly/ package hit import-cycle and honesty walls (QAT morph, silent fallbacks, god-layer). Welvet is the rewrite: one feature per folder, storage-truth dtypes/quants, Dense as the shared MatVec microkernel, tests only in w2a, apps only in apps/.

## What

An AI engine in Go: layers, 34 dtypes, 20 quant formats, and three backends (CPU tiled · Plan 9 SIMD · WebGPU). Version tracks a 100-point scorecard (today v1.1.0 · 100/100).

## Sample output (captured)

```
pre 4 post 4
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/01-welvet`).
