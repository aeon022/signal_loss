// Command wasm is the browser entry point for signal_loss: the exact same
// pkg/tui Bubble Tea program as the native CLI, bridged to JS instead of a
// real terminal. See ../../roadmap.md Phase 10 for how this fits together.
package main

import (
	"encoding/json"
	"syscall/js"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
	"github.com/aeon022/signal_loss/pkg/tui"
)

// terminalCols/terminalRows must match GameTerminal.astro's xterm.js
// dimensions exactly — see the WindowSizeMsg comment in slStart below.
const (
	terminalCols = 84
	terminalRows = 37
)

// bridgeReader feeds bytes pushed from JS (via slInput) to Bubble Tea's
// input loop, standing in for a real tty.
type bridgeReader struct {
	ch  chan []byte
	buf []byte
}

func (r *bridgeReader) Read(p []byte) (int, error) {
	if len(r.buf) == 0 {
		r.buf = <-r.ch
	}
	n := copy(p, r.buf)
	r.buf = r.buf[n:]
	return n, nil
}

// bridgeWriter hands Bubble Tea's rendered ANSI output straight to xterm.js
// via a JS callback — xterm.js already speaks ANSI/VT100, so no translation
// is needed on either side.
type bridgeWriter struct {
	onOutput js.Value
}

func (w *bridgeWriter) Write(p []byte) (int, error) {
	w.onOutput.Invoke(string(p))
	return len(p), nil
}

func main() {
	inputCh := make(chan []byte, 64)

	// slStart(storyJSON, savedStateJSON, fast, realtime, lang, onOutput, onSave, onExit)
	// lang only picks the UI chrome (headings, hints, boot log) — storyJSON
	// already carries whichever language's narrative content JS fetched,
	// so the two need to be kept in sync by the caller.
	js.Global().Set("slStart", js.FuncOf(func(_ js.Value, args []js.Value) any {
		storyJSON := args[0].String()
		savedJSON := args[1].String()
		fast := args[2].Bool()
		realtime := args[3].Bool()
		lang := args[4].String()
		onOutput := args[5]
		onSave := args[6]
		onExit := args[7]

		storyData, err := story.ParseStory([]byte(storyJSON))
		if err != nil {
			onOutput.Invoke("story error: " + err.Error() + "\r\n")
			return nil
		}

		state := loadOrNewState(storyData, savedJSON)

		save := func(s *game.GameState) error {
			data, err := json.Marshal(s)
			if err != nil {
				return err
			}
			onSave.Invoke(string(data))
			return nil
		}

		model := tui.New(storyData, state, save, fast, realtime, lang)
		prog := tea.NewProgram(model,
			tea.WithAltScreen(),
			tea.WithInput(&bridgeReader{ch: inputCh}),
			tea.WithOutput(&bridgeWriter{onOutput: onOutput}),
		)

		go func() {
			prog.Run() //nolint:errcheck // errors surface as terminal output, nothing to do with them here
			onExit.Invoke()
		}()

		// Bubble Tea normally learns the terminal size via a real tty's
		// ioctl, which the WASM build has none of (tty_js.go stubs that
		// path out entirely) — without this, the renderer's internal
		// height stays unset and its relative-cursor redraw math desyncs
		// against xterm.js's own scrolling: every frame lands one row
		// lower than the last instead of overwriting in place. Sent from
		// its own goroutine since Send() blocks until Run()'s loop is far
		// enough along to receive — calling it inline, before Run() even
		// starts, crashed the WASM module outright. Must match
		// terminalCols/terminalRows in GameTerminal.astro.
		go prog.Send(tea.WindowSizeMsg{Width: terminalCols, Height: terminalRows})

		return nil
	}))

	// slInput(bytes): raw terminal input from xterm.js (already encoded the
	// way a real terminal would send it — arrow keys as ESC sequences etc.)
	js.Global().Set("slInput", js.FuncOf(func(_ js.Value, args []js.Value) any {
		select {
		case inputCh <- []byte(args[0].String()):
		default:
		}
		return nil
	}))

	select {} // keep the module alive; work happens in callbacks above
}

func loadOrNewState(storyData *story.StoryData, savedJSON string) *game.GameState {
	if savedJSON != "" {
		var s game.GameState
		if err := json.Unmarshal([]byte(savedJSON), &s); err == nil {
			if idx := game.ChapterIndex(s.CurrentChapter); idx > s.MaxChapterIndex {
				s.MaxChapterIndex = idx
			}
			return &s
		}
	}
	return game.NewGameState(storyData)
}
