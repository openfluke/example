# mamba as a cam

**When:** SSM twin

**Where:** `go run ./05_layers -layer mamba`

**Why:** Long-seq state-space cams.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  mamba  SSM twins  loss=0.0625  plast=[0 0]
```

## Why this output

State-space twin smoke OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
