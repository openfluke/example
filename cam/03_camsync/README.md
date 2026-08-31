# 03 — CamSync

**When:** cams should share / pull weight DNA  
**Where:** `SetCamSync` / `SyncNow` / Cross  
**Why:** soft↔hard consensus, one-way teacher, groups, cross-layer

```bash
go run ./03_camsync
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 03_camsync — prove sync pulls cos up ===
  PASS  α=1_hard_sync  cos 0.3769 → 1.0000 (hard sync should ≈1 and beat pre)
  PASS  α=0.25_soft_pull  cos 0.3717→0.6413, still distinct ΔW=0.75
  PASS  BranchAlpha_one_way  cam1 unchanged Δ=0; cos 0.3717→0.6635 (cam0 pulled to mean)
  PASS  Groups_exclude_cam2  cam2 frozen Δ=0; cam0≡cam1 Δ=0; cam0≠cam2 Δ=3 (was 3)
  PASS  Cross_pair  Cross hard-sync cos 0.0000→1.0000
  PASS  Cross_stack_trains  TrainStackMSE err=<nil>

all CamSync proofs passed
```

## Why these PASS lines prove it

| Proof | Assertion | Why |
|-------|-----------|-----|
| **α=1_hard_sync** | cos 0.37 → **1.0** | Diverged cams hard-average to identical |
| **α=0.25_soft_pull** | cos rises, weights still distinct | Soft pull ≠ clone |
| **BranchAlpha_one_way** | cam1 Δ=0, cos rises | Teacher not overwritten; student pulls to mean |
| **Groups_exclude_cam2** | cam0≡cam1, cam2 frozen & different | Clique sync only |
| **Cross_pair** | Cross hard-sync cos → 1 | Stack endpoint glue works |
| **Cross_stack_trains** | TrainStackMSE ok | Training still works with Cross cfg |

