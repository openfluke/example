# 8. tiling — SC/MC & workgroups

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/tiling`  
**Status:** ok — ✅

## When

You need the `tiling` foundation package.

## Where

`import "github.com/openfluke/welvet/tiling"`

```bash
cd 08-tiling && source ../env.sh && go run .
```

## Why

MatVec throughput depends on tile size and when to go multi-core vs GPU workgroups. Centralizing caps keeps Dense and friends consistent.

## What

DefaultCPUTile, DefaultGPUWG, CPUTile, PreferMultiCore, GPUWorkgroupsX.

## Sample output (captured)

```
32 true 16
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/08-tiling`).
