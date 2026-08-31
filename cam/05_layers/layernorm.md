# layernorm as a cam

**When:** Norm as cam (rare)

**Where:** `go run ./05_layers -layer layernorm`

**Why:** Same — prefer inside a deep hemisphere.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  layernorm  norm cam  loss=0.8454  plast=[0.02069561693168054 0]
```

## Why this output

Same story as RMS — affine params plastic on cam0 only.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
