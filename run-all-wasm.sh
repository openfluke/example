#!/usr/bin/env bash
# Run every Welvet + cam chapter as npm WASM examples.
# Legacy quick smoke: WELVET_SMOKE_ONLY=1 bash run-all-wasm.sh
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)"
export WELVET_TS="${WELVET_TS:-$ROOT/../welvet/apps/w2a/typescript}"
if [[ ! -f "$WELVET_TS/dist/index.js" || ! -f "$WELVET_TS/dist/main.wasm" ]]; then
  echo "Building WASM package at $WELVET_TS …"
  (cd "$WELVET_TS" && npm run build:all)
fi

if [[ "${WELVET_SMOKE_ONLY:-}" == "1" ]]; then
  echo "==== welvet WASM smoke ===="
  node "$ROOT/welvet/wasm/run-smoke.mjs"
  echo "==== cam WASM smoke ===="
  node "$ROOT/cam/wasm/run-modes.mjs"
  echo "==== ALL EXAMPLE WASM OK ===="
  exit 0
fi

bash "$ROOT/_wasm/sync-assets.sh"
exec bash "$ROOT/run-all-npm.sh"
