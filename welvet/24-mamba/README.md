# 24. layers/mamba — selective SSM

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/mamba`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/mamba` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/mamba"`

```bash
cd 24-mamba && source ../env.sh && go run .
```

## Why

SSM mixers (KindSSM) are not MHA clones — they need their own selective-scan path.

## What

InProj → softplus(Δ) scan → OutProj. Full timed matrix + train grids (scan ALU host).

## Sample output (captured)

```
true <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/24-mamba`).
