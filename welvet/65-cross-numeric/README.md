# 65. Cross-numeric train + down-the-dem

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/runtime/training`  
**Status:** ok — ✅

## When

Walking a Grid/Stack with the shared `runtime/training` path.

## Where

`import "github.com/openfluke/welvet/runtime/training"`

```bash
cd 65-cross-numeric && source ../env.sh && go run .
```

## Why

Weight storage dtype and activation Tensor[T] are independent axes. Proving train without a retained float32 master means sweeping W×A — not only matched float32 acts.

## What

W2A Step Cross-Numeric Train: polyops.AllKinds() × FormatNone weight dtype × Go Numeric act host (smoke ~735; full ~10.7k) via StepMesh, then assert no retained f32 master. Public Dense volumetric showcase: down-the-dem — dtype demotion ladder, packed quants, and the full 34×15×3 perm matrix with charts/PDF.

## Sample output (captured)

```
1.5 <nil> false
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/65-cross-numeric`).
