# WASM / npm / HTML examples

Every Welvet book chapter and Cam cookbook entry has:

| Path | How to run |
|------|------------|
| `welvet/<slug>/npm/run.mjs` | Node + `@openfluke/welvet` |
| `welvet/<slug>/html/index.html` | Browser (static server) |
| `cam/<slug>/npm` + `html` | Same for cameral cookbook |

## Setup

```bash
# Build the WASM package once
cd /path/to/welvet/apps/w2a/typescript
npm run build:all

# Point examples at it (or rely on sibling path detection)
export WELVET_TS=/path/to/welvet/apps/w2a/typescript

# Sync wasm_exec + main.wasm into example/_wasm/assets (for HTML)
cd /path/to/example
bash _wasm/sync-assets.sh
```

## Run all npm examples

```bash
export WELVET_TS=...
bash run-all-npm.sh
```

## HTML

Serve the **example repo root** (so `/_wasm/assets/main.wasm` and chapter HTML resolve):

```bash
bash _wasm/sync-assets.sh
npx --yes serve -l 4173 .
# open http://localhost:4173/welvet/11-dense/html/
# open http://localhost:4173/cam/01_modes/html/
```

Native-only chapters (`06-simd`, `07-webgpu`, `10-fusedgpu`, donate/hardware/accel, HF FS, LoadUniversal) print `SKIP` and exit 0.
