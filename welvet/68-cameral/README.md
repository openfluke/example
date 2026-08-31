# 68. Cameral sandwiches + AAI Lucy

**Part:** VII · Apps  
**Package:** `github.com/openfluke/welvet/layers/parallel`  
**Status:** ok — ✅ cameral

## When

Building or training a net that needs the `layers/parallel` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/parallel"`

```bash
cd 68-cameral && source ../env.sh && go run .
```

## Why

A single Dense chain cannot host two independent weight copies that share an input, merge, and optionally train under different updates. That is the cameral graph: hemispheres, not screen-space sprites and not a second hidden size.

## What

Hemispheres / Bicameral / Sandwich / Mix. Stem → Parallel mid → Dense head. SetBranchModes + TrainStackMSE. Lucy races in AAI (test41 / test48 / test50) import this API; measuring is lucy/; harness is not engine.

## Sample output (captured)

```
mix loss 0 err <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/68-cameral`).
