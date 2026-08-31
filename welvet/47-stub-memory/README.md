# 47. stub/memory

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/memory`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/memory` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/memory"`

```bash
cd 47-stub-memory && source ../env.sh && go run .
```

## Why

HF→ENTITY and GPU upload need footprint accounting and optional history charts.

## What

FromGrid, Footprint, InitScavenger, ReleaseTransient; WELVET_MEMORY_HISTORY=1.

## Sample output (captured)

```
{HostWeightsMB:0 GPUWeightsMB:0 GPUKVMB:0}
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/47-stub-memory`).
