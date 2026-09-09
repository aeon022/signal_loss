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
- ✅ State Management: JSON serialization, save/resume, `max_chapter_index` tracking for the
  codex gate (pkg/game/state.go)
- ✅ Terminal UI: real Bubble Tea TUI, full-screen alt-buffer, nested monitor bezel (single-line
  "screen" inside an outer housing with corner screws and embedded label/status plates), boot
  sequence, glitch transitions, HUD instrument-panel line below the screen, arrow-key/digit-key
  choice navigation anchored to the bottom of a fixed-size viewport, codex/map overlays with
  codex scrolling (pkg/tui/) — see roadmap.md Phase 7/8 for how this evolved
- ✅ Story Parser: JSON loading, chapter/choice/codex (with unlock_chapter) schema, next_chapter
  validation (pkg/story/parser.go)
- ✅ Game Engine: pure logic only — state mutation, fatal/setback/ending resolution, game-over
  checks, chapter-index helpers (pkg/game/engine.go). No I/O; pkg/tui drives the loop.
- ✅ CLI Entry Point: flag parsing (`--fast`/`--realtime`/`--save`/`--story`), launches the Bubble
  Tea program (cmd/signal_loss/main.go)
- ✅ story.json: full 12-chapter narrative with real branching (fatal/setback paths), codex
  lexicon gated by story progress
- ✅ install.sh: one-script build (Go check + `go mod tidy` + `go build`)
- ✅ Codex only shows lexicon entries the story has actually introduced (by furthest chapter
  reached, not current — a setback doesn't make you "forget" things)
- ✅ `-realtime`: chapters 3/9's time-locks become genuine real hours instead of a compressed
  few seconds — the wait is enforced via a timestamp in the save file, so it survives quitting
  and can't be bypassed by relaunching with `-fast`

## Next Steps
Engine, content, and UI are feature-complete for the current scope (see roadmap.md Phase 8).
Open ideas, not started:
1. More than one setback destination (currently all setbacks route to chapter_1).
2. Automated tests (`go test`) beyond the scripted pty-driven playthroughs used so far.
3. Window-size-aware layout (currently a fixed 76-column content width, doesn't reflow on resize).
