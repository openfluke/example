# 56. stub/grouping

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/grouping`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/grouping` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/grouping"`

```bash
cd 56-stub-grouping && source ../env.sh && go run .
```

## Why

Detect layer archetypes from safetensor-style names before mounting.

## What

GroupRelatedTensors, DetectMHA/DetectSwiGLU/DetectRMSNorm → ArchetypeHint.

## Sample output (captured)

```
true {MultiHeadAttention 64 4}
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/56-stub-grouping`).
