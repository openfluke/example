# swiglu as a cam

**When:** FFN twin

**Where:** `go run ./05_layers -layer swiglu`

**Why:** Width=InputDim; pair with MHA stacks.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  swiglu  FFN twins  loss=0.0625  plast=[0 0]
```

## Why this output

Width=InputDim FFN cams.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
