# 61. stub/universal

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/universal`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/universal` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/universal"`

```bash
cd 61-stub-universal && source ../env.sh && go run .
```

## Why

Probe unknown safetensor geometry and mount placeholder grids until full weight import lands.

## What

LoadUniversal / LoadUniversalDetailed, ProbeDeepGeometry, MountGeometrically.

## Sample output (captured)

```
false universal: "/path/to/snapshot": open /path/to/snapshot: no such file or directory
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/61-stub-universal`).
