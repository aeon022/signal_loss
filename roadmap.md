# ROADMAP: SIGNAL LOSS ENGINE DEVELOPMENT

This roadmap outlines the phased implementation plan for the Go Terminal Game Engine.

---

## Phase 1: Foundation & State Management — done
- [x] Initialize Go module (`go mod init signal_loss`).
- [x] Create package directory tree: `cmd/signal_loss/`, `pkg/game/`, `pkg/story/`, `pkg/tui/` (formerly `pkg/ui/`, replaced in Phase 7).
- [x] Implement `pkg/game/state.go`:
  - `GameState` struct definition with JSON tags.
  - `SaveState(path string)` & `LoadState(path string)` with safe file I/O.
  - `ApplyMutations(hull, bat, scrap int, itemsAdded, itemsRemoved []string, flags map[string]bool)`.
  - Verified via a scripted end-to-end playthrough (see Phase 5), not unit tests.

---

## Phase 2: Terminal UI & ANSI Rendering — superseded by Phase 7
- [x] Original version: hand-rolled `pkg/ui/terminal.go` (box-drawing frames, ANSI colors,
  typewriter print, HUD gauges) driving a blocking `bufio`-read loop in `pkg/game/engine.go`.
  Replaced entirely by the Bubble Tea TUI in Phase 7 — `pkg/ui` no longer exists.

---

## Phase 3: Data Parser & Story Ingestion — done
- [x] Implement `pkg/story/parser.go`:
  - Struct definitions matching `story.json` schema (`StoryData`, `Chapter`, `Choice`, `StatMutations`), plus `Location` (map labels), `Codex` (lexicon), and `Fatal` (instant-death choices).
  - `LoadStory(path string) (*StoryData, error)`.
  - Validation check to ensure all `next_chapter` pointers exist.

---

## Phase 4: Engine Loop & Time-Lock System — done, reshaped in Phase 7
- [x] `pkg/game/engine.go` originally held both game logic *and* the blocking print/read loop.
  As of Phase 7 it's pure logic only: `NewGameState`, `ResolveChoice` (mutate + report what
  happens next), `CheckGameOver`, `RandomInvalidResponse`, `ChapterOrder`. No I/O, no UI import —
  `pkg/tui` drives the loop now, this package is just the rules.
  - Real branching (unchanged): each chapter but the finale has one choice that's a dead end —
    `fatal` (instant game over) or a setback (`next_chapter` back to `chapter_1`).
  - Game over (0% Hull / 0% Bat, or `fatal`) and 3-ending victory detection — unchanged logic,
    now returned as a `Resolution` value instead of being printed inline.

---

## Phase 5: CLI Entry Point & Verification — done
- [x] `cmd/signal_loss/main.go`: flag parsing (`--fast`, `--save`, `--story`), then hands off to
  `tea.NewProgram(tui.New(...), tea.WithAltScreen()).Run()`. State is saved after every resolved
  choice (see Phase 7), so no separate SIGINT/SIGTERM handler is needed — see the `ponytail:`
  comment in main.go for why that tradeoff is fine here.
