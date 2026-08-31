# 38. model/entity — .entity files

**Part:** VI · Model IO  
**Package:** `github.com/openfluke/welvet/model/entity`  
**Status:** ok — ✅

## When

Checkpoints, HF import, tokenize, sample, or decode (`model/entity`).

## Where

`import "github.com/openfluke/welvet/model/entity"`

```bash
cd 38-entity && source ../env.sh && go run .
```

## Why

HF safetensors are awkward for native topology + packed weights. ENTITY is the Welvet checkpoint.

## What

Open/Inspect/IsEntity, LoadBlob/LoadQuantBlob, PackFromHF/ImportFromHF, WriteTransformerFile, SerializeNetwork, WriteCameralFile / LoadCameral (sandwich Stack + TrainMode).

## Sample output (captured)

```
open model.entity: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/38-entity`).
