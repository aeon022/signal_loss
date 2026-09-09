#!/usr/bin/env bash
# Builds the signal_loss game. Installs Go via Homebrew if it's missing
# (macOS only outside that, you're on your own).
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"

if ! command -v go >/dev/null 2>&1; then
  echo "Go not found."
  if command -v brew >/dev/null 2>&1; then
    echo "Installing Go via Homebrew..."
    brew install go
  else
    echo "Install Go from https://go.dev/dl/ and re-run this script." >&2
    exit 1
  fi
fi

echo "Fetching dependencies..."
go mod tidy

echo "Building signal_loss..."
go build -o signal_loss ./cmd/signal_loss

echo
echo "Done. Play with:"
echo "  ./signal_loss          # normal speed"
echo "  ./signal_loss -fast    # skip typewriter effect and time-locks"