- [x] End-to-end playtests scripted via a pseudo-terminal (Python `pty`, since Bubble Tea needs a
  real tty for raw mode — plain piped stdin doesn't work post-Phase-7): full 12-chapter run to an
  ending, a `fatal` branch, a setback branch, save/resume, `codex`/`map` overlays, fast-mode boot
  skip — all verified against actual rendered frames and the save file, not just "it built".

---

## Phase 6: Content Depth (from user feedback, 2026-09-09)
- [x] Expanded prose per chapter (richer scene-setting, S.T.E.V.E. dialogue woven in).
- [x] Real branching: setback and fatal paths per chapter, not just flavor-only choices.
- [x] In-game Codex/lexicon (`c` key) explaining S.T.E.V.E., HULL/BAT/SCRAP, factions, etc.
- [x] In-game navigation map (`m` key) showing progress along the 12-location spine.
- [x] `story.md` regenerated from `story.json` (single source of truth) so both stay in sync — see `story.md` §3–5.

---

## Phase 7: Real TUI — Bubble Tea rewrite (from user feedback, 2026-09-09)
The plain-ANSI version above was "print statements with box-drawing," not a proper TUI: no
full-screen alt-buffer, no keyboard navigation, everything scrolled by in the terminal history.
Replaced with [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Elm-architecture TUI
runtime) + [Lipgloss](https://github.com/charmbracelet/lipgloss) (styling) — the standard Go
stack for this, not a hand-rolled alternative.
- [x] `pkg/tui/model.go`: the Bubble Tea `Model` (Init/Update/View), a state machine over
  `phaseBoot -> phaseChapter -> phaseGlitch -> phaseTimelock -> phaseGameOver/phaseEnding`,
  driven by `tea.Tick` for the typewriter reveal, boot animation, and glitch transitions —
  any keypress fast-forwards a reveal in progress instead of waiting it out.
- [x] `pkg/tui/view.go`: rendering — chapter frame, HUD bar, choice list (cursor-highlighted,
  arrow-key or digit-key selectable), codex/map as full overlays (`c`/`m`/`esc` toggle), all via
  Lipgloss styles instead of hand-built ANSI+padding math.
- [x] `pkg/tui/styles.go`: retro double-line border (`╔═╗║╚╝`), color palette, HUD gauge bars.
- [x] `cmd/signal_loss/main.go`: launches `tea.NewProgram(..., tea.WithAltScreen())` — real
  full-screen mode, not scrolling stdout.
- [x] `pkg/game/engine.go` reduced to pure logic (`ResolveChoice`, `CheckGameOver`, etc.) with
  zero UI imports, so the game rules are usable/testable independent of the rendering layer.
- [x] `pkg/ui` deleted — fully superseded, nothing left importing it.
- [x] Bug found and fixed during this rewrite: the chapter title in the frame was reading
  `state.CurrentChapter` live, which `ResolveChoice` already advances to the *next* chapter —
  so the title jumped ahead to "Chapter 2" while chapter 1's outcome text was still typewriter-
  revealing. Fixed by pinning `outcomeTitle` at the moment a choice is made.
- [x] Verified via a Python `pty`-driven test harness (Bubble Tea needs a real tty for raw mode,
  so plain piped stdin — used pre-Phase-7 — no longer works for testing): boot skip, codex/map
  toggle, a full chapter-1-to-outcome-to-chapter-2 transition, a full 12-chapter fast-mode run to
  an ending, and a `fatal` branch, all checked against actual rendered frames.

---

## Phase 8: Monitor Bezel, Fixed Viewport, Codex Gating, Real-Time Locks (2026-09-09)
- [x] **Fixed-size screen** — `frameStyle` got `.Height(FixedHeight)` alongside `.Width`, so
  every screen renders at the same 76×30 size regardless of content; codex (the one screen whose
  content can exceed that) scrolls inside it instead of growing the frame.
- [x] **Nested monitor bezel** — outer single-line housing with corner screws (`●───●`) and the
  terminal label / signal status embedded directly into its top/bottom edges, wrapping the inner
  "screen" (now a single-line border, not double) — see `chrome()` in pkg/tui/styles.go.
- [x] **Bug found and fixed:** `Width(MaxWidth)` already counts the padding columns, so the real
  text budget is `MaxWidth-2`, not `MaxWidth`. Wrapping text at the full `MaxWidth` let a few
  lines land exactly on the boundary, which Lipgloss then wrapped a second time — silently adding
  extra rows and breaking the fixed-height guarantee. Caught by comparing rendered line counts
  across screens, not by inspection. Fixed with a dedicated `TextWidth` constant.
- [x] **Second bug found and fixed:** `FrameWidth` (used to size the bezel) was off by 2 for the
  same padding-accounting reason. Only became visible once the bezel started wrapping every line
  in `│ … │`, where a 2-column mismatch is immediately obvious.
- [x] Instrument-panel layout: choices/continue-prompt anchored to the bottom of the fixed screen
  via `anchorBottom` (so they sit at a consistent position regardless of narrative length), and
  the HULL/BAT/SCRAP/INV status line moved out of the "screen" entirely, onto its own row between
  the screen and the bottom bezel edge — a console gauge strip, not part of the viewscreen.
- [x] **Codex gated by story progress** — `story.json` codex entries now carry `unlock_chapter`;
  the codex view only lists terms the player has actually reached (via a new `MaxChapterIndex` on
  `GameState`, which — unlike `CurrentChapter` — never regresses on a setback, so a setback
  doesn't make the codex "forget" things).
- [x] **`-realtime` flag** — chapters 3 and 9's time-locks become genuine real hours (via
  `TimeLockUntil`, a Unix timestamp already on `GameState`) instead of the compressed few-second
  wait. A pending real lock is enforced on load regardless of what flags the *next* launch uses —
  `-fast` alone can't skip it — and while the process is left running it self-unlocks the moment
  real time passes (1s recheck tick), so leaving the terminal open overnight also works.
- [x] Verified via pty tests: fixed line counts across chapter/codex/game-over screens, codex
  entry count growing from 9→13 as chapters unlock, a setback preserving `max_chapter_index`,
  triggering a real lock and confirming the save records the correct ~8h-out timestamp, that a
  fresh `-fast` (no `-realtime`) launch still honors a pending lock, and that an expired lock
  (simulated by editing the save) correctly resumes play — plus a full 12-chapter run throughout.

---

## Phase 9: More Time-Locks, Minute Granularity (2026-09-09)
- [x] `time_lock_hours` → `time_lock_minutes` throughout (`story.Choice`, `game.Resolution`,
  story.json) — hours-only couldn't express a "15 minute" beat, minutes cover both.
  `secondsPerGameHour` (compressed-mode constant) became `secondsPerGameMinute` at the same
  ratio (2s real / game-hour = 2s/60 real per game-minute), so the existing 8h/4h locks compress
  to the same on-screen duration as before (16s / 8s).
- [x] Three new shorter time-locks added, chosen per chapter rather than applied uniformly:
  Chapter 2 "Honest Barter" (15m, battery trickle-charging), Chapter 4 "Debug the Nomad" (20m,
  patch compiling), Chapter 7 "Patch the Leaks" (30m, waiting out the worst of the acid rain) —
  each choice's outcome text was lightly reworded so the wait reads as diegetic, not bolted on.
  Chapters 3 (8h) and 9 (4h, setback path) unchanged in substance, just re-expressed in minutes
  (480m / 240m).
- [x] Verified: compressed-mode timing for the new 15m lock, `-realtime` timing for the same
  lock (confirmed "14m 59s" / correct ~15min-out timestamp in the save), and a full 12-chapter
  fast-mode run completing normally with the new locks in place.
- [x] `story.md` §3–5 regenerated from `story.json` again (generator script updated for the
  renamed field and the `codex` entries' `unlock_chapter`).

---

## Phase 10: Play It in a Browser (2026-09-09)
The ask: embed the game in a website, playable online, same engine as the CLI — not a
reimplementation. Landing page lives at `~/Sites/signal-loss` (separate Astro repo, see its own
README); this phase is what changed here in `signal_loss` to make that possible.

- [x] **Feasibility spike first, before building anything:** `GOOS=js GOARCH=wasm go build` on
  `pkg/tui` failed outright — upstream `charmbracelet/bubbletea` has zero WASM support; `tea.go`
  references `p.initInput`, `openInputTTY`, `listenForResize`, etc., which only exist in
  `tty_unix.go` / `tty_windows.go`, both excluded on `GOOS=js`. Confirmed via a real compile
  attempt, not by reading docs.
- [x] Found the fix is tiny: `github.com/tmc/bubbletea` (a fork behind the community `bubbweb`
  project — verified via `gh api .../compare`) adds exactly two files for `GOOS=js`:
  `tty_js.go` (21 lines) and `signals_js.go` (9 lines), pure no-op stubs — no changes to existing
  files. Used directly via a pinned-commit `replace`, not vendored, so it's easy to audit and
  cheap to drop if upstream ever adds native `js` support.
- [x] **Scoped the fork to only the web build**, not the whole module: `cmd/wasm/` is its own Go
  module (own `go.mod`) with `replace signal_loss => ../..` (pulls in `pkg/tui`, `pkg/game`,
  `pkg/story` from the parent) and `replace github.com/charmbracelet/bubbletea => github.com/tmc/bubbletea ...`
  scoped to that module only. The native CLI's `go.mod` is untouched — still real upstream
  Bubble Tea v1.3.10. Verified independently: native build still passes after this change.
- [x] **Persistence decoupled from the filesystem.** `pkg/tui.Model` took a hardcoded
  `savePath string` + called `game.SaveState` directly — replaced with an injected
  `save func(*game.GameState) error`. The native CLI wires that to `game.SaveState`; the WASM
  build wires it to a JS callback that writes to `localStorage`. `pkg/tui` itself has no idea
  which one it's talking to.
- [x] **Story loading decoupled from the filesystem** the same way: `story.LoadStory(path)` now
  wraps a new `story.ParseStory(data []byte)` that does the real work. The WASM build fetches
  `story.json` over HTTP at runtime instead of reading a local file — which also means the site
  can swap that fetch for a CMS endpoint later (see the Orbiter idea below) without touching Go
  code at all.
- [x] `cmd/wasm/main.go`: bridges Bubble Tea's `tea.WithInput`/`tea.WithOutput` to two JS-exposed
  functions — `slInput` (xterm.js keystrokes in, already terminal-encoded — Bubble Tea's own
  input parser handles them exactly like a real tty) and an `onOutput` callback (Bubble Tea's
  ANSI output straight into `xterm.js`'s `write()`, no translation needed since xterm.js already
  speaks ANSI/VT100). `slStart` takes the story JSON, a saved-state JSON string (or empty), and
  the `fast`/`realtime` flags, and returns immediately — the actual program runs in a goroutine.
- [x] **Verified for real, not just "it compiles":** wrote a Puppeteer script driving real
  headless Chrome (no browser extension available this session) against the actual built Astro
  site — confirmed the WASM module boots, the boot sequence and chapter 1 render correctly
  (dumped via `.xterm-rows` per-row text, not `.xterm-screen innerText`, since xterm.js is
  canvas-rendered and `innerText` returns stale/inconsistent snapshots mid-animation), keyboard
  input reaches Bubble Tea and mutates state correctly (picked "Play Dead" → confirmed `-2 HULL`
  in the resulting `localStorage` save), and the chapter transition (glitch + chapter 2) fires.
  Also verified `npm run build` (the real static production build, not just `astro dev`) still
  serves a working `game.wasm` from `dist/`.
- [x] Binary size: stripped via `-ldflags="-s -w"` (7.1MB → 6.6MB). Not further optimized —
  gzip/brotli at the hosting layer would cut this substantially and is a deploy-config concern,
  not a code concern; noted as a TODO on the website side.

### Not built yet — on purpose (see also ~/Sites/signal-loss's own roadmap notes)
- **Orbiter as a content backend.** Right now the web build fetches a static `story.json` copy.
  Since `story.ParseStory` already takes raw bytes with no file-path assumption, pointing the
  fetch at an Orbiter-served endpoint instead of a static file is the whole change needed on the
  Go side — nothing to build here until that's actually wired up. The real work would be on the
  Orbiter/Astro side: modeling chapters as CMS content types matching the existing
  `Chapter`/`Choice`/`CodexEntry` shape in `pkg/story/parser.go`, so editing/adding chapters
  becomes a CMS edit, not a code change + WASM rebuild. `pkg/story`'s `LoadStory`/`ParseStory`
  split already exists specifically so this drop-in later doesn't require touching Go code again.
- **Story validation on the CMS side.** `ParseStory` already validates `next_chapter` references
  at parse time (existing check, not new) — worth surfacing that same validation in whatever
  Orbiter editing UI eventually exists, so a broken chapter reference is caught at save time in
  the CMS, not silently at runtime in someone's browser.

### Post-launch bugfix: typewriter reveal duplicating on screen (2026-09-09)
Reported live on the deployed dev site — during the typewriter reveal, each new frame appeared
as an *additional* line below the last instead of overwriting it in place, worst during longer
chapters.

- **My own verification method was the reason this shipped at all.** Every prior check in this
  session that "confirmed" correct rendering used either `-fast` (skips the incremental reveal
  entirely — the exact code path with the bug) or `.xterm-screen`/`.xterm-rows` snapshots taken
  *after* the animation had settled. A raw ANSI byte capture stripped of escape codes looks
  identical whether frames are correctly overwriting in place or genuinely stacking, because
  either way the same characters appear in transmission order — you can't tell "overwritten" from
  "duplicated" without something that actually *renders* the stream. Root-caused this properly
  using [`pyte`](https://pypi.org/project/pyte/) (a real VT100 emulator) to render the raw byte
  capture, which is what finally showed genuine duplicate rows rather than a clean single frame.
- **Root cause:** Bubble Tea's renderer needs to know the terminal's height to do its relative-
  cursor redraw math; on a real tty it gets this via `ioctl`/`SIGWINCH`, which is exactly what
  `tty_js.go` stubs out to a no-op for the WASM build (see Phase 10 above) — so the renderer's
  internal height silently stayed unset. Confirmed by reproducing the same bug on the *native*
  CLI too, by deliberately running it in a pty with an unset (0×0) window size, and confirming it
  went away the moment that pty was given an explicit, generous size via `TIOCSWINSZ`.
- **Fix:** `cmd/wasm/main.go` now explicitly sends `tea.WindowSizeMsg{Width: 84, Height: 37}`
  right after starting the program — 37 being the game's true total frame height (bezel + status
  line + the 30-row screen), matched exactly on the JS side by setting xterm.js's `rows: 37`
  (`GameTerminal.astro`) instead of the too-small `36` it had before. Sending it *before*
  `prog.Run()` starts crashed the WASM module outright (`Send()` blocks until the program's
  message loop is far enough along to receive) — has to go in its own goroutine, started after
  `go prog.Run()`.
- **Latent risk this surfaced, not yet addressed:** the *native* CLI has the same underlying
  fragility — a real terminal window shorter than 37 rows will hit this exact bug for real users,
  since Bubble Tea's own auto-detected height would then genuinely be too small. Not fixed here;
  worth either shrinking `FixedHeight` to fit more comfortably inside a typical terminal window,
  or detecting a too-small terminal at startup and failing with a clear message instead of
  silently corrupting the display.

---

## Phase 11: Launch Config, Quit Screen, and Going Public (2026-09-09)
- [x] **Scroll-triggers-input bug** — xterm.js turns mouse-wheel scrolling over a focused
  alt-screen terminal into arrow-key escape sequences (standard emulator behavior, for curses
  apps that want to handle their own scrolling) — Bubble Tea read those as real keypresses, so
  scrolling past the terminal on the page silently advanced the boot prompt. Fixed on the site
  side (capture-phase wheel listener that intercepts before xterm sees it, scrolls the page
  itself instead) — nothing to fix here in the engine, `cmd/wasm`/`pkg/tui` untouched.
- [x] **`-realtime` now has a UI**, not just a CLI flag: the site's pre-launch panel lets players
  pick text speed and compressed-vs-real-time locks before `slStart` is even called. Required
  extending `slStart`'s JS signature (adds a 7th `onExit` callback, see below) — no Go-side
  change beyond that, `fast`/`realtime` were already plain bool params.
- [x] **`[q]uit` had no graceful ending in a browser** — the CLI's Bubble Tea exit (return to a
  shell prompt) doesn't exist on a web page, so the terminal just went blank. `cmd/wasm/main.go`
  now invokes a JS `onExit` callback once `prog.Run()` returns; the site shows a "SIGNAL LOST"
  panel with a random S.T.E.V.E. one-liner, a Reconnect button (page reload — simplest reliable
  way to get a clean re-init given the existing double-boot guard), and a link to the repo.
- [x] **First real commits, and it's public now.** Both repos had been sitting as uncommitted
  working-tree changes all session. `signal_loss` (this repo) pushed to
  `github.com/aeon022/signal_loss` — **public**, a deliberate departure from every sibling repo
  in `~/Sites`/`~/Developing/Projects` being private, because the site explicitly links to it as
  a "view source" / download link. `~/Sites/signal-loss` (the Astro site) is *not* a separate
  GitHub repo — it keeps its own local git history but pushes to this same repo's `deploy/landing`
  branch, so `main` is the game and `deploy/landing` is the deployable site.

---

## Phase 12: German Localization (2026-09-09)
- [x] **Full `-lang de` mode**: a second complete story dataset (`story_de.json`, all 12
  chapters/36 choices/13 codex entries/endings translated, mechanically identical — same
  `next_chapter`/`flags`/`stat_mutations`/item names cross-referenced consistently within the
  file) plus a `pkg/tui/labels.go` catalog for every UI-chrome string (boot log, codex/map
  headings, hints, death/quit lines) that isn't story content. `HULL`/`BAT`/`SCRAP`/`INV` stay
  English acronyms in both languages — the status bar's fixed-width bar math is tuned to their
  exact lengths, and it's common practice even in fully localized games.
- [x] `pkg/game.CheckGameOver`/`ResolveChoice` now take the story data too, so the two
  stats-triggered death messages (hull breach, battery depleted) come from the story file's own
  `system_messages` section instead of being hardcoded English in a UI-agnostic package — with an
  English fallback if a story file omits that section, so old/custom story files don't break.
- [x] Native CLI: `-lang en|de` picks both the UI labels and (unless `-story` is passed
  explicitly) the default story file. Browser build: the launch panel gets a LANGUAGE toggle next
  to speed/mode; picking German fetches `story_de.json` instead of `story.json` and passes
  `lang` through to `slStart` (which grew an 8th param). Saves are namespaced per language
  (`signal_loss_save_en`/`_de`) since a German run's translated item names wouldn't match against
  English `items_removed` lookups if a save crossed languages.
- [x] Verified end-to-end via a pty test harness answering Bubble Tea's terminal-capability
  queries (OSC 11 background-color probe, cursor-position report) that a bare `pty.fork()`
  otherwise leaves hanging: full chapter progression with correct stat mutations, codex/map
  overlays, a fatal-choice death, and the chapter 12 finale — in German, cross-checked against
  identical English behavior to rule out regressions. Browser side verified via Puppeteer against
  a locally served `/signal_loss/` subpath build, both languages.
