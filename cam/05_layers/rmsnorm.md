# rmsnorm as a cam

**When:** Norm as cam (rare)

**Where:** `go run ./05_layers -layer rmsnorm`

**Why:** Usually nest inside Sequential cam.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  rmsnorm  norm cam  loss=0.5227  plast=[0.016977749851921242 0]
```

## Why this output

Higher MSE: norm posts ≠ flat 0.25 target. Cam0 γ moves (plast>0); Freeze cam1 idle.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
