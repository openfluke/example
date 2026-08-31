# cnn2 as a cam

**When:** 2D vision cam

**Where:** `go run ./05_layers -layer cnn2`

**Why:** MNIST/CIFAR-style; put CNN *inside* Parallel for CamSync.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  cnn2  2D vision cams  loss=0.0625  plast=[0 0]
```

## Why this output

MNIST-style twin. Same plast-meter caveat as cnn1; use longer runs / Acc for learning proof.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
