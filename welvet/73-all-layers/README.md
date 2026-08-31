# 73. All layers — one forward each

**When:** you want a single command that proves every `layers/*` Op (plus seqmix + dispatch) actually forwards.  
**Where:** this folder imports every concrete layer package.  
**Why:** chapters 11–28 split by topic; this is the regression checklist for “does it run?”

Cameral proofs (Freeze / CamSync on each Op) live in [`../../cam/05_layers`](../../cam/05_layers/).

```bash
cd 73-all-layers && source ../env.sh && go run .
```

Each line prints **when / where / why** for that Op.
