# 64. Scorecard → v1.0 / minors

**Part:** IX · Validate  
**Package:** `—`  
**Status:** ok — v1.1.0

## When

Orientation chapter: Scorecard → v1.0 / minors.

## Where

Book / repo map (no single import)

```bash
cd 64-scorecard && source ../env.sh && go run .
```

## Why

Version is earned from a weighted board, not marketing. v1.0 is 100/100 on the engine board. Minor tags (v1.1.0, …) pack features without a new board. Apps, stubs, and NPU sit off-board.

## What

version = 0.{round(earned)} until 100 → v1.0. Scorecard today 100/100; this book tags v1.1.0. Training credit is §9 on the board (8 pts), not an afterthought.

## Sample output (captured)

```
v1.1.0
scorecard 100/100
— 64-scorecard (122ms)

======== welvet book: ok=70 fail=0 skip=0 ========
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/64-scorecard`).
