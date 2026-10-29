# Coverage gaps (honest)

## Now covered (this pass)

| Gap | Chapter |
|-----|---------|
| cnn1 + cnn3 | `20-cnn` |
| convt2 + convt3 | `25-convt` |
| runtime/dispatch | `71-dispatch` |
| model/wav2vec2 API | `72-wav2vec2` |
| every layer Forward | `73-all-layers` |
| empty model/kokoro + apps/kokoro | `74-kokoro` (documents placeholders; points to mosstts/qwentts) |
| apps inventory | `75-apps-map` |

## Still soft / external

| Item | Why |
|------|-----|
| Full wav2vec2 ASR | Needs HF `model.safetensors` (`WAV2VEC2_DIR`) |
| WebGPU (`07`) | Needs adapter; soft-ok without GPU |
| HF load (`39`) | Fake/missing path soft path |
| Apps runtime (`octo`, `flux2`, …) | Separate modules + weights; map only in `75` |
| Cameral Freeze/CamSync proofs | `example/cam` (asserted), not book printouts |
| stub/* | Intentional stubs — print + partial |
| w2a / validation / scorecard | Separate harness modules — print pointers |

## Cam deep-dive

Everything ModeFreeze / Shadow / CamSync / DNAReg / Rotate / Disagree → [`../cam`](../cam/).

## Apple examples 82–88

Mounted forward, Mamba's 34 native parameter dtypes, the 29 training modes,
explicit weight return and ENTITY saving, JSON placement, and mixed
Go/C++/Metal routing are demonstrated locally. Apple GPU examples need a
compatible Apple silicon macOS/cgo runtime and installed Metal dylib. The
results chapter is a labelled supplied snapshot; live timings come from
`apple_stress_test`. Browser/WASM or ANE execution is not claimed by these
native Metal examples.

## NVIDIA examples 89–95

Same mounted-session tour for `nativecuda` + welvet.cuda **0.79** (Go Welvet
**v1.4.0**). Soft-SKIP without an NVIDIA GPU so book smoke stays green.
`EnableRouting` (93–94) is Linux amd64 only. Live matrices live in w2a
`go_v_cuda` and `welvet.cuda/examples/cuda_smoke` — chapter 95 only prints pointers.
Browser/WASM CUDA is not claimed.
