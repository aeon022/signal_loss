package story

import (
	"fmt"
	"testing"
)

func load(t *testing.T, name string) *StoryData {
	t.Helper()
	s, err := LoadStory("../../" + name)
	if err != nil {
		t.Fatalf("%s: %v", name, err)
	}
	return s
}

// mechanics is everything except prose (an ending counts as present or not; its title is prose): what must be identical across languages.
func mechanics(c Choice) string {
	return fmt.Sprintf("%+v|%v|%v|%v|%d|%s|%t|%t|%s", c.StatMutations, len(c.ItemsAdded), len(c.ItemsRemoved),
		c.Flags, c.TimeLockMinutes, c.NextChapter, c.Ending != "", c.Fatal, c.RequiresFlag)
}

func TestGermanMatchesEnglishMechanics(t *testing.T) {
	en, de := load(t, "story.json"), load(t, "story_de.json")
	if len(en.Chapters) != len(de.Chapters) {
		t.Fatalf("chapters: en %d, de %d", len(en.Chapters), len(de.Chapters))
	}
	if en.InitialState.Hull != de.InitialState.Hull || en.InitialState.Bat != de.InitialState.Bat || en.InitialState.Scrap != de.InitialState.Scrap || en.InitialState.CurrentChapter != de.InitialState.CurrentChapter {
		t.Errorf("initial stats differ: %+v vs %+v", en.InitialState, de.InitialState)
	}
	for id, ec := range en.Chapters {
		dc, ok := de.Chapters[id]
		if !ok {
			t.Errorf("%s missing in German", id)
			continue
		}
		if len(ec.Choices) != len(dc.Choices) {
			t.Errorf("%s: en %d choices, de %d", id, len(ec.Choices), len(dc.Choices))
			continue
		}
		for i := range ec.Choices {
			if a, b := mechanics(ec.Choices[i]), mechanics(dc.Choices[i]); a != b {
				t.Errorf("%s choice %d mechanics differ:\n en %s\n de %s", id, i, a, b)
			}
		}
	}
	if len(en.Codex) != len(de.Codex) {
		t.Errorf("codex: en %d, de %d", len(en.Codex), len(de.Codex))
	}
}

func TestStoryIsConsistent(t *testing.T) {
	for _, name := range []string{"story.json", "story_de.json"} {
		s := load(t, name)
		if _, ok := s.Chapters[s.InitialState.CurrentChapter]; !ok {
			t.Errorf("%s: start chapter %q missing", name, s.InitialState.CurrentChapter)
		}
		setFlags := map[string]bool{}
		for _, c := range s.Chapters {
			for _, ch := range c.Choices {
				for f := range ch.Flags {
					setFlags[f] = true
				}
			}
		}
		for id, c := range s.Chapters {
			if len(c.Choices) == 0 {
				t.Errorf("%s: %s has no choices", name, id)
			}
			for i, ch := range c.Choices {
				if ch.Text == "" || ch.Outcome == "" {
					t.Errorf("%s: %s choice %d has empty text/outcome", name, id, i)
				}
				if ch.RequiresFlag != "" && !setFlags[ch.RequiresFlag] {
					t.Errorf("%s: %s choice %d requires flag %q that nothing sets", name, id, i, ch.RequiresFlag)
				}
				if ch.NextChapter == "" && ch.Ending == "" && !ch.Fatal {
					t.Errorf("%s: %s choice %d leads nowhere (no next_chapter, ending or fatal)", name, id, i)
				}
			}
		}
		for term, e := range s.Codex {
			if _, ok := s.Chapters[e.UnlockChapter]; !ok {
				t.Errorf("%s: codex %q unlocks at unknown chapter %q", name, term, e.UnlockChapter)
			}
		}
	}
}
