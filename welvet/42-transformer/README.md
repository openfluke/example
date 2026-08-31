# 42. model/transformer — generate

**Part:** VI · Model IO  
**Package:** `github.com/openfluke/welvet/model/transformer`  
**Status:** ok — ✅

## When

Checkpoints, HF import, tokenize, sample, or decode (`model/transformer`).

## Where

`import "github.com/openfluke/welvet/model/transformer"`

```bash
cd 42-transformer && source ../env.sh && go run .
```

## Why

ENTITY packs must run as Llama-style decoders with KV cache, profiles (SIMD/WebGPU/fused), and chat templates.

## What

LoadEntity → Model; Generate; ApplyExec profiles; ExportFusedGPUSpec / hybrid sync.

## Sample output (captured)

```
open model.entity: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/42-transformer`).
