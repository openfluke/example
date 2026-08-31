# 17. layers/softmax

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/softmax`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/softmax` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/softmax"`

```bash
cd 17-softmax && source ../env.sh && go run .
```

## Why

Classification heads and attention need stable softmax variants, including sparse/Gumbel/Entmax for research paths.

## What

Weightless layer; KindStandard/Temperature/Grid/Hierarchical/Gumbel/Masked/Sparse/… WebGPU covers std family; exotic kinds hard-error on GPU (no silent host).

## Sample output (captured)

```
[0.6380664 0.23473153 0.09543471 0.031767454] <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/17-softmax`).
