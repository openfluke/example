# Welvet feature examples

Runnable Go for **every Welvet book chapter** — copied from
`openfluke.github.io/welvet/examples` (and local NativeCPU chapters), pointed at local
`../../welvet` via `go.mod` replace.

Each folder has **When / Where / Why / What** plus captured sample output.

**v1.1.7 NativeCPU:** chapters **76–78** — opt-in `BackendNativeCPU`, per-layer `TrainNativeMSE`, TrainMode stack/parallel/mesh/line + `SyncWeightsFromNative`.

**v1.2.0 Hexagon NPU:** chapters **79–81** — opt-in `BackendQNN` + `SyncToQNN`, session layer tour, RESULTS.md showcase (`welvet.qnn` 1.0.0 · Windows ARM64). Put `welvet/native/qnn/windows_arm64` on `PATH` so `libwelvet_qnn.dll` loads.

**v1.3.0 Apple Metal:** chapters **82–88** — mounted forward, Mamba native dtypes,
all 29 training modes, explicit weight return/ENTITY saving, JSON placement and
mixed Go/C++/Metal execution. The native backend is independent welvet.metal 0.68.

```sh
./run_apple.sh                       # small Apple usage tour
cd 82-apple-metal && go run .        # first Metal forward
```

Cameral deep-dive (Freeze, CamSync, DNAReg, …) lives next door: [`../cam`](../cam/).

```bash
cd welvet   # or /home/openfluke/gs/example/welvet
source ./env.sh
go run ./cmd/runall          # smoke all mains that exist
cd 11-dense && go run .      # one chapter
```

## Index

