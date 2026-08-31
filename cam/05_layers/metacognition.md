# metacognition as a cam

**When:** Observer twin

**Where:** `go run ./05_layers -layer metacognition`

**Why:** Meta Dense observer.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  metacognition  observer twin  loss=0.0625  plast=[0 0]
```

## Why this output

Observer Dense twin smoke OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
