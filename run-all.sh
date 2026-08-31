#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")" && pwd)"
echo "==== welvet book ===="
(
  cd "$ROOT/welvet"
  # shellcheck disable=SC1091
  source ./env.sh
  go run ./cmd/runall
)
echo "==== cam cookbook ===="
(
  cd "$ROOT/cam"
  go run ./cmd/runall
)
echo "==== ALL EXAMPLE SUITES OK ===="
