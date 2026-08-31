# rnn as a cam

**When:** Temporal cam

**Where:** `go run ./05_layers -layer rnn`

**Why:** Short seq twins.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  rnn  temporal cams  loss=0.0572  plast=[0 0]
```

## Why this output

Seq twin OK; loss slightly below the 0.0625 plateau.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
