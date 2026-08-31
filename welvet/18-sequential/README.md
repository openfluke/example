# 18. layers/sequential

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/sequential`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/sequential` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/sequential"`

```bash
cd 18-sequential && source ../env.sh && go run .
```

## Why

Some cells need an ordered Dense chain without burning grid hops.

## What

Dense→Dense compose, or mixed Ops via NewFromOps (Dense, SwiGLU, RMSNorm, LayerNorm). Parallel as a child is parallel.ResidualGraft / a Sandwich — Sequential cannot import parallel.

## Sample output (captured)

```
16 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/18-sequential`).
