# 20. CNN family — cnn1 / cnn2 / cnn3

**When:** spatial / temporal convolution senses (1D seq, 2D images, 3D volumes).  
**Where:** `layers/cnn1`, `layers/cnn2`, `layers/cnn3`  
**Why:** im2col → Dense `Proj` so quant/SIMD/WebGPU and CamSync (via Proj) are shared.

```bash
cd 20-cnn && source ../env.sh && go run .
```

| Op | Use |
|----|-----|
| **cnn1** | Audio / 1D sequences |
| **cnn2** | Images / MNIST-style |
| **cnn3** | Volumetric / video voxels |

Put CNN **inside** Parallel cams when you want CamSync on kernels.
