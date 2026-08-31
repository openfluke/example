# 15. layers/layernorm

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/layernorm`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/layernorm` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/layernorm"`

```bash
cd 15-layernorm && source ../env.sh && go run .
```

## Why

Classic mean+var normalization with γ/β — still required for many HF architectures.

## What

WebGPU forward; backward host today. Same dtype×quant axes as other weighted layers.

## Sample output (captured)

```
8 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/15-layernorm`).
