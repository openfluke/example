# Cam examples — Welvet Parallel / cameral cookbook

Working Go examples for **every cameral feature** in Welvet’s `layers/parallel`,
plus **every layer type** that can sit inside a Parallel as a cam.

```bash
cd /home/openfluke/git/example/cam
go run ./cmd/runall          # smoke everything
go run ./01_modes            # one chapter
go run ./05_layers -layer cnn2
```

## Map

| Folder | What |
|--------|------|
| [`FEATURES.md`](FEATURES.md) | When / where / why for every feature |
| [`LAYERS.md`](LAYERS.md) | Which layer as a cam, and why |
| [`01_modes`](01_modes/) | BranchModes: BP, Tween, Freeze, Shadow, Adv, Memory… |
| [`02_combine`](02_combine/) | concat · add · avg · filter · max · sparsek · disagree |
| [`03_camsync`](03_camsync/) | α, groups, Cross, asymmetric `BranchAlpha` |
| [`04_kit`](04_kit/) | BranchLRs, Rotate, DNAReg, Dream, metrics |
| [`05_layers`](05_layers/) | Dense→GDN: each Op as twin cams |
| [`06_recipes`](06_recipes/) | End-to-end patterns (teacher, sleep, twin fight…) |

## Mental model (30 seconds)

```
input ──► cam0 ──┐
                 ├─ Combine ──► output
input ──► cam1 ──┘
                 ▲
         BranchModes / LRs / CamSync / CamKit
```

- **Forward** always runs all cams (unless you build a different graph).
- **Train** updates only cams whose mode/LR allow it (`Freeze`/`Shadow`/`LR=0` skip).
- **CamSync** blends weights between cams (optional, α / one-way).
- **Combine** decides how posts fuse (avg is the usual bicameral merge).

Parent train mode should be `NormalBP` (or Tween…) so `BranchModes` are honored.

## Proofs (not smoke)

Every example **asserts** the feature and exits 1 on failure:

```bash
go run ./cmd/runall    # must print: all cam examples ok
```

| Chapter | What a PASS means |
|---------|-------------------|
| `01_modes` | Freeze/Shadow leave cam1 weights unchanged; Memory wakes/sleeps; Adv fights |
| `02_combine` | concat widens; max/sparsek route; disagree formula; MoE gate trains |
| `03_camsync` | hard sync cos→1 from ~0.37; one-way keeps teacher; groups exclude cam2 |
| `04_kit` | Rotate swaps who learns; DNAReg ± moves cos; DreamPulse ΔW>0 |
| `05_layers` | Each Op learns (seeded) + Freeze holds |
| `06_recipes` | Composed recipes each assert their behavior |

Captured stdout + explanations live in each folder’s `README.md`.
