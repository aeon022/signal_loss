package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	signal_loss "github.com/aeon022/signal_loss"
	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
	"github.com/aeon022/signal_loss/pkg/tui"
)

func main() {
	fast := flag.Bool("fast", false, "skip typewriter effect and time-lock waits")
	realTime := flag.Bool("realtime", false, "make story time-locks real minutes/hours instead of a compressed few seconds — come back later, can't be skipped with -fast")
	savePath := flag.String("save", "save.json", "path to the save file")
	statsPath := flag.String("stats", "stats.json", "path to the stats file (endings found, deaths; unlocks New Game+)")
	lang := flag.String("lang", "", "game language: en or de (asked at the start of a new run if omitted)")
	storyPath := flag.String("story", "", "path to the story data file (default: built into the binary)")
	flag.Parse()

	// A run keeps the language it started in (item names differ per language),
	// so ask only when starting fresh and no -lang was given.
	state, loadErr := game.LoadState(*savePath)
	if *lang == "" {
		switch {
		case loadErr == nil && state.Lang != "":
			*lang = state.Lang
		case loadErr == nil:
			*lang = "en" // saves from before the language was recorded
		default:
			*lang = askLanguage()
		}
	}

	var storyData *story.StoryData
	var err error
	if *storyPath != "" {
		storyData, err = story.LoadStory(*storyPath)
	} else {
		name := "story.json"
		if *lang == "de" {
			name = "story_de.json"
		}
		var data []byte
		if data, err = signal_loss.Stories.ReadFile(name); err == nil {
			storyData, err = story.ParseStory(data)
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load story: %v\n", err)
		os.Exit(1)
	}

	if loadErr != nil {
		state = game.NewGameState(storyData)
	} else if idx := game.ChapterIndex(state.CurrentChapter); idx > state.MaxChapterIndex {
		// Backfills max_chapter_index for saves written before it existed.
		state.MaxChapterIndex = idx
	}
	state.Lang = *lang

	// ponytail: state is saved after every resolved choice (see
	// pkg/tui.selectChoice), so an abrupt kill/SIGTERM only loses the
	// current unconfirmed input, not run progress. No separate signal
	// handler needed on top of what Bubble Tea already restores.
	save := func(s *game.GameState) error { return game.SaveState(*savePath, s) }
	stats := game.LoadStats(*statsPath)
	saveStats := func(s *game.Stats) error { return game.SaveStats(*statsPath, s) }
	model := tui.New(storyData, state, save, stats, saveStats, *fast, *realTime, *lang)
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\ngame error: %v\n", err)
		os.Exit(1)
	}
}

func askLanguage() string {
	fmt.Print("Language / Sprache:  [1] English   [2] Deutsch  (Enter = English) > ")
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	switch strings.ToLower(strings.TrimSpace(line)) {
	case "2", "d", "de", "deutsch":
		return "de"
	}
	return "en"
}
