package story

import (
	"encoding/json"
	"fmt"
	"os"
)

// StoryData is the full narrative dataset loaded from story.json.
type StoryData struct {
	InitialState          InitialState          `json:"initial_state"`
	Chapters              map[string]Chapter    `json:"chapters"`
	InvalidInputResponses []string              `json:"invalid_input_responses"`
	Codex                 map[string]CodexEntry `json:"codex"`
	SystemMessages        SystemMessages        `json:"system_messages"`
}

// SystemMessages are death lines triggered by stats alone (as opposed to a
// choice's own Outcome text) — localized here rather than hardcoded in
// pkg/game, since that package stays UI/language-agnostic. Empty fields
// fall back to an English default in pkg/game.CheckGameOver.
type SystemMessages struct {
	HullBreach      string `json:"hull_breach"`
	BatteryDepleted string `json:"battery_depleted"`
}

// CodexEntry is one lexicon term. UnlockChapter is the chapter id (e.g.
// "chapter_5") the player must have reached before it shows up in-game —
// the codex only reveals what the story has actually introduced so far.
type CodexEntry struct {
	Text          string `json:"text"`
	UnlockChapter string `json:"unlock_chapter"`
}

// InitialState seeds a fresh GameState for a new playthrough.
type InitialState struct {
	Hull           int      `json:"hull"`
	Bat            int      `json:"bat"`
	Scrap          int      `json:"scrap"`
	CurrentChapter string   `json:"current_chapter"`
	Inventory      []string `json:"inventory"`
}

// Chapter is one narrative beat with its available choices, in display order.
type Chapter struct {
	Title    string   `json:"title"`
	Location string   `json:"location"`
	Content  string   `json:"content"`
	Choices  []Choice `json:"choices"`
}

// Choice is one option a player can pick within a chapter.
type Choice struct {
	Text            string          `json:"text"`
	Outcome         string          `json:"outcome"`
	StatMutations   StatMutations   `json:"stat_mutations"`
	ItemsAdded      []string        `json:"items_added"`
	ItemsRemoved    []string        `json:"items_removed"`
	Flags           map[string]bool `json:"flags"`
	TimeLockMinutes int             `json:"time_lock_minutes"`
	NextChapter     string          `json:"next_chapter"`
	Ending          string          `json:"ending"`
	Fatal           bool            `json:"fatal"`
}

// StatMutations are the resource deltas a choice applies.
type StatMutations struct {
	Hull  int `json:"hull"`
	Bat   int `json:"bat"`
	Scrap int `json:"scrap"`
}

// LoadStory loads and validates the story from a JSON file on disk.
func LoadStory(path string) (*StoryData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open story file: %w", err)
	}
	return ParseStory(data)
}

// ParseStory validates and parses story JSON already in memory — the file
// isn't the only source: the WASM build fetches it over HTTP instead of
// reading it off a local disk.
func ParseStory(data []byte) (*StoryData, error) {
	var story StoryData
	if err := json.Unmarshal(data, &story); err != nil {
		return nil, fmt.Errorf("failed to parse story JSON: %w", err)
	}

	for id, chapter := range story.Chapters {
		for _, choice := range chapter.Choices {
			if choice.NextChapter == "" {
				continue
			}
			if _, ok := story.Chapters[choice.NextChapter]; !ok {
				return nil, fmt.Errorf("invalid next_chapter reference %q in chapter %q", choice.NextChapter, id)
			}
		}
	}

	return &story, nil
}
