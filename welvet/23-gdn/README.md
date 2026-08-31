# 23. layers/gdn — gated delta net

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/gdn`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/gdn` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/gdn"`

```bash
cd 23-gdn && source ../env.sh && go run .
```

## Why

Linear attention / decode-first mixers (Gated DeltaNet) need a first-class package under KindLinearAttn.

## What

Exec CPU/SIMD/WebGPU; ForwardDecode; truncated BPTT. Timed matrix is Float32-primary: PermutationOK is f32 × FormatNone/BinaryPacked × CPU/SIMD/WebGPU; other cells GAP (declared), not FAIL.

## Sample output (captured)

```
<nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/23-gdn`).
