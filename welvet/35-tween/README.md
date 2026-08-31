# 35. systems/tween

**Part:** V · Systems  
**Package:** `github.com/openfluke/welvet/systems/tween`  
**Status:** ok — ✅

## When

Adaptation, measurement, or evolution (`systems/tween`).

## Where

`import "github.com/openfluke/welvet/systems/tween"`

```bash
cd 35-tween && source ../env.sh && go run .
```

## Why

Target propagation (chain-rule or Hebbian layerwise gaps) is an alternative credit-assignment path. Not the same package as TrainMode Tween / TweenChain on a Sandwich (those are layers/parallel — §67).

## What

NewState, Forward, BackwardChainRule / BackwardLayerwise, ApplyGaps; SIMD DotTile/Saxpy budgets.

## Sample output (captured)

```
forward: no op at {0 0 0 0} (type Dense)
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/35-tween`).
