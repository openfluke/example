# 79. QNN — opt-in Hexagon HTP (welvet.qnn)

**Part:** XI · Hexagon NPU (welvet.qnn)  
**Package:** `github.com/openfluke/welvet/nativeqnn` + layers  
**Status:** ok — ✅ windows/arm64 + cgo

## When

You want Dense (and friends) on Qualcomm Hexagon HTP instead of Go CPU / NativeCPU / WebGPU.

## Where

```bash
cd 79-qnn && source ../env.sh && go run .
```

Needs shipped `libwelvet_qnn.dll` + QAIRT on PATH (Snapdragon X / Windows ARM64).

## Why

Opt-in `BackendQNN` + explicit `SyncToQNN` / `PrepareQNN` — never Prefer as default. Caps reject unsupported mounts.

## What

`Available()`, `Version()`, `SupportedLayers()`, Dense Sync → Prepare → Forward.

See [feature book §79](https://openfluke.github.io/welvet/chapters/79-qnn.html). Layer tour: §80. RESULTS: §81.

## Sample output (captured)

```
(run locally — needs windows/arm64 libwelvet_qnn)
```

## Source

Mirrored for the [Welvet feature book](https://openfluke.github.io/welvet/) (`openfluke.github.io` §79).
