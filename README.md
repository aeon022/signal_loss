# signal-loss

**Live: https://aeon022.github.io/signal_loss/**

Landing page + browser-playable build of [SIGNAL LOSS: SECTOR SCHRÖDINGER](https://github.com/aeon022/signal_loss),
a retro-futuristic sci-fi text adventure. The game itself is the real Go/Bubble Tea engine from
that repo, compiled to WebAssembly — not a JS reimplementation. See that repo's `roadmap.md`
Phase 10 for how the WASM bridge works.

This source lives on the `deploy/landing` branch of the [`signal_loss`](https://github.com/aeon022/signal_loss)
repo itself (the game's `main` branch has the Go source; this one has the site) — not a separate
GitHub repository.

## Requirements

- Node.js ≥ 22.12
- Go (same toolchain used by `signal_loss`) — only needed to run `build:wasm`, not for the Astro
  site itself
- A checkout of `signal_loss` as a sibling of `~/Sites` at `~/Developing/Projects/signal_loss`
  (or set `SIGNAL_LOSS_DIR` to point elsewhere — see `scripts/build-wasm.sh`)

## Development

```sh
npm install
npm run build:wasm   # compiles signal_loss's Go game to public/game/game.wasm
npm run dev          # http://localhost:4321
```

Re-run `npm run build:wasm` any time the Go source or `story.json` in the `signal_loss` repo
changes — the compiled `.wasm` is not checked in (see `.gitignore`), so a fresh clone needs it
built once before `npm run dev`/`npm run build` will actually show a working game.

## Production build

```sh
npm run build:wasm
npm run build         # -> dist/
npm run preview       # sanity-check the static build locally
```

## How the game gets into the page

`src/components/GameTerminal.astro` boots `public/game/game.wasm` (Go's WASM glue via
`wasm_exec.js`) and wires its input/output to an [xterm.js](https://xtermjs.org) terminal:
keystrokes go into Bubble Tea's own input parser exactly like a real tty would send them, and
Bubble Tea's ANSI output goes straight into `xterm.write()` — no translation layer, xterm.js
already speaks the same ANSI/VT100 that a real terminal does. Save state round-trips through
`localStorage` via a JS callback the Go side calls after every resolved choice.

## Deployment

`.github/workflows/deploy-pages.yml` builds and deploys to GitHub Pages on every push to
`deploy/landing`: it checks out this branch plus `signal_loss`'s `main` as a sibling directory
(so `build:wasm` has Go source to compile), then `astro build` → Pages. `base`/`site` in
`astro.config.mjs` are set for Pages' project-site subpath (`/signal_loss/`) — every asset
reference in the codebase goes through `import.meta.env.BASE_URL` rather than a root-relative
path so this isn't hardcoded in more than one place.

## Content

`public/game/story.json` is a build-time copy of the game's story data, fetched at runtime by
the WASM build (not baked into the binary) specifically so a future content backend (see
`signal_loss/roadmap.md` Phase 10's Orbiter note) only needs to change *where* that fetch points,
not the Go code.
