# Feature guide — when / where / why

## BranchModes (per cam)

| Feature | When | Where | Why |
|---------|------|-------|-----|
| **NormalBP** | Default learning | Any cam | Chain-rule SGD; sharp credit |
| **Tween / TweenChain** | Hot LR / local updates | Mix with BP cams | Softer updates; often needs lower LR than seed SGD |
| **Freeze** | Keep a prior / eval cam | Idle hemisphere | Still contributes to Combine; no ΔW |
| **Shadow** | Distillation / teacher | Frozen teacher cam | Students get KD toward teacher post |
| **Adversarial** | Robustness / debate | One twin | Negated LR — maximizes local gap |
| **Memory** | Sparse plasticity | Surprise-gated cam | Updates only if `LastLoss ≥ SurpriseThresh` |
| **Inherit** | Slot filler | BranchModes list | Use parent mode |
| **Split / Alt / Mesh\*** | Credit experiments | Advanced | See Welvet train_mode docs; Mesh needs Grid |

**Stamp:** `para.SetBranchModes(ModeNormalBP, ModeFreeze)`

## Combine

| Mode | When | Why |
|------|------|-----|
| **concat** | Different widths / keep both views | No merge loss; wider head |
| **add / avg** | Same-width twins | Classic bicameral; avg = soft ensemble |
| **filter** | Learned MoE | Gate picks cams per sample |
| **max** | Winner-take-all features | Hard route; sparse expert feel |
| **sparsek** | Keep top-K strong cams | Ignore weak hemispheres |
| **disagree** | Explicit debate signal | `avg + β(cam0−cam1)` |

Equal widths required except **concat**.

## CamSync

| Knob | When | Why |
|------|------|-----|
| **Alpha** | Soft coupling | 1% gentle, 100% hard mean |
| **When** | AfterSample / Step / Pulse | How often minds meet |
| **Groups** | Pairwise / cliques | Don’t sync enemies |
| **Cross** | Glue across Stack layers | Same-shape stores only |
| **BranchAlpha** | One-way / asymmetric | Teacher writes 0; student pulls |

Frozen/Shadow cams are skipped as sync *targets* (and filtered from default groups’ trainable set).

## CamKit / schedules

| Feature | When | Why |
|---------|------|-----|
| **BranchLRs** | Per-cam step size | Mix algorithms *and* rates; `0` = soft freeze |
| **RotateSchedule** | Sleep cycles | Alternate who is plastic |
| **ShadowCoef** | KD strength | Scale student←teacher gap |
| **DNAReg >0** | Speciation | Push cams apart in weight space |
| **DNAReg <0** | Soft sync without CamSync | Attract toward mean |
| **SurpriseThresh** | Memory cams | Hippocampus-ish |
| **DreamBuffer / DreamPulse** | Offline consolidation | Replay after `Pulse` |
| **RefreshMetrics** | Debug / River | Cosines, plasticity, modes |

## Cross-modal

Use `HemispheresFrom` with different Ops (e.g. CNN2 ∥ Dense) **only with concat**
(or matching flattened widths). Sync only works on same-shape stores (`Proj` for CNN).

## Sample outputs

Each chapter folder README embeds **captured `go run` stdout** and a **why** table explaining the numbers:

- [`01_modes/README.md`](01_modes/README.md)
- [`02_combine/README.md`](02_combine/README.md)
- [`03_camsync/README.md`](03_camsync/README.md)
- [`04_kit/README.md`](04_kit/README.md)
- [`05_layers/README.md`](05_layers/README.md) (+ per-layer `*.md`)
- [`06_recipes/README.md`](06_recipes/README.md)

## Proofs

Chapter READMEs embed **live `PASS` output** from `go run` with assertions (not smoke).
See [`README.md`](README.md).
