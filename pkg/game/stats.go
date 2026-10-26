package game

import (
	"encoding/json"
	"os"
)

// Stats outlive a single run: they are what makes New Game+ possible and
// what the end screen reports. Endings are keyed by Choice.EndingID, never
// by title, so a run in German and one in English count the same ending.
type Stats struct {
	Runs    int            `json:"runs"`
	Deaths  int            `json:"deaths"`
	Endings map[string]int `json:"endings"`
}

// NGPlusEndings is how many different endings unlock New Game+.
const NGPlusEndings = 3

// Record notes one finished run; an empty endingID means the run ended in death.
func (s *Stats) Record(endingID string) {
	s.Runs++
	if endingID == "" {
		s.Deaths++
		return
	}
	if s.Endings == nil {
		s.Endings = map[string]int{}
	}
	s.Endings[endingID]++
}

func (s *Stats) NGPlus() bool { return len(s.Endings) >= NGPlusEndings }

// LoadStats returns empty stats when the file is missing or unreadable.
func LoadStats(path string) *Stats {
	var s Stats
	if data, err := os.ReadFile(path); err == nil {
		json.Unmarshal(data, &s) //nolint:errcheck // a broken file just means starting over
	}
	return &s
}

func SaveStats(path string, s *Stats) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