| # | Folder | Title | Package | Status | main.go |
|---|--------|-------|---------|--------|---------|
| 1 | [`01-welvet`](01-welvet/) | What Welvet is | `github.com/openfluke/welvet` | ok | yes |
| 2 | [`02-tree`](02-tree/) | Repository map | `—` | ok | yes |
| 3 | [`03-core`](03-core/) | core — types & backends | `github.com/openfluke/welvet/core` | ok | yes |
| 4 | [`04-weights`](04-weights/) | weights — FormatNone MatVec | `github.com/openfluke/welvet/weights` | ok | yes |
| 5 | [`05-quant`](05-quant/) | quant — 20 pack formats | `github.com/openfluke/welvet/quant` | ok | yes |
| 6 | [`06-simd`](06-simd/) | simd — Plan 9 kernels | `github.com/openfluke/welvet/simd` | ok | yes |
| 7 | [`07-webgpu`](07-webgpu/) | webgpu — device GEMV & shaders | `github.com/openfluke/welvet/webgpu` | ok | yes |
| 8 | [`08-tiling`](08-tiling/) | tiling — SC/MC & workgroups | `github.com/openfluke/welvet/tiling` | ok | yes |
| 9 | [`09-architecture`](09-architecture/) | architecture — volumetric grid | `github.com/openfluke/welvet/architecture` | ok | yes |
| 10 | [`10-fusedgpu`](10-fusedgpu/) | fusedgpu — decoder on device | `github.com/openfluke/welvet/fusedgpu` | ok | yes |
| 11 | [`11-dense`](11-dense/) | layers/dense — MatVec microkernel | `github.com/openfluke/welvet/layers/dense` | ok | yes |
| 12 | [`12-mha`](12-mha/) | layers/mha — attention | `github.com/openfluke/welvet/layers/mha` | ok | yes |
| 13 | [`13-swiglu`](13-swiglu/) | layers/swiglu — gated FFN | `github.com/openfluke/welvet/layers/swiglu` | ok | yes |
| 14 | [`14-rmsnorm`](14-rmsnorm/) | layers/rmsnorm | `github.com/openfluke/welvet/layers/rmsnorm` | ok | yes |
| 15 | [`15-layernorm`](15-layernorm/) | layers/layernorm | `github.com/openfluke/welvet/layers/layernorm` | ok | yes |
| 16 | [`16-embedding`](16-embedding/) | layers/embedding | `github.com/openfluke/welvet/layers/embedding` | ok | yes |
| 17 | [`17-softmax`](17-softmax/) | layers/softmax | `github.com/openfluke/welvet/layers/softmax` | ok | yes |
| 18 | [`18-sequential`](18-sequential/) | layers/sequential | `github.com/openfluke/welvet/layers/sequential` | ok | yes |
| 19 | [`19-residual`](19-residual/) | layers/residual | `github.com/openfluke/welvet/layers/residual` | ok | yes |
| 20 | [`20-cnn`](20-cnn/) | layers/cnn1 · cnn2 · cnn3 | `github.com/openfluke/welvet/layers/cnn{1,2,3}` | ok | yes |
| 21 | [`21-rnn-lstm`](21-rnn-lstm/) | layers/rnn · lstm | `github.com/openfluke/welvet/layers/lstm` | ok | yes |
| 22 | [`22-seqmix`](22-seqmix/) | layers/seqmix — mixer contract | `github.com/openfluke/welvet/layers/seqmix` | ok | yes |
| 23 | [`23-gdn`](23-gdn/) | layers/gdn — gated delta net | `github.com/openfluke/welvet/layers/gdn` | ok | yes |
| 24 | [`24-mamba`](24-mamba/) | layers/mamba — selective SSM | `github.com/openfluke/welvet/layers/mamba` | ok | yes |
| 25 | [`25-convt`](25-convt/) | layers/convt1 · convt2 · convt3 | `github.com/openfluke/welvet/layers/convt{1,2,3}` | ok | yes |
| 26 | [`26-kmeans`](26-kmeans/) | layers/kmeans | `github.com/openfluke/welvet/layers/kmeans` | ok | yes |
| 27 | [`27-parallel`](27-parallel/) | layers/parallel — MoE + cameral | `github.com/openfluke/welvet/layers/parallel` | ok | yes |
| 28 | [`28-metacognition`](28-metacognition/) | layers/metacognition | `github.com/openfluke/welvet/layers/metacognition` | ok | yes |
| 29 | [`29-forward`](29-forward/) | runtime/forward | `github.com/openfluke/welvet/runtime/forward` | ok | yes |
| 30 | [`30-backward`](30-backward/) | runtime/backward | `github.com/openfluke/welvet/runtime/backward` | ok | yes |
| 31 | [`31-training`](31-training/) | runtime/training | `github.com/openfluke/welvet/runtime/training` | ok | yes |
| 32 | [`32-step`](32-step/) | runtime/step — step mesh | `github.com/openfluke/welvet/runtime/step` | ok | yes |
| 65 | [`65-cross-numeric`](65-cross-numeric/) | Cross-numeric train + down-the-dem | `github.com/openfluke/welvet/runtime/training` | ok | yes |
| 67 | [`67-train-modes`](67-train-modes/) | TrainMode — 29 named updates | `github.com/openfluke/welvet/layers/parallel` | ok | yes |
| 33 | [`33-dna`](33-dna/) | systems/dna | `github.com/openfluke/welvet/systems/dna` | ok | yes |
| 34 | [`34-evolution`](34-evolution/) | systems/evolution | `github.com/openfluke/welvet/systems/evolution` | ok | yes |
| 35 | [`35-tween`](35-tween/) | systems/tween | `github.com/openfluke/welvet/systems/tween` | ok | yes |
| 36 | [`36-tanhi`](36-tanhi/) | systems/tanhi — TANHI · UDP HUD | `github.com/openfluke/welvet/systems/tanhi` | ok | yes |
| 37 | [`37-telemetry`](37-telemetry/) | systems/telemetry | `github.com/openfluke/welvet/systems/telemetry` | ok | yes |
| 66 | [`66-lucy`](66-lucy/) | lucy — SoftAcc / Score measuring | `github.com/openfluke/welvet/lucy` | ok | yes |
| 69 | [`69-lucy-density`](69-lucy-density/) | Lucy density — synthetic organism | `github.com/openfluke/welvet/lucy` | ok | yes |
| 38 | [`38-entity`](38-entity/) | model/entity — .entity files | `github.com/openfluke/welvet/model/entity` | ok | yes |
| 39 | [`39-hf`](39-hf/) | model/hf — snapshots | `github.com/openfluke/welvet/model/hf` | ok | yes |
| 40 | [`40-tokenizer`](40-tokenizer/) | model/tokenizer | `github.com/openfluke/welvet/model/tokenizer` | ok | yes |
| 41 | [`41-sampling`](41-sampling/) | model/sampling | `github.com/openfluke/welvet/model/sampling` | ok | yes |
| 42 | [`42-transformer`](42-transformer/) | model/transformer — generate | `github.com/openfluke/welvet/model/transformer` | ok | yes |
| 43 | [`43-apps`](43-apps/) | apps — octo · flux2 · mosstts | `github.com/openfluke/welvet/apps/…` | partial | yes |
| 44 | [`44-octo`](44-octo/) | Octo — model shell | `github.com/openfluke/welvet/apps/octo` | ok | yes |
| 68 | [`68-cameral`](68-cameral/) | Cameral sandwiches + AAI Lucy | `github.com/openfluke/welvet/layers/parallel` | ok | yes |
| 70 | [`70-cam-sync`](70-cam-sync/) | CamSync — inter-cameral / cross-mesh weight blend | `github.com/openfluke/welvet/layers/parallel` | ok | yes |
| 45 | [`45-stub-seed`](45-stub-seed/) | stub/seed | `github.com/openfluke/welvet/stub/seed` | partial | yes |
| 46 | [`46-stub-serialization`](46-stub-serialization/) | stub/serialization | `github.com/openfluke/welvet/stub/serialization` | partial | yes |
| 47 | [`47-stub-memory`](47-stub-memory/) | stub/memory | `github.com/openfluke/welvet/stub/memory` | partial | yes |
| 48 | [`48-stub-donate`](48-stub-donate/) | stub/donate | `github.com/openfluke/welvet/stub/donate` | partial | yes |
| 49 | [`49-stub-fountain`](49-stub-fountain/) | stub/fountain | `github.com/openfluke/welvet/stub/fountain` | partial | yes |
| 50 | [`50-stub-hardware`](50-stub-hardware/) | stub/hardware | `github.com/openfluke/welvet/stub/hardware` | partial | yes |
| 51 | [`51-stub-accel`](51-stub-accel/) | stub/accel — NPU/Metal/QNN | `github.com/openfluke/welvet/stub/accel` | missing | yes |
| 52 | [`52-stub-clustering`](52-stub-clustering/) | stub/clustering | `github.com/openfluke/welvet/stub/clustering` | partial | yes |
| 53 | [`53-stub-ensemble`](53-stub-ensemble/) | stub/ensemble | `github.com/openfluke/welvet/stub/ensemble` | partial | yes |
| 54 | [`54-stub-evaluation`](54-stub-evaluation/) | stub/evaluation | `github.com/openfluke/welvet/stub/evaluation` | partial | yes |
| 55 | [`55-stub-grafting`](55-stub-grafting/) | stub/grafting | `github.com/openfluke/welvet/stub/grafting` | partial | yes |
| 56 | [`56-stub-grouping`](56-stub-grouping/) | stub/grouping | `github.com/openfluke/welvet/stub/grouping` | partial | yes |
| 57 | [`57-stub-introspection`](57-stub-introspection/) | stub/introspection | `github.com/openfluke/welvet/stub/introspection` | partial | yes |
| 58 | [`58-stub-observer`](58-stub-observer/) | stub/observer | `github.com/openfluke/welvet/stub/observer` | partial | yes |
| 59 | [`59-stub-pipeline`](59-stub-pipeline/) | stub/pipeline | `github.com/openfluke/welvet/stub/pipeline` | partial | yes |
| 60 | [`60-stub-templates`](60-stub-templates/) | stub/templates | `github.com/openfluke/welvet/stub/templates` | partial | yes |
| 61 | [`61-stub-universal`](61-stub-universal/) | stub/universal | `github.com/openfluke/welvet/stub/universal` | partial | yes |
| 62 | [`62-w2a`](62-w2a/) | w2a — validation harness | `github.com/openfluke/w2a` | ok | yes |
| 63 | [`63-validation`](63-validation/) | Validation report — full suite | `github.com/openfluke/w2a` | ok | yes |
| 64 | [`64-scorecard`](64-scorecard/) | Scorecard → v1.0 / minors | `—` | ok | yes |
| 71 | [`71-dispatch`](71-dispatch/) | runtime/dispatch — Op switchboard | `github.com/openfluke/welvet/runtime/dispatch` | ok | yes |
| 72 | [`72-wav2vec2`](72-wav2vec2/) | model/wav2vec2 — CTC ASR | `github.com/openfluke/welvet/model/wav2vec2` | ok | yes |
| 73 | [`73-all-layers`](73-all-layers/) | All layers — one forward each | `layers/*` | ok | yes |
| 74 | [`74-kokoro`](74-kokoro/) | Kokoro — empty model vs apps | `model/kokoro` · `apps/kokoro` | ok | yes |
| 75 | [`75-apps-map`](75-apps-map/) | Apps map — when/where/why | `apps/*` | ok | yes |
| 76 | [`76-nativecpu`](76-nativecpu/) | NativeCPU — opt-in welvet.cpp C ABI | `github.com/openfluke/welvet/nativecpp` | ok | yes |
| 77 | [`77-native-layers`](77-native-layers/) | NativeCPU — TrainNativeMSE per layer | `layers/*` + nativecpp | ok | yes |
| 78 | [`78-native-trainmodes`](78-native-trainmodes/) | NativeCPU — TrainMode stack / parallel / mesh / line | `layers/parallel` + nativecpp | ok | yes |
| 79 | [`79-qnn`](79-qnn/) | QNN — opt-in Hexagon HTP (welvet.qnn) | `nativeqnn` + dense | ok | yes |
| 80 | [`80-qnn-layers`](80-qnn-layers/) | QNN — HTP layer tour (session ABI) | `nativeqnn` | ok | yes |
| 81 | [`81-qnn-results`](81-qnn-results/) | QNN — go_v_qnn RESULTS (large shapes) | w2a RESULTS.md | ok | yes |
| 82 | [`82-apple-metal`](82-apple-metal/) | Apple Metal — mount and forward | `nativemetal` | ok | yes |
| 83 | [`83-apple-mamba-dtypes`](83-apple-mamba-dtypes/) | Apple Metal — Mamba parameter dtypes | `nativemetal + mamba` | ok | yes |
| 84 | [`84-apple-training`](84-apple-training/) | Apple Metal — training modes | `nativemetal + parallel + architecture` | ok | yes |
| 85 | [`85-apple-save-weights`](85-apple-save-weights/) | Apple Metal — weight return and ENTITY save | `nativemetal + entity` | ok | yes |
| 86 | [`86-apple-json-routing`](86-apple-json-routing/) | Apple Metal — JSON backend placement | `nativemetal + runtime/placement` | ok | yes |
| 87 | [`87-apple-mixed-backends`](87-apple-mixed-backends/) | Apple Metal — mixed Go/C++/Metal | `nativecpp + nativemetal + placement` | ok | yes |
| 88 | [`88-apple-mamba-results`](88-apple-mamba-results/) | Apple Metal — Mamba results and stress runner | `apple_stress_test` | ok | yes |

