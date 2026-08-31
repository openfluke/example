# 19. layers/residual

**Part:** III · Layers  
**Package:** `github.com/openfluke/welvet/layers/residual`  
**Status:** ok — ✅

## When

Building or training a net that needs the `layers/residual` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/residual"`

```bash
cd 19-residual && source ../env.sh && go run .
```

## Why

Skip connections stabilize deep stacks: y = F(x) + x with correct skip grads.

## What

F is Dense Dim→Dim, or mixed Ops via NewFromOps (Dense, SwiGLU, RMSNorm, LayerNorm). Parallel as F is parallel.ResidualGraft (y = F(x)+x) — Residual cannot import parallel.

## Sample output (captured)

```
16 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/19-residual`).
