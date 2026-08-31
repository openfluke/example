# 11. layers/dense — MatVec microkernel

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/dense`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/dense` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/dense"`

```bash
cd 11-dense && source ../env.sh && go run .
```

## Why

Most FLOPs are W@x. One Dense stack owns FormatNone×34 and all quants × three backends so every composite proj shares one correctness surface — including native in-dtype SGD.

## What

New / NewConfigured[T], Forward/Backward (dispatch on Exec.Backend), Place, ApplyGradSGD (→ weights.ApplySGD on the store). SIMD forward (v1.0.3): dtype switch by MatVec strategy — DotTile / DotI8 / lowp packed / expand-once→DotTile / WireF64+DotTileF64; fused Dot* for classic Q*, k/IQ, AffinePacked. Composites (MHA, SwiGLU, CNN im2col, RNN/LSTM/Mamba, residual·sequential·parallel) reuse Dense children via syncProjExec. BackwardSIMD still DecodeRow/saxpy (not the new expand wires).

## Sample output (captured)

```
8 32
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/11-dense`).
