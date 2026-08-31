# 10. fusedgpu — decoder on device

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/fusedgpu`  
**Status:** ok — ✅

## When

You need the `fusedgpu` foundation package.

## Where

`import "github.com/openfluke/welvet/fusedgpu"`

```bash
cd 10-fusedgpu && source ../env.sh && go run .
```

## Why

Token-by-token host round-trips kill decode. A fused engine keeps weights and scratch resident for Q4_0 and BinaryG128 hybrid paths.

## What

Engine from Spec (AppendTokens/Reset/Close); HybridEngine (PrefillSample/DecodeSample/DecodeChunk). Specs come from model/transformer export.

## Sample output (captured)

```
need an ENTITY file: open model.entity: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/10-fusedgpu`).
