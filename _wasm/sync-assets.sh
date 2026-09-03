#!/usr/bin/env bash
# Copy main.wasm + wasm_exec.js into example/_wasm/assets for HTML demos.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
ASSETS="$ROOT/_wasm/assets"
mkdir -p "$ASSETS"

if [[ -n "${WELVET_TS:-}" && -f "$WELVET_TS/dist/main.wasm" ]]; then
  SRC="$WELVET_TS/dist"
elif [[ -f "$ROOT/../welvet/apps/w2a/typescript/dist/main.wasm" ]]; then
  SRC="$ROOT/../welvet/apps/w2a/typescript/dist"
elif [[ -f "$ROOT/../chaosglue/welvet/apps/w2a/typescript/dist/main.wasm" ]]; then
  SRC="$ROOT/../chaosglue/welvet/apps/w2a/typescript/dist"
else
  echo "Set WELVET_TS to apps/w2a/typescript (with dist/main.wasm after npm run build:all)" >&2
  exit 1
fi

cp -f "$SRC/main.wasm" "$ASSETS/main.wasm"
cp -f "$SRC/wasm_exec.js" "$ASSETS/wasm_exec.js"
echo "synced from $SRC -> $ASSETS"
ls -lh "$ASSETS"
