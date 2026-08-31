# 63. Validation report — full suite

**Part:** IX · Validate  
**Package:** `github.com/openfluke/w2a`  
**Status:** ok — ✅ 246k cells

## When

You need `github.com/openfluke/w2a`.

## Where

`import "github.com/openfluke/welvet/github.com/openfluke/w2a"`

```bash
cd 63-validation && source ../env.sh && go run .
```

## Why

Claims are cheap; a stamped matrix is not. This is the actual output of one full w2a [0] Run ALL so the book's ✅ marks are backed by numbers you can reproduce, not asserted.

## What

Every timed layer sweeps its dtype × format × backend matrix; every suite runs its case checks. v1.0 board: 246,032 matrix cells, FAIL 0, RESULT PASS. GAP cells are declared skips (not silent fails).

## Sample output (captured)

```
w2a [0] ALL: 246032 cells  FAIL 0  RESULT PASS
GAP = declared skip (GDN non-f32, AffinePacked, …) — not a fail
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/63-validation`).
