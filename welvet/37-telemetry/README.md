# 37. systems/telemetry

**Part:** V · Systems  
**Package:** `github.com/openfluke/welvet/systems/telemetry`  
**Status:** ok — ✅

## When

Adaptation, measurement, or evolution (`systems/telemetry`).

## Where

`import "github.com/openfluke/welvet/systems/telemetry"`

```bash
cd 37-telemetry && source ../env.sh && go run .
```

## Why

Static structural blueprints (sizes, op kinds) differ from live tanhi events.

## What

ExtractNetworkBlueprint, ExtractLayerTelemetry for introspection/UIs.

## Sample output (captured)

```
{ID:demo TotalLayers:1 TotalParams:0 Layers:[{Z:0 Y:0 X:0 L:0 Type:Dense Activation:ReLU Parameters:0 InputShape:[] OutputShape:[] Branches:[] CombineMode:}]}
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/37-telemetry`).
