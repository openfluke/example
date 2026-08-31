# 06 — Recipes

**When:** compose modes + sync + kit into behaviors  
**Where:** `go run ./06_recipes`  
**Why:** teacher, sleep, debate, memory, dream, concat

```bash
go run ./06_recipes
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 06_recipes — prove composed behaviors ===
  PASS  teacher_student  teacher frozen, student moved
  PASS  sleep_slot0  only cam0 Δ=0.00257/0
  PASS  sleep_slot1  only cam1 Δ=0/0.00249
  PASS  debate  both move; debate loss 0.0121 vs coop 0.0110
  PASS  surprise_memory  asleep Δ=0 awake Δ=0.00582
  PASS  dream_consolidate  avg=0.0012 weights moved on replay
  PASS  cross_modal_concat  feat=8 want 3+5=8
  PASS  cross_modal_trains  cam0 moved under concat

all recipe proofs passed
```

## Why these PASS lines prove it

| Proof | Assertion | Why |
|-------|-----------|-----|
| **teacher_student** | teacher frozen, student moved | Shadow + one-way sync |
| **sleep_slot0/1** | alternating single-cam ΔW | Rotate schedule |
| **debate** | both move; debate loss ≥ coop | Adv + Disagree |
| **surprise_memory** | asleep vs awake ΔW | SurpriseThresh gate |
| **dream_consolidate** | replay moves weights | DreamBuffer |
| **cross_modal_concat** | feat=8=3+5; cam0 trains | Unequal cams need concat |

