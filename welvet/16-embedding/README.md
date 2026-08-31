# 16. layers/embedding

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/embedding`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/embedding` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/embedding"`

```bash
cd 16-embedding && source ../env.sh && go run .
```

## Why

Token IDs must gather rows from a table — not a Dense MatVec — with scatter grads on backward.

## What

Config{VocabSize, EmbeddingDim, SeqLen}; table on weights.Store; host gather on SIMD/WebGPU today.

## Sample output (captured)

```
[1 4 16] <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/16-embedding`).
