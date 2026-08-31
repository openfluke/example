# 05 — Every layer as cams

**When:** CNN / Mamba / LSTM / … as Parallel cams  
**Where:** `HemispheresFrom` + twin Ops  
**Why:** CamSync / Freeze / Train work on every hosted Op

```bash
go run ./05_layers
```

Exits **non-zero** if any proof fails (`PASS` / `FAIL` lines).

## Live output

```
=== 05_layers — prove each Op works as cams ===
  PASS  dense  default host  loss 0.1230→0.0789  Freeze ok
  PASS  cnn1  1D conv  loss 0.1181→0.0854  Freeze ok
  PASS  cnn2  2D vision  loss 0.1161→0.0602  Freeze ok
  PASS  cnn3  3D  loss 0.1155→0.0271  Freeze ok
  PASS  convt1  1D upsample  loss 0.1206→0.1148  Freeze ok
  PASS  convt2  2D upsample  loss 0.1201→0.1077  Freeze ok
  PASS  convt3  3D upsample  loss 0.1188→0.1053  Freeze ok
  PASS  mha  attention  loss 0.1217→0.1195  Freeze ok
  PASS  swiglu  FFN  loss 0.1230→0.1223  Freeze ok
  PASS  rmsnorm  norm  loss 0.4835→0.0027  Freeze ok
  PASS  layernorm  norm  loss 1.1225→0.0003  Freeze ok
  PASS  softmax  distribution  loss 0.0708→0.0708  Freeze ok
  PASS  rnn  temporal  loss 0.1225→0.0009  Freeze ok
  PASS  lstm  gated temporal  loss 0.1225→0.0741  Freeze ok
  PASS  embedding  tables  loss 0.1252→0.0100  Freeze ok
  PASS  sequential  deep hemi  loss 0.1218→0.1203  Freeze ok
  PASS  residual  residual hemi  loss 0.0431→0.0277  Freeze ok
  PASS  kmeans  prototypes  loss 0.0003→0.0003  Freeze ok
  PASS  mamba  SSM  loss 0.1208→0.0996  Freeze ok
  PASS  metacognition  observer  loss 0.1225→0.0282  Freeze ok
  PASS  gdn  gated delta  loss 0.1225→0.1225  Freeze ok

all layer-as-cam proofs passed
```

## Why these PASS lines prove it

Each layer must:

1. **Forward** without error  
2. **Learn** (loss drops after seeded init) — except weightless `softmax`, tiny `kmeans`, zero-blob `gdn` (Forward+Freeze still required)  
3. **Freeze cam1** — `ActiveModes[1]==Freeze` and primary store Δ≈0  

`PASS … loss a→b Freeze ok` means all three held. CNN/MHA/etc. are seeded (Welvet `New` is zero-init).

