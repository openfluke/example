# 93. NVIDIA CUDA — JSON backend placement

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecuda + runtime/placement`  
**Status:** ok — Linux amd64 + NVIDIA runtime (`EnableRouting`)

## When

Assign some Stack layers to CUDA via execution JSON while Go owns the schedule.

## Where

```sh
go run .
go run . -config /path/to/placement.json
```

## Why

Show that `nativecuda.EnableRouting` + `placement.Stack` drive native backward/SGD without an explicit `Mount` session.

## What

Embeds a three-layer plan (Go → CUDA → CUDA), trains NormalBP once, prints per-layer backends and router stats. Routed CUDA weight updates are already synced into Go stores.

## Requirements

**Linux amd64** (routing requirement), cgo, NVIDIA runtime, bundled `libwelvet_cuda.so`.
Do **not** use `-tags welvet_cuda` if you also need C++ in other chapters.
Windows mounts work for 89–92; automatic routing is Linux-only today.

## Captured sample

```text
layer=0 backend=cpu-tiled
layer=1 backend=cuda
layer=2 backend=cuda
loss=… cuda routing={Mounts:… Forwards:… Backwards:… Updates:… TrainingSteps:0}
Routed updates are already in Go stores; Close releases cached mounts.
```
