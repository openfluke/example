# 70. CamSync — inter-cameral / cross-mesh weight blend

**Part:** VII · Apps  
**Package:** `github.com/openfluke/welvet/layers/parallel`  
**Status:** ok — ✅ CamSync

## When

Building or training a net that needs the `layers/parallel` Op (also usable as a Parallel cam).

## Where

`import "github.com/openfluke/welvet/layers/parallel"`

```bash
cd 70-cam-sync && source ../env.sh && go run .
```

## Why

Mix BranchModes and different inits let cams diverge on purpose. Sometimes you want them to share — gently (1% pull) or hard (full average) — within a Parallel, across Stack children, or even between same-shaped stores sitting on wildly different mesh layouts (1×2×1 ↔ 2×4×5). That is CamSync: couple weights without collapsing the graph into one Dense.

## What

CamSyncConfig{Alpha, When, Groups, Cross}. BlendStores pulls every store in a clique toward the mean. Empty Groups = all cams; Cross wires SyncEndpoint pairs across layers/cams. Shape rule: Rows×Cols must match. Bidirectional today; one-way teacher→student not yet.

## Sample output (captured)

```
cosine before=-0.0697 after=1.0000
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/70-cam-sync`).
