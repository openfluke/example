# mha as a cam

**When:** Attention hemisphere

**Where:** `go run ./05_layers -layer mha`

**Why:** Same DModel; seq cams.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  mha  attention hemispheres  loss=0.0625  plast=[0 0]
```

## Why this output

Same DModel twins; seq input. Short smoke may not move metered store.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
