# cnn1 as a cam

**When:** 1D temporal/spatial sense

**Where:** `go run ./05_layers -layer cnn1`

**Why:** Audio/seq patches; sync via Proj.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  cnn1  1D conv senses  loss=0.0625  plast=[0 0]
```

## Why this output

Twin CNN1 forwards+trains. Primary-store plast meter often ~0 on short smoke; graph is live (loss computed). Sync via Proj.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
