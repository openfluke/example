# sequential as a cam

**When:** Deep hemisphere

**Where:** `go run ./05_layers -layer sequential`

**Why:** Whole MLP stack = one mind.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  sequential  deep hemisphere  loss=0.0625  plast=[0 0]
```

## Why this output

Whole MLP stack = one cam mind.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
