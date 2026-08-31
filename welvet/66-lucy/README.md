# 66. lucy — SoftAcc / Score measuring

**Part:** V · Systems  
**Package:** `github.com/openfluke/welvet/lucy`  
**Status:** ok — ✅

## When

Adaptation, measurement, or evolution (`lucy`).

## Where

`import "github.com/openfluke/welvet/lucy"`

```bash
cd 66-lucy && source ../env.sh && go run .
```

## Why

Adaptation benches (test41-w, tide, live_gpt) need one shared measuring math — SoftAcc, Availability, AdaptPct, Score — not three copies of the formulas.

## What

Pure measuring package: SoftAcc / SoftAccProb, Window + Snapshot, Finalize. No datasets, no train loops. Sine scale 0.10; classification SoftAccProb scale 1.0. Density / synthetic-organism board is §69.

## Sample output (captured)

```
soft=20.0 class=91.0 avail=80.0 score=7680
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/66-lucy`).
