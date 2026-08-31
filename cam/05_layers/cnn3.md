# cnn3 as a cam

**When:** 3D / volumetric

**Where:** `go run ./05_layers -layer cnn3`

**Why:** Medical/video volumes as twin senses.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  cnn3  3D cams  loss=0.0625  plast=[0 0]
```

## Why this output

Volumetric twin smoke OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
