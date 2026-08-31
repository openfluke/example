# 58. stub/observer

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/observer`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/observer` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/observer"`

```bash
cd 58-stub-observer && source ../env.sh && go run .
```

## Why

Attach forward/backward observers for debugging without coupling to tanhi UDP.

## What

Observer interface; ConsoleObserver / HTTPObserver / BufferObserver; ComputeLayerStats.

## Sample output (captured)

```
*observer.ConsoleObserver
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/58-stub-observer`).
