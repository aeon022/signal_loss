# Sci-Fi Text Adventure Game

> Note: the premise below (lifeline device, day/night cycle, flashlight) is an early draft
> that was superseded by the actual story bible in `story.md` (SIGNAL LOSS: SECTOR SCHRÖDINGER
> — Arthur Pendelton, S.T.E.V.E., Gryps-4). Kept here for history, not implemented as written.

Welcome to the game! You are a space explorer who has crash landed on an unknown planet. Your ship is damaged and you have limited resources. Your mission is to survive long enough to repair the ship and escape the planet.

## Story
You are a lone space explorer who has crash landed on an unknown planet. The player's ship is equipped with a "lifeline" device that allows them to communicate with their team back home. However, using the lifeline consumes valuable resources and puts the player at risk of being detected by hostile forces.

Throughout the game, you will encounter various challenges such as hostile alien creatures, dangerous terrain, and scarce resources. To overcome these challenges, you must make strategic decisions about how to allocate your limited resources. For example, should you spend resources on repairing the ship or stockpiling food?

The planet is inhabited by several alien factions, each with its own agenda and relationship with the player. Some factions may be friendly and willing to trade resources for information, while others may be hostile and attack on sight. Players must navigate these relationships carefully to survive.

The game has a day/night cycle that affects visibility, energy levels, and enemy behavior. During the night, players will be more vulnerable to enemies and must use their flashlight sparingly to conserve battery life.

## Current Implementation Status

_Last updated 2026-09-09 — see `roadmap.md` for the full phase-by-phase build log._

- ✅ State Management: JSON serialization, save/resume, `max_chapter_index` tracking for the
  codex gate (pkg/game/state.go)
- ✅ Terminal UI: real Bubble Tea TUI, full-screen alt-buffer, nested monitor bezel (single-line
  "screen" inside an outer housing with corner screws and embedded label/status plates), boot
  sequence, glitch transitions, HUD instrument-panel line below the screen, arrow-key/digit-key
  choice navigation anchored to the bottom of a fixed-size viewport, codex/map overlays with
  codex scrolling (pkg/tui/) — see roadmap.md Phase 7/8 for how this evolved
- ✅ Story Parser: JSON loading, chapter/choice/codex (with unlock_chapter) schema, `next_chapter`
  validation, per-story `system_messages` (localized hull-breach/battery-depleted death lines)
  (pkg/story/parser.go)
- ✅ Game Engine: pure logic only — state mutation, fatal/setback/ending resolution, game-over
  checks, chapter-index helpers (pkg/game/engine.go). No I/O, no UI, no language awareness;
  `pkg/tui` drives the loop and owns all localized chrome.
- ✅ CLI Entry Point: flag parsing (`-fast`/`-realtime`/`-lang`/`-save`/`-story`), launches the
  Bubble Tea program (cmd/signal_loss/main.go)
- ✅ story.json + story_de.json: full 12-chapter narrative in **English and German**, real
  branching (fatal/setback paths), codex lexicon gated by story progress — both files
  mechanically identical (same flags/mutations/branching), verified via cross-diff and the
  actual Go parser
- ✅ `pkg/tui/labels.go`: every UI-chrome string (boot log, codex/map headings, hints, death/quit
  lines) as an English/German catalog, set once per `Model` via `-lang`/the browser's language
  toggle — `HULL`/`BAT`/`SCRAP`/`INV` deliberately stay English acronyms in both (status-bar
  width math is tuned to them)
- ✅ install.sh: one-script build (Go check + `go mod tidy` + `go build`)
- ✅ Codex only shows lexicon entries the story has actually introduced (by furthest chapter
  reached, not current — a setback doesn't make you "forget" things)
- ✅ `-realtime`: several chapters' time-locks (15m/30m/8h/4h, minute-granularity) become genuine
  real time instead of a compressed few seconds — the wait is enforced via a timestamp in the
  save file, so it survives quitting and can't be bypassed by relaunching with `-fast`
- ✅ **Browser build**: `cmd/wasm/` (separate Go module, pins a small `tmc/bubbletea` fork for
  `GOOS=js` support) compiles the exact same engine to WebAssembly — no reimplementation.
  Bridged into an xterm.js terminal on an Astro landing page (`~/Sites/signal-loss`, this repo's
  `deploy/landing` branch), with a pre-launch panel (text speed / time-lock mode / **language**),
  a "SIGNAL LOST" quit screen, and localStorage-backed saves namespaced per language.
- ✅ **Public and deployed**: `github.com/aeon022/signal_loss` (public, MIT-licensed, README,
  `FUNDING.yml`), tagged `v0.1.0` for `go install`. Landing page live at
  **https://aeon022.github.io/signal_loss/** via a GitHub Actions workflow that builds the WASM
  game + Astro site and deploys to Pages on every push to `deploy/landing`. Retro-futuristic
  site design (starfield backdrop, scanline/CRT overlay, chromatic title glow, hazard-stripe
  dividers, shrink-on-scroll header with a mobile burger menu), abteilung83 branding/signature,
  a Polar.sh "Support this project" link, and an honest cookie-consent notice (no tracking,
  save file only, in `localStorage`).

## Next Steps

Engine, content, UI, browser port, deployment, and localization are all feature-complete for the
current scope. Open backlog — see `roadmap.md`'s **What's Next** section for the full list with
rationale; condensed here:

**Site:** localize the marketing page itself (game speaks German now, the page around it
doesn't yet), a custom 404, a devlog page adapted from `roadmap.md`'s war stories, privacy-
respecting analytics (currently zero visibility into whether anyone's playing), an itch.io
listing, a custom domain under abteilung83.at.

**Game:** terminal SFX (browser-only, via Web Audio), cross-run stats beyond the current
per-save codex gate (endings seen, causes of death), more languages if there's real demand
(the `-lang`/`story_<lang>.json` pattern already generalizes), Orbiter-backed chapters (once
there's actually more content to justify it), a Homebrew formula (needs a `goreleaser` release
pipeline, bigger follow-up).

Smaller/older open ideas, still not started:
1. More than one setback destination (currently all setbacks route to chapter_1).
2. Automated tests (`go test`) beyond the scripted pty-driven playthroughs used so far.
3. Window-size-aware layout (currently a fixed 76-column content width, doesn't reflow on resize).
