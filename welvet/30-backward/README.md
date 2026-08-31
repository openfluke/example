# 30. runtime/backward

**Part:** IV · Runtime  
**Package:** `github.com/openfluke/welvet/runtime/backward`  
**Status:** ok — ✅

## When

Walking a Grid/Stack with the shared `runtime/backward` path.

## Where

`import "github.com/openfluke/welvet/runtime/backward"`

```bash
cd 30-backward && source ../env.sh && go run .
```

## Why

Training needs a reverse tape over the same ops forward used — no separate graph framework.

## What

Backward[T](fwdResult, gradOut) walks the tape and returns per-op weight grads.

## Sample output (captured)

```
true <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/30-backward`).
