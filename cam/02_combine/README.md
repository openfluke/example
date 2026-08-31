# 02 — Combine modes

**When:** choosing how hemisphere posts become one tensor  
**Where:** `Config.Combine`  
**Why:** avg/add ensemble; max WTA; sparsek; disagree; filter MoE; concat widths

```bash
go run ./02_combine
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 02_combine — prove merge modes ===
  PASS  avg_width  outShape=[1 4] want [1 4]
  PASS  add_width  outShape=[1 4] want [1 4]
  PASS  concat_width  outShape=[1 8] want feat=8
  PASS  max_routes_to_cam0  combined should match cam0 when cam0≫cam1
  PASS  sparsek_keeps_strong  SparseK=1 should ≈ strongest cam (cam0)
  PASS  disagree_formula  y = mean + β(a-b) holds
  PASS  filter_gate_trains  MoE gate Δ=0.01582

all Combine proofs passed
```

## Why these PASS lines prove it

| Proof | Assertion | Why |
|-------|-----------|-----|
| **avg/add_width** | out feat = 4 | Same-width merge |
| **concat_width** | out feat = **8** | Unequal stack = 4+4 |
| **max_routes_to_cam0** | combined ≡ cam0 when cam0≫cam1 | Hard max routing |
| **sparsek_keeps_strong** | SparseK=1 ≈ strongest cam | Top-K by ‖out‖₂ |
| **disagree_formula** | `y = mean + β(a−b)` numerically | Debate combine is exact |
| **filter_gate_trains** | MoE gate weights move | Gate is live, not decoration |

