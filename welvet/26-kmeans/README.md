# 26. layers/kmeans

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/kmeans`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/kmeans` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/kmeans"`

```bash
cd 26-kmeans && source ../env.sh && go run .
```

## Why

Soft clustering as a differentiable layer lets topology experiments sit inside the same train loop.

## What

Centers on Dense (K×FeatureDim); soft assignment outputs. Full timed matrix + train grids.

## Sample output (captured)

```
[1 4] <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/26-kmeans`).
