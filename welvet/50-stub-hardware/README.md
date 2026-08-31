# 50. stub/hardware

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/hardware`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/hardware` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/hardware"`

```bash
cd 50-stub-hardware && source ../env.sh && go run .
```

## Why

Dispatchers and UIs need a portable host audit (OS/CPU/RAM/GPU).

## What

Audit() → SystemAudit with linux /proc or platform fallbacks.

## Sample output (captured)

```
{Model:Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz Logical:12 GOMAXPROCS:12}
OS: linux | CPU: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz (12) | RAM: 31.12 GB | GPU: Unknown GPU (0.0 GB)
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/50-stub-hardware`).
