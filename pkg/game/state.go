package game

import (
	"encoding/json"
	"os"
)

// GameState is the full serializable state of a playthrough.
type GameState struct {
	Hull           int             `json:"hull"`
	Bat            int             `json:"bat"`
	Scrap          int             `json:"scrap"`
	CurrentChapter string          `json:"current_chapter"`
	Inventory      []string        `json:"inventory"`
	Flags          map[string]bool `json:"flags"`
	TimeLockUntil  int64           `json:"time_lock_until"`
	LastUpdated    int64           `json:"last_updated"`
	// MaxChapterIndex is the furthest point ever reached in ChapterOrder.
	// Unlike CurrentChapter it never regresses on a setback, so the codex
	// doesn't "forget" things a setback sent you back past.
	MaxChapterIndex int `json:"max_chapter_index"`
}

// ApplyMutations applies stat/inventory/flag changes from a chosen choice.
func (s *GameState) ApplyMutations(hull, bat, scrap int, itemsAdded, itemsRemoved []string, flags map[string]bool) {
	s.Hull += hull
	s.Bat += bat
	s.Scrap += scrap
	if s.Scrap < 0 {
		s.Scrap = 0
	}

	for _, item := range itemsAdded {
		s.Inventory = append(s.Inventory, item)
	}
	for _, item := range itemsRemoved {
		s.Inventory = removeItem(s.Inventory, item)
	}
	for k, v := range flags {
		s.Flags[k] = v
	}
}

func removeItem(items []string, target string) []string {
	out := items[:0]
	for _, item := range items {
		if item != target {
			out = append(out, item)
		}
	}
	return out
}

// SaveState writes the state to path as JSON.
func SaveState(path string, s *GameState) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadState reads a previously saved state from path.
func LoadState(path string) (*GameState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s GameState
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
