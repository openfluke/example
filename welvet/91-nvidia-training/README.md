# 91. NVIDIA CUDA — training modes

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecuda + parallel + architecture`  
**Status:** ok — needs Linux/Windows amd64 + NVIDIA runtime

## When

Train a mounted model using native CUDA credit and optimizer kernels.

## Where

```sh
go run .
go run . -mode StepBP -steps 8
go run . -mode MeshTweenChain
go run . -mode all
```

## Why

Keep one session alive for Step pipeline state; use an actual Grid for Mesh modes.

## What

Runs a two-layer Dense model in NormalBP by default. All 29 update modes are selectable. Mesh modes use a disabled ingress plus two active Grid cells and four clock ticks per call.

## Requirements

Same as [89-nvidia-cuda](../89-nvidia-cuda/README.md). Bundled engine **welvet.cuda 0.79** covers train modes through StepTweenSplitSparse.

## Captured sample

```text
mode=NormalBP schedule=Stack calls=8 loss=… weights=synced-to-Go
```
