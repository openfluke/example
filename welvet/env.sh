#!/usr/bin/env bash
# Shared Go build cache for book examples (avoids /tmp quota exhaustion with webgpu CGO).
_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
export GOCACHE="${GOCACHE:-$_root/.cache/gocache}"
export GOTMPDIR="${GOTMPDIR:-$_root/.cache/gotmp}"
export TMPDIR="${TMPDIR:-$_root/.cache/tmp}"
mkdir -p "$GOCACHE" "$GOTMPDIR" "$TMPDIR"

# Hexagon HTP runtime DLL (welvet.qnn) — needed for §79–80 on windows/arm64.
_qnn_dll="$_root/../../welvet/native/qnn/windows_arm64"
if [[ -d "$_qnn_dll" ]]; then
  case ":${PATH}:" in
    *":$_qnn_dll:"*) ;;
    *) export PATH="$_qnn_dll:$PATH" ;;
  esac
fi
