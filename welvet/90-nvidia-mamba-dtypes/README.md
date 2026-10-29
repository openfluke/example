# 90. NVIDIA CUDA — Mamba parameter dtypes

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecuda + mamba`  
**Status:** ok — needs Linux/Windows amd64 + NVIDIA runtime

## When

Forward Mamba with Float32 activations while parameters stay FormatNone in each
of Welvet's 34 native dtypes.

## Where

```sh
go run .
go run . -dtype float32
go run . -dtype float16
```

## Why

Confirm the CUDA mount accepts the same dtype matrix the engine board claims.

## What

Mounts a small Mamba, forwards once per selected dtype, prints output shape and a checksum.

## Requirements

Same as [89-nvidia-cuda](../89-nvidia-cuda/README.md). Use Go **v1.4.0** +
bundled **welvet.cuda 0.79**.

## Captured sample

```text
float32    shape=[1 3 8] checksum=…
…
Mamba dtypes=34; activations=float32; parameters=FormatNone native dtype
```
