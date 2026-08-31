# 3. core — types & backends

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/core`  
**Status:** ok — ✅

## When

You need the `core` foundation package.

## Where

`import "github.com/openfluke/welvet/core"`

```bash
cd 03-core && source ../env.sh && go run .
```

## Why

Every polymorphic path needs one place for DType, LayerType, Activation, Backend, Tensor[T], and slim Layer metadata — without QAT morph defaults.

## What

34 storage dtypes, Numeric generics, Tensor[T], ExecConfig (BackendCPUTiled | BackendSIMD | BackendWebGPU), Activate/ActivateDeriv, and converters for f16/bf16/fp8/fp4.

## Sample output (captured)

```
simd bfloat16 8
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/03-core`).
