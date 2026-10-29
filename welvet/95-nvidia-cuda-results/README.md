# 95. NVIDIA CUDA — results and suite pointers

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** w2a / cuda_smoke  
**Status:** ok — host-only; no GPU required

## When

Find where live CUDA timing/parity matrices live after the small usage chapters.

## Where

```sh
go run .
../run_nvidia.sh
```

## Why

Separate labelled pointers from live GPU measurement (same idea as Apple chapter 88).

## What

Prints Go **v1.4.0** / welvet.cuda **0.79** pins and commands for w2a `go_v_cuda`
and `welvet.cuda/examples/cuda_smoke`. Does not execute CUDA.

## Captured sample

```text
NVIDIA CUDA results pointers (not a live benchmark)

Engine pin: welvet.cuda 0.79 · model ABI 1 · Go binding: nativecuda (Welvet v1.4.0)
…
```
