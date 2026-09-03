# 36. systems/tanhi — TANHI · UDP HUD

**Part:** V · Systems  
**Package:** `github.com/openfluke/welvet/systems/tanhi`  
**Status:** ok — ✅

## When

Adaptation, measurement, or evolution (`systems/tanhi`).

## Where

`import "github.com/openfluke/welvet/systems/tanhi"`

```bash
cd 36-tanhi && source ../env.sh && go run .
```

## Why

Training visualization must never block the engine — best-effort UDP JSON-lines to a HUD.

## What

**TANHI** = *Tensor Activation Network Holographic Interface*. Sparse non-blocking JSON-line UDP events for per-layer forward/backward HUD visualization. ConfigFromGrid, Emit/EmitSweep, DefaultUDPPort (17481). SoulGlitch-style consumers.

_No captured output in `_manifest.json` yet — run `go run .` and paste here._
## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/36-tanhi`).

## Sample output (captured)

```
— 36-tanhi (996ms)
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).

