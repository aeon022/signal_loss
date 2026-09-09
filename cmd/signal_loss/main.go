package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"signal_loss/pkg/game"
	"signal_loss/pkg/story"
	"signal_loss/pkg/tui"
)

func main() {
	fast := flag.Bool("fast", false, "skip typewriter effect and time-lock waits")
	realTime := flag.Bool("realtime", false, "make story time-locks real minutes/hours instead of a compressed few seconds — come back later, can't be skipped with -fast")
	savePath := flag.String("save", "save.json", "path to the save file")
	storyPath := flag.String("story", "story.json", "path to the story data file")
	flag.Parse()

	storyData, err := story.LoadStory(*storyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load story: %v\n", err)
		os.Exit(1)
	}

	state, err := game.LoadState(*savePath)
	if err != nil {
		state = game.NewGameState(storyData)
	} else if idx := game.ChapterIndex(state.CurrentChapter); idx > state.MaxChapterIndex {
		// Backfills max_chapter_index for saves written before it existed.
		state.MaxChapterIndex = idx
	}

	// ponytail: state is saved after every resolved choice (see
	// pkg/tui.selectChoice), so an abrupt kill/SIGTERM only loses the
	// current unconfirmed input, not run progress. No separate signal
	// handler needed on top of what Bubble Tea already restores.
	save := func(s *game.GameState) error { return game.SaveState(*savePath, s) }
	model := tui.New(storyData, state, save, *fast, *realTime)
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\ngame error: %v\n", err)
		os.Exit(1)
	}
}
