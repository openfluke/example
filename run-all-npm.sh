#!/usr/bin/env bash
# Run every chapter npm/run.mjs (welvet + cam).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)"
export WELVET_TS="${WELVET_TS:-$ROOT/../welvet/apps/w2a/typescript}"
if [[ ! -f "$WELVET_TS/dist/index.js" ]]; then
  echo "Build first: cd \$WELVET_TS && npm run build:all" >&2
  exit 1
fi

ok=0
skip=0
fail=0
run_one() {
  local label="$1" script="$2"
  if out=$(node "$script" 2>&1); then
    if echo "$out" | grep -q '^SKIP '; then
      echo "[SKIP] $label"
      skip=$((skip + 1))
    else
      echo "[OK]   $label"
      ok=$((ok + 1))
    fi
  else
    echo "[FAIL] $label"
    echo "$out" | tail -20
    fail=$((fail + 1))
  fi
}

while IFS= read -r -d '' f; do
  slug=$(basename "$(dirname "$(dirname "$f")")")
  run_one "welvet/$slug" "$f"
done < <(find "$ROOT/welvet" -path '*/npm/run.mjs' -print0 | sort -z)

while IFS= read -r -d '' f; do
  slug=$(basename "$(dirname "$(dirname "$f")")")
  run_one "cam/$slug" "$f"
done < <(find "$ROOT/cam" -path '*/npm/run.mjs' -print0 | sort -z)

echo "=== ok=$ok skip=$skip fail=$fail ==="
[[ "$fail" -eq 0 ]]
