# 28. layers/metacognition

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/metacognition`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/metacognition` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/metacognition"`

```bash
cd 28-metacognition && source ../env.sh && go run .
```

## Why

Observed layers can apply heuristic stability rules (gate/scale/reset) without dtype morph/QAT.

## What

Wraps Dense + DefaultStabilityRules(); Stats exposed. Full timed matrix + train grids.

## Sample output (captured)

```
16 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/28-metacognition`).
