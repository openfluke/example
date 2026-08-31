# 46. stub/serialization

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/serialization`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/serialization` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/serialization"`

```bash
cd 46-stub-serialization && source ../env.sh && go run .
```

## Why

Volumetric grids need JSON/ENTITY persist beyond transformer packs.

## What

SerializeEntity/LoadEntity, SerializeGrid/GridFromSpec — native FormatNone bytes; packed via wire.

## Sample output (captured)

```
101 <nil>
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/46-stub-serialization`).
