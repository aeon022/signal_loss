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

	// Spine, set only on a side chapter (e.g. "chapter_2b"), names the main-
	// line chapter it sits beside. The map shows the player there, and a
	// side chapter never moves the furthest-reached marker that gates the codex.
	Spine string `json:"spine"`
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
	EndingID        string          `json:"ending_id"` // language-independent key for stats; set exactly when Ending is
	Fatal           bool            `json:"fatal"`

	// RequiresFlag, if set, hides this choice unless that flag is already
	// true in the run's state — how an earlier decision unlocks a later
	// one (a callback, a shortcut, a way around a risk that choice itself
	// created). Empty means always available. See game.AvailableChoices.
	RequiresFlag string `json:"requires_flag"`

	// Reputation is added to the run's standing with factions (negative
	// lowers it); RequiresRep hides the choice until every named faction
	// has at least that much. Standing is hidden from the player, which is
	// why some options only ever appear to people who were nice earlier.
	Reputation  map[string]int `json:"reputation"`
	RequiresRep map[string]int `json:"requires_rep"`

	// RequiresItem hides the choice unless the item (by its name in this
	// language's story file) is in the inventory.
	RequiresItem string `json:"requires_item"`

	// A fatal choice with SavedByItem set is survivable by someone carrying
	// that item: the item is used up, SavedOutcome is shown and play goes on
	// at SavedNext instead of ending in death.
	SavedByItem  string `json:"saved_by_item"`
	SavedOutcome string `json:"saved_outcome"`
	SavedNext    string `json:"saved_next"`

	// Luck is a chance, per choice and per run, of an extra twist on top of
	// the normal outcome. Rolls come from the run's seed, so a given run is
	// reproducible; a negative seed (tests) switches luck off.
	Luck *Luck `json:"luck"`
}

// Luck is the optional twist of a choice: Chance percent of runs get the extra text and stat changes.
type Luck struct {
	Chance        int           `json:"chance"`
	Outcome       string        `json:"outcome"`
	StatMutations StatMutations `json:"stat_mutations"`
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
			for _, next := range []string{choice.NextChapter, choice.SavedNext} {
				if next == "" {
					continue
				}
				if _, ok := story.Chapters[next]; !ok {
					return nil, fmt.Errorf("invalid next_chapter reference %q in chapter %q", next, id)
				}
			}
		}
	}

	return &story, nil
}
