# signal_loss — Projektkontext

**Tech-Stack:** Go (`go.mod`: `module signal_loss`, go 1.27.1). Kein Node/JS/Python — nicht nach `package.json`/`README.md`/`*.js`/`*.ts` suchen, das ist hier irrelevant.

**Status:** Migration von einem alten Flat-Prototyp zu einer neuen `pkg/`-Architektur läuft, laut `roadmap.md` (Phase 1–5).

- **Alt/veraltet** (Root, `package main`): `game.go`, `player.go`, `item.go`, `enemy.go` — Prototyp, wird durch die neue Architektur ersetzt, aktuell nicht konsistent (referenziert einen nicht existierenden globalen `game`).
- **Neu, in Arbeit** (Zielarchitektur):
  - `cmd/signal_loss/main.go` — CLI-Entrypoint
  - `pkg/game/state.go` — `GameState` struct
  - `pkg/game/engine.go` — Engine-Loop (**aktuell kaputt/leer, Phase 4 unvollständig**)
  - `pkg/story/parser.go` — lädt `story.json`
  - `pkg/ui/terminal.go` — ANSI-Terminal-Rendering
- `story.json` — Kapitel-/Dialogdaten für das Spiel
- `roadmap.md` — Phasenplan, Quelle der Wahrheit für die Zielstruktur

**Vor jeder Aufgabe:** `go build ./...` laufen lassen, um den aktuellen Kompilierstatus zu sehen, bevor Rückfragen zum Tech-Stack gestellt werden.