## Last smoke

```
go run ./cmd/runall
======== welvet book: ok=76+ fail=0 skip=0  # re-run after §79–81 ========
```

## Related

- Book HTML: https://openfluke.github.io/welvet/ (§76–78 NativeCPU · §79–81 QNN)
- Engine: `../../welvet` (local v1.3.0 checkout) · NPU: `welvet.qnn` 1.0.0
- Cam cookbook: [`../cam`](../cam/)
- Matrix: [`../../w2a/suites/go_v_qnn/RESULTS.md`](../../w2a/suites/go_v_qnn/RESULTS.md)

See also [`GAPS.md`](GAPS.md) for what still needs external weights.

## Apple examples

The small examples 82–87 need Apple silicon macOS, cgo and the installed Metal
dylib; 88 prints a labelled captured excerpt and works without a GPU. Keep a
normal build for example 87 so C++ and Metal can coexist. In 82–86,
`-tags welvet_metal` excludes the C++ binding.

For the full Mamba timing/determinism matrix, see
[`apple_stress_test`](apple_stress_test/README.md): all 34 parameter dtypes and
29 training modes. `./run_apple.sh` runs the usage tour without launching that
large stress matrix.

Apple usage tour 82–88: all seven examples verified on Apple M5/macOS arm64.
Example 83 also exercised all 34 parameter dtypes; example 84 exercised all
29 training modes. Tests live in w2a as `TestAppleNumberedExamples`.
