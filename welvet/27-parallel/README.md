# 27. layers/parallel — MoE + cameral

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/parallel`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/parallel` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/parallel"`

```bash
cd 27-parallel && source ../env.sh && go run .
```

## Why

Mixture-of-experts and multi-path cells need concat/add/avg/filter combines. Cameral graphs need sibling hemispheres that share input, merge outputs, and optionally train under distinct modes.

## What

Parallel branches + Stack sandwiches. Hemispheres / Bicameral / Sandwich build nested multi-cameral nets. SetBranchModes + TrainStackMSE let each hemi use a different TrainMode (all 29 named updates — BP, Tween, Split, FastProxy, Sparse, Step*, Mesh*, …).

## Sample output (captured)

```
loss 0 err <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/27-parallel`).
