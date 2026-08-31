# dense as a cam

**When:** Default cam host

**Where:** `go run ./05_layers -layer dense`

**Why:** Any BranchModes / CamSync experiment; fastest.

Setup in the smoke: identical twins + `CombineAvg` + `BranchModes=NormalBP∥Freeze` + CamSync α=1%.

## Sample output

```
  ok  dense  default host  loss=0.0623  plast=[0.0012478139251470566 0]
```

## Why this output

Default BranchModes host. Cam0 primary Dense store moves; cam1 Freeze → plast 0.

Cam1 plasticity is **0** on purpose (`ModeFreeze`). See [`README.md`](README.md) for the full matrix.
