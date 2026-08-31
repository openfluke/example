# softmax as a cam

**When:** Distribution cam

**Where:** `go run ./05_layers -layer softmax`

**Why:** Avg/max of probability views.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  softmax  distribution cams  loss=0.0169  plast=[0 0]
```

## Why this output

Softmax already in (0,1) → easy MSE to 0.25. Softmax itself has no weights → plast 0 is expected.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
