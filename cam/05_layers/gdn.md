# gdn as a cam

**When:** Gated Δ-net twin

**Where:** `go run ./05_layers -layer gdn`

**Why:** Modern seq cam alternative to Mamba.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  gdn  gated delta twins  loss=0.0625  plast=[0 0]
```

## Why this output

Gated Δ-net twin smoke OK.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
