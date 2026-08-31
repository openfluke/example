# convt2 as a cam

**When:** 2D generator twin

**Where:** `go run ./05_layers -layer convt2`

**Why:** Decoder cams.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  convt2  2D upsample  loss=0.0625  plast=[0 0]
```

## Why this output

2D transpose twin OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
