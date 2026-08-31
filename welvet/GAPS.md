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
