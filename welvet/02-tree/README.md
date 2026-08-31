# 2. Repository map

**Part:** I · Orientation  
**Package:** `—`  
**Status:** ok — ✅ layout

## When

Orientation chapter: Repository map.

## Where

Book / repo map (no single import)

```bash
cd 02-tree && source ../env.sh && go run .
```

## Why

Readers need a single map of what is engine vs harness vs app vs stub.

## What

Top-level folders and their ownership. Engine packages never import w2a.

## Sample output (captured)

```
import github.com/openfluke/welvet/core/…
import github.com/openfluke/welvet/weights/…
import github.com/openfluke/welvet/quant/…
import github.com/openfluke/welvet/simd/…
import github.com/openfluke/welvet/webgpu/…
import github.com/openfluke/welvet/tiling/…
import github.com/openfluke/welvet/architecture/…
import github.com/openfluke/welvet/fusedgpu/…
import github.com/openfluke/welvet/layers/…
import github.com/openfluke/welvet/runtime/…
import github.com/openfluke/welvet/systems/…
import github.com/openfluke/welvet/lucy/…
import github.com/openfluke/welvet/model/…
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/02-tree`).
