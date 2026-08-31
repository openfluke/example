# 4. weights — FormatNone MatVec

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/weights`  
**Status:** ok — ✅

## When

You need the `weights` foundation package.

## Where

`import "github.com/openfluke/welvet/weights"`

```bash
cd 04-weights && source ../env.sh && go run .
```

## Why

Unquantized matrices still need a typed store that streams MatVec and SGD without forcing a float32 master or Morph-as-training.

## What

Store holds DType + Format + Native/Packed bytes. New[T], MatVec, MatVecT, DecodeRow, SelectWire (F32/F64/I8), ApplySGD. FormatNone: update in native lanes (float32 payload is the only f32 buffer; float64 uses native ALU; other dtypes decode→update→re-encode). Packed: unpack→update→re-Pack then drop scratch. RetainsF32Master() is true only for FormatNone+float32. Dense and composite projs share this store.

## Sample output (captured)

```
[3 4]
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/04-weights`).
