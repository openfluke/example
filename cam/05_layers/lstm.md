# lstm as a cam

**When:** Gated temporal cam

**Where:** `go run ./05_layers -layer lstm`

**Why:** Richer memory than RNN.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  lstm  gated temporal  loss=0.0620  plast=[0 0]
```

## Why this output

LSTM twin smoke OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
