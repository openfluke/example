# embedding as a cam

**When:** Dual embedding tables

**Where:** `go run ./05_layers -layer embedding`

**Why:** Two vocab views; Freeze one as prior.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  embedding  dual tables  loss=0.0619  plast=[0.00043333118975077967 0]
```

## Why this output

Token IDs in; cam0 embedding rows move (plast>0), Freeze cam1 idle.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
