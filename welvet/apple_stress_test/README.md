# Apple Mamba stress test

A Go program comparing **Go CPU**, **welvet.cpp through Go's C ABI binding**, and
**welvet.metal through Go's mounted-session binding**. It runs Mamba forward
passes and all 29 training modes over all 34 native parameter dtypes. Activations
and boundary gradients are Float32. No external C++/Metal benchmark executable
is launched.

## Run

From this folder on Apple silicon macOS, with the sibling Go checkout and its
C++/Metal artifacts installed:

```sh
go run .
```

The default is the full **1,020-row stress matrix**: 34 forward comparisons plus
34 × 29 training comparisons. Each row runs all three backends. The full run is
intentionally long. Progress prints after every completed comparison and each
result is written immediately to CSV and JSONL. Ctrl+C stops between native calls.

For a small verification run:

```sh
go run . -profile smoke -samples 1 -iters 1 -steps 8 -out results/smoke
```

To measure the large Float32 workload first:

```sh
go run . -profile stress -dtypes float32 -samples 3 -out results/float32
```

Useful filters:

```sh
go run . -phase forward -dtypes all
go run . -phase train -modes NormalBP,StepBP,MeshBP -dtypes float32,float16,bfloat16
go run . -list
```

Profiles use Mamba Expand=2, DState=8:

| Profile | DModel | Batch | Sequence |
|---|---:|---:|---:|
| smoke | 8 | 2 | 3 |
| medium | 128 | 4 | 16 |
| stress (default) | 512 | 8 | 32 |
| large | 1024 | 16 | 64 |

Forward cases contain one Mamba. Training cases contain two independently seeded
Mamba layers. Step modes retain their actual Line pipeline between calls. Mesh
modes use an actual three-cell clocked Grid: disabled input ingress followed by
two Mamba cells. The Go/C++ cases use the public Go scheduler and credit APIs;
Metal uses the native Grid training API. Eight training calls per sample allow
the Line pipeline to produce updates; `-steps` can increase the training horizon.

## What the results mean

`results/<timestamp>/` contains `results.csv`, `results.jsonl` and `summary.json`.
Metadata records the native versions, Metal device identity, Go runtime, geometry,
parameters and options. Every row compares all three pairs: Go/C++, Go/Metal,
and C++/Metal.

- `steady_us_per_call`: median wall-clock latency across independently rebuilt,
  identically seeded samples. Forward uses `-iters` calls per sample; training
  uses `-steps`. GPU calls are synchronous, so completion is included.
- `setup_ms`: construction/mount and per-sample forward warm-up time, recorded
  separately. Training shader/credit warm-up uses a separate disposable model,
  preserving identical initial parameters and empty Line pipelines in timed samples. `final_pull_ms` records explicit
  weight return separately; these costs are not hidden in the reported kernel/API latency.
- `speedup_x = time(A)/time(B)` and latency reduction `100*(1-time(B)/time(A))`.
  A negative percentage means B was slower. Large workloads can expose GPU
  benefits, but the program reports measured slowdowns too—it does not assume Metal wins.
- `exact_pct`: actual percentage of Float32 bit patterns that match.
  `within_tolerance_pct`: percentage satisfying
  `abs(a-b) <= 1e-5 + 1e-3*max(abs(a),abs(b))`.
  Outputs and updated weights have separate metrics; these are numerical
  agreement percentages, **not model prediction accuracy**.
- `exact100`, `strict80`, `industry60`: the project's compatibility-band labels,
  distinct from those measured percentages. Strict/industry use maximum
  normalized error thresholds `1e-5`/`1e-3`. Lower agreement remains visible as
  `deviation`; `-strict` makes such rows fail the process.
- `replay_exact_pct`: repeatability against an independent same-backend run.
  It is null / `not_measured` when `-samples 1` is selected.
- `successful_cpp_calls`: checked C ABI calls made by the measured workload.
  A C++ row with no native calls is an execution error, not a Go fallback timing.

The C++ path calls the native Mamba forward/backward APIs. Go retains training
scheduling, SGD and some proxy-credit arithmetic for that path. Metal sessions
perform their training on the GPU and keep parameters mounted until the explicit
final pull. The table measures these **available end-to-end binding paths**, not
an assertion that every path has identical host overhead.

Best-effort determinism aims for 100%; the project's accepted target is industry60
through exact100. Low-bit encodings and longer training can diverge. Non-finite
values and execution errors are reported as errors, never as successful matches.

## Build requirements

Use a normal cgo Go build; **do not use `-tags welvet_metal`**, which disables C++.
The existing parent module resolves Go Welvet through `../../welvet`. Its C++
artifact must be available under `welvet/native/cpp/`; the Metal dylib is resolved
by the `nativemetal` package. Set `WELVET_METAL_LIBRARY` to select another dylib.
Rebuild/install Metal from its checkout with `./scripts/macos_install_go.sh`.

The Go checkout must contain the completed C++ layer dispatch and call-counter
API used by this example. The program fails clearly if either backend is absent.
