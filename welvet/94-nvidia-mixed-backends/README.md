# 94. NVIDIA CUDA — mixed Go/C++/CUDA

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecpp + nativecuda + placement`  
**Status:** ok — Linux amd64 + NVIDIA runtime + normal cgo build

## When

Run one Stack with Go, welvet.cpp, and welvet.cuda layers in the same train step.

## Where

```sh
go run .
# do NOT use -tags welvet_cuda here — that disables C++
```

## Why

Prove heterogeneous handoff: Float32 activations forward, gradients back, CUDA SGD auto-syncs to Go.

## What

Thirteen Dense layers: Go 0–3, C++ 4–8, CUDA 9–12. Trains NormalBP five times and reports C++ ABI call counts plus CUDA router stats.

## Requirements

Linux amd64, normal cgo build (both `libwelvet.so` and `libwelvet_cuda.so`), NVIDIA runtime. Soft-SKIPs when C++ or CUDA is missing.

## Captured sample

```text
Go layers 0-3 -> C++ layers 4-8 -> CUDA layers 9-12
loss=… cpp calls=… cuda routing={…}
Float32 activations travel forward; gradients travel back; routed CUDA weights return automatically to Go.
```
