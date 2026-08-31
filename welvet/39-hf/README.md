# 39. model/hf — snapshots

**Part:** VI · Model IO  
**Package:** `github.com/openfluke/welvet/model/hf`  
**Status:** ok — ✅

## When

Checkpoints, HF import, tokenize, sample, or decode (`model/hf`).

## Where

`import "github.com/openfluke/welvet/model/hf"`

```bash
cd 39-hf && source ../env.sh && go run .
```

## Why

Import starts with probing HF/MLX layouts before packing ENTITY.

## What

InspectSnapshot, DetectArchitecture, safetensors/MLX loaders, Qwen3.5 hybrid helpers.

## Sample output (captured)

```
hf: config.json: stat /path/to/hf-snapshot/config.json: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/39-hf`).
