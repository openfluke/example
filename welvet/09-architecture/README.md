# 9. architecture — volumetric grid

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/architecture`  
**Status:** ok — ✅

## When

You need the `architecture` foundation package.

## Where

`import "github.com/openfluke/welvet/architecture"`

```bash
cd 09-architecture && source ../env.sh && go run .
```

## Why

Networks are spatial (Depth×Rows×Cols×LayersPerCell), not only linear stacks. Topology lives here; compute lives in layer packages.

## What

Grid/Cell/Coord, NewGrid, BindOp, HopOrder, SetRemoteLink/ResolveHop. VolumetricNetwork is an alias for Grid.

## Sample output (captured)

```
cells 2
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/09-architecture`).
