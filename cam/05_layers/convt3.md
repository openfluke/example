# convt3 as a cam

**When:** 3D generator twin

**Where:** `go run ./05_layers -layer convt3`

**Why:** Volumetric decode.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  convt3  3D upsample  loss=0.0625  plast=[0 0]
```

## Why this output

3D transpose twin OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
