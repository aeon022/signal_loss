#!/bin/sh
# Installs signal_loss on Linux and macOS: downloads the release binary for
# this machine (no Go needed), puts it in ~/.local/bin and says what to do if
# that folder is not on PATH. Falls back to `go install` if no binary exists.
#
#   curl -fsSL https://raw.githubusercontent.com/aeon022/signal_loss/main/install.sh | sh
#
# VERSION=v0.4.1 picks a version, INSTALL_DIR=/usr/local/bin another folder.
set -eu

repo=aeon022/signal_loss
dest=${INSTALL_DIR:-$HOME/.local/bin}

os=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$os" in
  linux | darwin) ;;
  *) echo "Unsupported system '$os'. On Windows use install.ps1." >&2; exit 1 ;;
esac
case "$(uname -m)" in
  x86_64 | amd64) arch=amd64 ;;
  arm64 | aarch64) arch=arm64 ;;
  *) echo "Unsupported CPU '$(uname -m)'." >&2; exit 1 ;;
esac

fetch() { if command -v curl >/dev/null 2>&1; then curl -fsSL "$1" -o "$2"; else wget -q "$1" -O "$2"; fi; }

tag=${VERSION:-}
if [ -z "$tag" ]; then # the redirect of /releases/latest ends in the newest tag
  tag=$(curl -fsSIL -o /dev/null -w '%{url_effective}' "https://github.com/$repo/releases/latest" | sed 's|.*/||')
fi

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
name=signal_loss_${tag#v}_${os}_${arch}

mkdir -p "$dest"
if fetch "https://github.com/$repo/releases/download/$tag/$name.tar.gz" "$tmp/pkg.tar.gz" 2>/dev/null; then
  if fetch "https://github.com/$repo/releases/download/$tag/checksums.txt" "$tmp/sums" 2>/dev/null; then
    want=$(grep " $name.tar.gz\$" "$tmp/sums" | cut -d' ' -f1)
    if command -v sha256sum >/dev/null 2>&1; then got=$(sha256sum "$tmp/pkg.tar.gz" | cut -d' ' -f1); else got=$(shasum -a 256 "$tmp/pkg.tar.gz" | cut -d' ' -f1); fi
    [ "$want" = "$got" ] || { echo "Checksum mismatch, not installing." >&2; exit 1; }
  fi
  tar -xzf "$tmp/pkg.tar.gz" -C "$tmp" signal_loss
  install -m 755 "$tmp/signal_loss" "$dest/signal_loss"
elif command -v go >/dev/null 2>&1; then
  echo "No binary for $tag ($os/$arch); building with Go instead..."
  GOBIN=$dest go install "github.com/$repo/cmd/signal_loss@${tag:-latest}"
else
  echo "No binary for $tag ($os/$arch) and no Go installed. See https://github.com/$repo/releases" >&2
  exit 1
fi

echo "Installed signal_loss $tag to $dest/signal_loss"
case ":$PATH:" in
  *":$dest:"*) echo "Play with:  signal_loss" ;;
  *) echo "Add it to your PATH, then run 'signal_loss':"
     echo "  echo 'export PATH=\"$dest:\$PATH\"' >> ~/.profile   # and restart the terminal" ;;
esac
