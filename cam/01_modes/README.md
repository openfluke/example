# 01 — BranchModes

**When:** different update rules (or idle/teacher) per cam  
**Where:** `para.SetBranchModes(...)` with parent `NormalBP`  
**Why:** mix credit, freeze priors, distill, adversarial twins, surprise memory

```bash
go run ./01_modes
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 01_modes — prove BranchModes ===
  PASS  both_BP  both cams moved Δ=0.009232 / 0.009232
  PASS  Freeze  cam0 moved Δ=0.009794, frozen cam1 Δ=0 (want ~0)
  PASS  Shadow  teacher frozen Δ=0, student moved Δ=0.1343
  PASS  Adversarial  both move Δ=0.00832/0.00832; Adv loss 0.0013 ≥ BP loss 0.0007 (fight)
  PASS  Memory  asleep Δ=0 (~0), awake Δ=0.007141 (>0)
  PASS  Tween  Tween cam moved Δ=0.003871

all BranchMode proofs passed
```

## Why these PASS lines prove it

| Proof | Assertion | Why that shows the feature |
|-------|-----------|----------------------------|
| **both_BP** | both cams ΔW ≫ 0 | Shared learning works |
| **Freeze** | cam0 moves, cam1 Δ≈0 | Frozen cam still forwards; no weight update |
| **Shadow** | teacher Δ≈0, student moves | Shadow = frozen KD teacher |
| **Adversarial** | both move; Adv loss ≥ BP loss | Negated LR fights the objective |
| **Memory** | asleep Δ≈0, awake Δ≫0 | SurpriseThresh gates updates |
| **Tween** | Tween cam Δ≫0 | Alternate update family still applies |

Cams are **diverged** first so Freeze isn’t vacuous.

