# 89. NVIDIA CUDA — mount and forward

**Part:** XIII · NVIDIA acceleration (welvet.cuda)  
**Package:** `nativecuda`  
**Status:** ok — needs Linux/Windows amd64 + NVIDIA runtime

## When

Mount a Go layer and execute its forward pass on NVIDIA GPU.

## Where

From `example/welvet/89-nvidia-cuda`:

```sh
go run .
go run . -library /path/to/libwelvet_cuda.so
```

## Why

See the backend version and CUDA runtime identity; release the session explicitly.

## What

Creates an identity Dense layer, mounts it, forwards a Float32 tensor, and closes the session.

## Requirements

Linux x86_64 or Windows amd64, Go with cgo, NVIDIA driver/runtime, and the
installed independent CUDA shared library. Set `WELVET_CUDA_LIBRARY` or install
from the CUDA checkout:

```sh
cd ../welvet.cuda   # sibling of welvet
./scripts/linux_install_go.sh
# Windows: ./scripts/windows_install_go.ps1
```

Use the Go **v1.4.0** checkout containing `nativecuda`. Examples 89–93 can also
use `go run -tags welvet_cuda .` to exclude the independent C++ binding. Example
94 needs the normal build to keep both native backends. Unsupported platforms
print SKIP. Missing libraries / no GPU on a supported host also print SKIP.

These are usage examples with small shapes. For speed matrices, use w2a
`go_v_cuda` / the fat online suite. Results are numerical CPU/GPU comparisons,
not model prediction accuracy.

Backend labels in routed examples: `cpu-tiled` / `go_cpu` is Go CPU,
`native_cpu` is welvet.cpp, and `cuda` is the independent NVIDIA GPU backend.

## Captured sample

```text
backend=cuda version=0.79
device=cuda:…
input=[0.5 -0.25] output=[0.5 -0.25] shape=[1 2]
```
