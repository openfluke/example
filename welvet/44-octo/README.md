# 44. Octo — model shell

**Part:** VII · Apps  
**Package:** `github.com/openfluke/welvet/apps/octo`  
**Status:** ok — ✅ runs

## When

You need the `apps/octo` foundation package.

## Where

`import "github.com/openfluke/welvet/apps/octo"`

```bash
cd 44-octo && source ../env.sh && go run .
```

## Why

A model is only useful with a shell around it: pull weights from Hugging Face, convert them to a Welvet .entity, then chat, serve, or benchmark. Octo is that shell, kept in its own module so the engine never depends on an app.

## What

Subcommands cover the whole loop: hub download/ensure repos, convert pack to a single .entity, interactive run/serve/chat, plus image and speech menus. Its bench harness sweeps every quant format across a CPU Plan 9 SIMD fused profile and a WebGPU fused profile.

## Sample output (captured)

```
Octo is its own module — the engine never imports it.
  cd apps/octo && go run .
  download -> convert -> .entity -> chat / serve / bench
```

Live capture from `go run ./cmd/runall` on local `../../welvet` (exit 0).


## Source

Copied from the [Welvet feature book](https://openfluke.github.io/welvet/) examples (`openfluke.github.io/welvet/examples/44-octo`).
