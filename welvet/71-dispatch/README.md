# 71. runtime/dispatch — Op switchboard

**When:** walking a grid / Parallel / Sequential and you only have `any` cells.  
**Where:** `github.com/openfluke/welvet/runtime/dispatch`  
**Why:** one Forward/Backward/Pack/SetDType/ApplyGrad path for every Welvet Op; **unknown types hard-error** (never silently pretend they are Dense).

Do **not** import dispatch from `layers/parallel` (import cycle) — Parallel dispatches branches itself.

```bash
cd 71-dispatch && source ../env.sh && go run .
```
