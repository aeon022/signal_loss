#!/usr/bin/env bash
# Builds the Go game to WASM and copies it (plus its story data) into
# public/game/, where the Astro page loads it from. Run this whenever
# signal_loss's Go source or story.json changes.
set -euo pipefail

SITE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GAME_DIR="${SIGNAL_LOSS_DIR:-$SITE_DIR/../../Developing/Projects/signal_loss}"
OUT_DIR="$SITE_DIR/public/game"

if [ ! -d "$GAME_DIR" ]; then
  echo "signal_loss repo not found at $GAME_DIR (set SIGNAL_LOSS_DIR to override)" >&2
  exit 1
fi

mkdir -p "$OUT_DIR"

echo "Building game.wasm from $GAME_DIR ..."
(cd "$GAME_DIR/cmd/wasm" && GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o "$OUT_DIR/game.wasm" .)

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$OUT_DIR/wasm_exec.js"
cp "$GAME_DIR/story.json" "$OUT_DIR/story.json"
cp "$GAME_DIR/story_de.json" "$OUT_DIR/story_de.json"

echo "Done: $OUT_DIR/{game.wasm,wasm_exec.js,story.json,story_de.json}"
