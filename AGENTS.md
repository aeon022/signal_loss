# signal_loss — Projektkontext

**Tech-Stack:** Go (Bubble Tea TUI), Modul `github.com/aeon022/signal_loss`. Kein Node im Spiel-Repo; die Landingpage liegt im Branch `deploy/landing`.

**Stand:** Spiel, Story (EN + DE, 12 Kapitel, 49 Entscheidungen, 4 Enden), Terminal-UI, Browser-Build (WASM) und Deployment sind fertig. Details: `stage.md` (Übergabestand), `roadmap.md` (Bau-Protokoll und Backlog), `story.md` (Story-Bibel).

**Aufbau:**
- `cmd/signal_loss/` CLI (`-fast`, `-realtime`, `-lang`, `-save`, `-story`); `cmd/wasm/` eigenes Modul für den Browser, dieselbe Engine
- `pkg/game/` reine Logik (State, Engine), `pkg/story/` JSON-Parser, `pkg/tui/` Oberfläche und alle UI-Texte (`labels.go`)
- `story.json` (EN) und `story_de.json` (DE): müssen mechanisch identisch sein; der Test `pkg/story/story_test.go` prüft das

**Regeln:**
- Vor jeder Aufgabe `go build ./... && go vet ./... && go test ./...`.
- Story-Änderungen immer in beiden Sprachdateien; `go test ./pkg/story/` muss grün bleiben.
- Live-Seite: `deploy/landing` baut bei jedem Push (oder `workflow_dispatch`) das Spiel aus `main`. Spieländerungen auf `main` erscheinen erst nach einem Deploy-Lauf im Browser.
- Terminal und Browser teilen sich Engine und Story; UI-Texte nur in `pkg/tui/labels.go` ändern, nicht doppelt pflegen.
