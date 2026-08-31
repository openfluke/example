# 34. systems/evolution

**Part:** V · Systems  
**Package:** `github.com/openfluke/welvet/systems/evolution`  
**Status:** ok — ✅

## When

Adaptation, measurement, or evolution (`systems/evolution`).

## Where

`import "github.com/openfluke/welvet/systems/evolution"`

```bash
cd 34-evolution && source ../env.sh && go run .
```

## Why

Topology search and weight crossover need first-class splice + NEAT on CPU-resident grids.

## What

SpliceDNA, NEATMutate, NewNEATPopulation, CloneGrid — dtype/quant preserved via SetFromF32.

## Sample output (captured)

```
<nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/34-evolution`).
