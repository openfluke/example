# 04 — CamKit / LRs / Rotate / DNA / Dream

**When:** schedules, plasticity, replay, speciation  
**Where:** `SetBranchLRs` / `SetRotateSchedule` / `SetCamKit` / `DreamPulse`  
**Why:** sleep cycles, soft freeze, DNA push/pull, offline consolidation

```bash
go run ./04_kit
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 04_kit — prove LRs / rotate / DNA / dream ===
  PASS  BranchLRs_1|0  cam0 Δ=0.009794 cam1 Δ=0
  PASS  Rotate_slot0  first 5: cam0 Δ=0.002567 cam1 Δ=0
  PASS  Rotate_slot1  next 5: cam0 Δ=0 cam1 Δ=0.002486
  PASS  DNAReg_diversify  diversify cos 0.8446 → -0.8649
  PASS  DNAReg_attract  attract cos 0.3717 → 0.9994
  PASS  DreamPulse  replay avgLoss=0.0012 moved cam0 Δ=0.0009936
  PASS  RefreshMetrics  modes=[NormalBP Freeze] plast=[0.000660944211525738 0]

all CamKit proofs passed
```

## Why these PASS lines prove it

| Proof | Assertion | Why |
|-------|-----------|-----|
| **BranchLRs 1\|0** | cam1 Δ=0 | LR×0 = soft freeze |
| **Rotate_slot0/1** | only cam0 then only cam1 moves | Sleep schedule flips plasticity |
| **DNAReg_diversify** | cos 0.84 → **negative** | Push away from mean |
| **DNAReg_attract** | cos 0.37 → **~1** | Pull toward mean |
| **DreamPulse** | replay loss>0 and ΔW>0 | Buffer replay moves weights |
| **RefreshMetrics** | modes show Freeze, plast[1]=0 | Meter matches reality |

