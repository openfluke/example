# 7. webgpu — device GEMV & shaders

**Part:** II · Foundation  
**Package:** `github.com/openfluke/welvet/webgpu`  
**Status:** ok — ✅

## When

You need the `webgpu` foundation package.

## Where

`import "github.com/openfluke/welvet/webgpu"`

```bash
cd 07-webgpu && source ../env.sh && go run .
```

## Why

GPU paths must bind a real adapter. Host “fake GPU” was banned so suites cannot stamp WebGPU done when ALU ran on CPU.

## What

DenseGEMV family (incl. quant/I8/resident), DenseGEMVT/DenseDW, RMSNorm, LayerNorm fwd, Softmax family, SwiGLUFuse. Available()/InitError() gate use.

## Sample output (captured)

```
webgpu probe goos=linux goarch=amd64 WELVET_WGPU_BACKEND=""
  adapter high-perf: name="NVIDIA GeForce GTX 1650 SUPER" vendor="NVIDIA" arch="" driver="610.57.04" backend=vulkan type=discrete-gpu | maxBuf=1099511627776 maxSSBO=2147483644 maxSSBO/stage=524288 maxBindGroups=8 maxComputeWG=65535 maxInvoc/WG=1024
  → selected high-perf
  adapter low-power: name="Intel(R) UHD Graphics 630 (CML GT2)" vendor="Intel open-source Mesa driver" arch="" driver="Mesa 26.1.8" backend=vulkan type=integrated-gpu | maxBuf=2147483647 maxSSBO=2147483644 maxSSBO/stage=103 maxBindGroups=8 maxComputeWG=65535 maxInvoc/WG=1024
  adapter force-fallback: name="llvmpipe (LLVM 22.1.8, 256 bits)" vendor="llvmpipe" arch="" driver="Mesa 26.1.8 (LLVM 22.1.8)" backend=vulkan type=cpu | maxBuf=2147483647 maxSSBO=134217728 maxSSBO/stage=48 maxBindGroups=8 maxComputeWG=65535 maxInvoc/WG=1024
  adapter default: name="NVIDIA GeForce GTX 1650 SUPER" vendor="NVIDIA" arch="" driver="610.57.04" backend=vulkan type=discrete-gpu | maxBuf=1099511627776 maxSSBO=2147483644 maxSSBO/stage=524288 maxBindGroups=8 maxComputeWG=65535 maxInvoc/WG=1024
  try RequestDevice inflated: maxBuf=2147483648 maxSSBO=1073741824 maxSSBO/stage=524288 maxBindGroups=8 maxComputeWG=65535 maxInvoc/WG=1024
RESULT: OK via=inflated adapter="NVIDIA GeForce GTX 1650 SUPER" backend=vulkan maxSSBO=2147483644 maxBuf=1099511627776
[1.5 2.5] <nil> NVIDIA GeForce GTX 1650 SUPER
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/07-webgpu`).
