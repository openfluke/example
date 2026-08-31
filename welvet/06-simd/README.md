# 6. simd — Plan 9 kernels

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/simd`  
**Status:** ok — ✅

## When

You need the `simd` foundation package.

## Where

`import "github.com/openfluke/welvet/simd"`

```bash
cd 06-simd && source ../env.sh && go run .
```

## Why

CPU peak needs hand-written AVX2/NEON without a silent Go fallback that pretends SIMD ran.

## What

amd64/arm64 .s kernels: DotTile, DotI8/U8, DotQ4_0, Saxpy, BitNet helpers, packed f16/bf16/fp8/fp4 dots; amd64 AVX2 DotTileF64 (WireF64); Go fused DotKRow / DotIQRow / DotAffineRow for k/IQ/Affine. SimdEnabled() false → BackendSIMD hard-errors.

## Sample output (captured)

```
10
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/06-simd`).
