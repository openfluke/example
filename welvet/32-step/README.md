# 32. runtime/step — step mesh

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/runtime/step`  
**Status:** ok — ✅

## When

Walking a Grid/Stack with the shared `runtime/step` path.

## Where

`import "github.com/openfluke/welvet/runtime/step"`

```bash
cd 32-step && source ../env.sh && go run .
```

## Why

Spatial feedback (remote links) needs a discrete-time mesh where every cell updates from a double buffer — different from a decoder wavefront. Cross-numeric train also needs the same mesh with weight DType ⊥ activation Tensor[T].

## What

State[T], StepForward/StepBackward/StepApplyTween / StepMesh across the grid for all wired Ops × dtype × quant × CPU/SIMD. W2A Cross-Numeric Train: polyops.AllKinds() × weight dtype × act host (smoke ~21×7×5 ≈ 735; full ~21×34×15 ≈ 10.7k) — asserts no retained f32 master after StepMesh. Stack Step* is a different clock: IsLineStep / TrainLine (1D pipe) — see §67.

## Sample output (captured)

```
<nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/32-step`).
