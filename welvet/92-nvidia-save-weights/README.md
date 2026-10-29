# 92. NVIDIA CUDA — weight return and ENTITY save

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecuda + entity`  
**Status:** ok — needs Linux/Windows amd64 + NVIDIA runtime

## When

Prove explicit ownership: native train does not mutate Go stores until `SyncToGo`, then save/reload ENTITY.

## Where

```sh
go run .
go run . -out /tmp/trained.entity
```

## Why

CUDA sessions own copies; Go saving stays explicit.

## What

Mounts a Dense Stack, trains NormalBP five steps, asserts Go weights unchanged, syncs, writes Cameral ENTITY, remounts, and checks native snapshots match.

## Requirements

Same as [89-nvidia-cuda](../89-nvidia-cuda/README.md).

## Captured sample

```text
Go parameters unchanged before SyncToGo
weights returned to Go; saved=…/trained.entity; reload=exact
```
