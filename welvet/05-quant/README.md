# 5. quant — 20 pack formats

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/quant`  
**Status:** ok — ✅

## When

You need the `quant` foundation package.

## Where

`import "github.com/openfluke/welvet/quant"`

```bash
cd 05-quant && source ../env.sh && go run .
```

## Why

Inference, storage, and train need classic Q-packs, k-quants, IQ, Ternary/Binary, and Affine without a separate QAT mode or retained f32 master. Format is storage truth.

## What

Pack / Unpack / MatVec / MatVecT for FormatNone…AffinePacked. Dense SIMD once-projects codes into Int8QS + scales (EnsureQ* / EnsureK/IQ/AffineSIMDCache) — no full-matrix F32 inflate for k/IQ/Affine. SGD on packed stores: short-lived unpack scratch → update → re-Pack; Packed stays truth.

## Sample output (captured)

```
rows 2 cols 4 y [0.6857143 -1.4901161e-08]
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/05-quant`).
