# AGENT SPECIFICATION: GO TERMINAL GAME ENGINE

## 1. Role & Objective
* **Role:** Senior Systems Engineer & Senior Go Developer.
* **Objective:** Build a robust, modular, data-driven terminal text-adventure engine in pure Go.
* **Scope:** Provide the complete runtime architecture, state persistence, ANSI UI rendering, real-time timer locks, and dynamic story ingestion. All narrative content is decoupled and loaded from `story.json`.

---

## 2. Tech Stack & Constraints
* **Language:** Go (Latest stable version).
* **Dependencies:** **Strictly zero external dependencies.** Only the Go Standard Library (`bufio`, `encoding/json`, `flag`, `fmt`, `math/rand`, `os`, `strings`, `sync`, `time`).
* **Cross-Platform:** Pure ANSI escape sequences for terminal control, compatible with macOS, Linux, and Windows terminals supporting VT100/ANSI.

---

## 3. Architecture & Package Structure
```text
.
├── cmd/
│   └── signal_loss/
│       └── main.go          # CLI entry point, flag parsing, engine initialization
├── pkg/
│   ├── game/
│   │   ├── state.go         # GameState struct, JSON save/load, stat mutation logic
│   │   └── engine.go        # Main game loop, time-lock enforcement, input processing
│   ├── story/
│   │   └── parser.go        # Story, Chapter, Choice, Dialogue Go structs & loader
│   └── ui/
│       └── terminal.go      # ANSI color wrappers, typewriter effect, reader, status bar
├── agent.md                 # Technical architecture and constraints (this file)
├── story.json               # Pure narrative and gameplay data
├── story.md                 # Human-readable story bible
└── roadmap.md               # Step-by-step development checklist
```

---

## 4. Core System Mechanics

### 4.1. CLI Flags
* `--fast` or `-f`: Boolean flag to bypass `printSlow` typewriter delays globally during testing and debugging.
* `--save <path>`: Optional flag to specify custom save file path (defaults to `.signal_loss_save.json`).
* `--story <path>`: Optional flag to specify custom story JSON path (defaults to `story.json`).

### 4.2. Terminal UI & Input Handling (`pkg/ui/terminal.go`)
* **Typewriter Effect:** `PrintSlow(text string, delay time.Duration, fastMode bool)`.
* **Safe Input Reader:** Read user input cleanly using `bufio.NewReader(os.Stdin)`, trimming CRLF/LF and handling empty lines gracefully.
* **Dynamic Status Bar:** Render a persistent stats bar before every prompt:
  ```text
  [HULL: 12% | BAT: 45% | SCRAP: 2 | INV: 1 items]
  ```
* **ANSI Color Palette:**
  * **Cyan (`\033[36m`):** AI / Companion dialogue (S.T.E.V.E.) and system status headers.
  * **Yellow (`\033[33m`):** Prompts, choices, resource warnings.
  * **Green (`\033[32m`):** Stat increases, item acquisitions, positive outcomes.
  * **Red/Pink (`\033[31m` / `\033[35m`):** Critical alarms, damage, threats, combat.
  * **Reset (`\033[0m`):** Normal prose.

### 4.3. State Management & Persistence (`pkg/game/state.go`)
* **`GameState` Struct:**
  ```go
  type GameState struct {
      Hull            int             `json:"hull"`
      Bat             int             `json:"bat"`
      Scrap           int             `json:"scrap"`
      CurrentChapter  string          `json:"current_chapter"`
      Inventory       []string        `json:"inventory"`
      Flags           map[string]bool `json:"flags"`
      TimeLockUntil   int64           `json:"time_lock_until"` // Unix timestamp
      LastUpdated     int64           `json:"last_updated"`
  }
  ```
* **Auto-Save:** Atomically save state to `.signal_loss_save.json` upon every chapter transition.
* **Auto-Load:** Load save file if present on startup; fallback to `story.json`'s `initial_state` if not found.

### 4.4. Real-World Time Locks
* When a chapter choice imposes a cooldown (`time_lock_seconds > 0`), set `TimeLockUntil = time.Now().Unix() + time_lock_seconds`.
* On subsequent boots or chapter checks, compare `time.Now().Unix()` against `TimeLockUntil`.
* If locked:
  * Display a low-power standby screen with remaining hours/minutes/seconds.
  * Prevent story progression until the duration has elapsed.

### 4.5. Data-Driven Story Architecture (`pkg/story/parser.go`)
* Ingest `story.json` into typed structs:
  * `StoryData`, `Chapter`, `Dialogue`, `Choice`, `StatMutations`.
* Process choice outcomes dynamically: apply stat deltas, add/remove items, set boolean flags, and update `current_chapter`.

### 4.6. Invalid Input Fallback
* If input does not match any valid choice ID:
  * Select a randomized sarcastic quip from `story.json`'s `invalid_input_responses`.
  * Print in ANSI Cyan/Red, pause briefly, and re-render the prompt.
