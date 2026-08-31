# 13. layers/swiglu — gated FFN

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/swiglu`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/swiglu` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/swiglu"`

```bash
cd 13-swiglu && source ../env.sh && go run .
```

## Why

Modern decoder FFNs are SiLU(gate)⊙up → down. Projections must share Dense’s quant/backend matrix.

## What

Gate/Up/Down Dense children; DefaultFFN(dModel); WebGPU SiLU⊙ fuse on forward.

## Sample output (captured)

```
64 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/13-swiglu`).
