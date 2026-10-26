package game

import (
	"encoding/json"
	"testing"

	"github.com/aeon022/signal_loss/pkg/story"
)

func loadStory(t *testing.T) *story.StoryData {
	t.Helper()
	st, err := story.LoadStory("../../story.json")
	if err != nil {
		t.Fatal(err)
	}
	return st
}

// choose resolves the choice whose text starts with prefix in the state's current chapter.
func choose(t *testing.T, st *story.StoryData, s *GameState, prefix string) Resolution {
	t.Helper()
	for _, c := range AvailableChoices(s, st.Chapters[s.CurrentChapter]) {
		if len(c.Text) >= len(prefix) && c.Text[:len(prefix)] == prefix {
			return ResolveChoice(s, c, st)
		}
	}
	t.Fatalf("no choice starting %q in %s", prefix, s.CurrentChapter)
	return Resolution{}
}

func TestFailedTheftLeadsToLedgerNotChapterOne(t *testing.T) {
	st := loadStory(t)
	s := NewGameState(st)
	s.CurrentChapter, s.Scrap = "chapter_2", 3
	choose(t, st, s, "Covert Theft")
	if s.CurrentChapter != "chapter_2b" || s.Scrap != 0 || !s.Flags["wanted_by_cartel"] {
		t.Fatalf("after theft: chapter %s scrap %d flags %v", s.CurrentChapter, s.Scrap, s.Flags)
	}
	if got := s.MaxChapterIndex; got != 0 {
		t.Errorf("a side chapter must not move the codex marker, got %d", got)
	}
	if SpineIndex(st, "chapter_2b") != 1 {
		t.Errorf("map position of chapter_2b should be chapter 2")
	}
	choose(t, st, s, "Work the Grill")
	if s.CurrentChapter != "chapter_3" || s.Flags["cartel_debt"] != true || len(s.Inventory) != 1 || s.Hull != 11 {
		t.Fatalf("after work: chapter %s inv %v hull %d flags %v", s.CurrentChapter, s.Inventory, s.Hull, s.Flags)
	}
}

func TestCoolerIsFatalAndRunIsAWayBack(t *testing.T) {
	st := loadStory(t)
	s := NewGameState(st)
	s.CurrentChapter = "chapter_2b"
	if r := choose(t, st, s, "Raid the Cooler"); !r.Fatal {
		t.Error("the cooler should kill")
	}
	s = NewGameState(st)
	s.CurrentChapter = "chapter_2b"
	choose(t, st, s, "Run for It")
	if s.CurrentChapter != "chapter_1" {
		t.Errorf("run should still end at the wreck, got %s", s.CurrentChapter)
	}
}

// Reachability on paper (story tests) ignores gates and survival; this plays every
// legal line from a fresh run and collects the endings that someone can actually get.
// A state is skipped when an earlier visit had the same chapter, flags and (capped)
// standing with at least as much hull and battery: it could do everything this one can.
func TestEveryEndingIsWinnable(t *testing.T) {
	st := loadStory(t)
	type pair struct{ hull, bat int }
	seen := map[string][]pair{}
	endings := map[string]bool{}
	var walk func(s *GameState)
	walk = func(s *GameState) {
		rep := map[string]int{}
		for k, v := range s.Rep {
			rep[k] = min(v, 4) // no gate asks for more
		}
		key, _ := json.Marshal([]any{s.CurrentChapter, s.Flags, rep})
		for _, p := range seen[string(key)] {
			if p.hull >= s.Hull && p.bat >= s.Bat {
				return
			}
		}
		seen[string(key)] = append(seen[string(key)], pair{s.Hull, s.Bat})
		for _, c := range AvailableChoices(s, st.Chapters[s.CurrentChapter]) {
			next := *s
			next.Flags, next.Rep = map[string]bool{}, map[string]int{}
			for k, v := range s.Flags {
				next.Flags[k] = v
			}
			for k, v := range s.Rep {
				next.Rep[k] = v
			}
			next.Inventory = nil
			r := ResolveChoice(&next, c, st)
			switch {
			case r.Ending != "":
				endings[r.Ending] = true
			case !r.Fatal && !r.GameOver:
				walk(&next)
			}
		}
	}
	walk(NewGameState(st))
	if len(endings) != 7 {
		t.Errorf("want 7 reachable endings, got %d: %v", len(endings), endings)
	}
}
