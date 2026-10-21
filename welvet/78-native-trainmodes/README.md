# 78. NativeCPU — TrainMode stack / parallel / mesh / line

**Part:** X · NativeCPU (welvet.cpp)  
**Package:** `github.com/openfluke/welvet/nativecpp` + layers  
**Status:** ok — ✅ linux/amd64 + cgo

## When

You want in-process welvet.cpp MatVec / NormalBP / TrainMSE instead of Go CPU tiled.

## Where

`import "github.com/openfluke/welvet/nativecpp"` and layer `TrainNativeMSE` / `TrainStackMSE`.

```bash
cd 78-native-trainmodes && source ../env.sh && go run .
```

## Why

Opt-in `BackendNativeCPU` keeps Go as default. Mirror Capture + `SyncWeightsFromNative` before `.entity`.

## What

See [feature book §78](https://openfluke.github.io/welvet/chapters/78-native-trainmodes.html).

## Sample output (captured)

```
(run locally — needs linux/amd64 libwelvet)
```

## Source

Mirrored from the [Welvet feature book](https://openfluke.github.io/welvet/) (`openfluke.github.io/welvet/examples/78-native-trainmodes`).
