# residual as a cam

**When:** Residual hemisphere

**Where:** `go run ./05_layers -layer residual`

**Why:** Skip-connected cam.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  residual  residual cam  loss=0.0625  plast=[0 0]
```

## Why this output

Skip-connected hemisphere twin.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
