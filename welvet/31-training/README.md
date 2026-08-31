# 31. runtime/training

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/runtime/training`  
**Status:** ok — ✅

## When

Walking a Grid/Stack with the shared `runtime/training` path.

## Where

`import "github.com/openfluke/welvet/runtime/training"`

```bash
cd 31-training && source ../env.sh && go run .
```

## Why

Suites and small nets need MSE+SGD and tween hooks without inventing an external trainer or a retained float32 master beside storage.

## What

MSE/MSEGrad, SGD, Step, ApplyTween/StepTween, StepMesh. Layer-agnostic ApplyGradSGD dispatch (Dense…Mamba/GDN/…). FormatNone: in-dtype ApplySGD; packed: unpack→update→re-Pack. No QAT dual path — storage dtype/format is truth after every step. Sandwich credit (Split / FastProxy / Sparse / …) lives in layers/parallel TrainMode — see §67.

## Sample output (captured)

```
0 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/31-training`).
