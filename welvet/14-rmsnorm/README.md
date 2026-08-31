# 14. layers/rmsnorm

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/rmsnorm`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/rmsnorm` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/rmsnorm"`

```bash
cd 14-rmsnorm && source ../env.sh && go run .
```

## Why

Llama-style blocks normalize by RMS, not mean+var. Needs native fwd/bwd and WebGPU shaders.

## What

Per-token RMS + γ on weights.Store; WebGPU fwd+bwd; SIMD DotTile stats + host scale.

## Sample output (captured)

```
[0 0 0 0] <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/14-rmsnorm`).
