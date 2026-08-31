# 48. stub/donate

**Part:** VIII · Stubs  
**Package:** `github.com/openfluke/welvet/stub/donate`  
**Status:** partial — 🚧

## When

Scaffold / experimental stub `stub/donate` (may be partial).

## Where

`import "github.com/openfluke/welvet/stub/donate"`

```bash
cd 48-stub-donate && source ../env.sh && go run .
```

## Why

LAN donors should accept framed JSON jobs without embedding HTTP in the engine.

## What

u32-LE + JSON frames; ServeTCP/Dial; model_push vs local_lm. v0 workers echo.

## Sample output (captured)

```
default port 17001
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/48-stub-donate`).
