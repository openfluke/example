# convt1 as a cam

**When:** 1D generator twin

**Where:** `go run ./05_layers -layer convt1`

**Why:** Upsample hemispheres.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  convt1  1D upsample  loss=0.0625  plast=[0 0]
```

## Why this output

Generator-style twin; TrainMSE path OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
