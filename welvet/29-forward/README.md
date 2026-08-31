# 29. runtime/forward

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/runtime/forward`  
**Status:** ok — ✅

## When

Walking a Grid/Stack with the shared `runtime/forward` path.

## Where

`import "github.com/openfluke/welvet/runtime/forward"`

```bash
cd 29-forward && source ../env.sh && go run .
```

## Why

A grid of heterogeneous ops needs one walker that dispatches by concrete type and fails loudly on unknowns.

## What

Forward[T](grid, input) → Result tape; Cell[T] for single-cell dispatch. Covers Dense…Residual + extended wired ops.

## Sample output (captured)

```
true <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/29-forward`).
