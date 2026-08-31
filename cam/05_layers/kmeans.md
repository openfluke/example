# kmeans as a cam

**When:** Prototype cam

**Where:** `go run ./05_layers -layer kmeans`

**Why:** Cluster / codebook twins.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  kmeans  prototype cams  loss=0.0069  plast=[0 0]
```

## Why this output

Lowest loss — cluster probs sit near constant target quickly.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
