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
	return fmt.Sprintf("%+v|%v|%v|%v|%d|%s|%t|%t|%s|%v|%v|%s", c.StatMutations, len(c.ItemsAdded), len(c.ItemsRemoved),
		c.Flags, c.TimeLockMinutes, c.NextChapter, c.Ending != "", c.Fatal, c.RequiresFlag, c.Reputation, c.RequiresRep, c.EndingID)
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
				if ch.RequiresFlag != "" && !setFlags[ch.RequiresFlag] && ch.RequiresFlag != "ng_plus" { // ng_plus is set by the host from Stats
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

// Every chapter must be reachable from the start (ignoring flag gates), every
// ending must be reachable, and a side chapter's spine must be a real chapter.
func TestEverythingIsReachable(t *testing.T) {
	for _, name := range []string{"story.json", "story_de.json"} {
		s := load(t, name)
		seen := map[string]bool{}
		endings := 0
		queue := []string{s.InitialState.CurrentChapter}
		for len(queue) > 0 {
			id := queue[0]
			queue = queue[1:]
			if seen[id] {
				continue
			}
			seen[id] = true
			for _, ch := range s.Chapters[id].Choices {
				if ch.NextChapter != "" {
					queue = append(queue, ch.NextChapter)
				}
				if ch.Ending != "" {
					endings++
				}
			}
		}
		for id, c := range s.Chapters {
			if !seen[id] {
				t.Errorf("%s: %s is unreachable", name, id)
			}
			if c.Spine != "" {
				if _, ok := s.Chapters[c.Spine]; !ok {
					t.Errorf("%s: %s names unknown spine %q", name, id, c.Spine)
				}
			}
		}
		if endings == 0 {
			t.Errorf("%s: no ending reachable", name)
		}
	}
}

// Every requires_rep must be reachable by some route, or its choice (an ending, usually) is dead text. Best standing at a chapter is the max over incoming
// routes; len(chapters) passes cover any path without repeats, and a setback loop that pays reputation would only make this bound more generous.
func TestReputationGatesAreReachable(t *testing.T) {
	for _, name := range []string{"story.json", "story_de.json"} {
		s := load(t, name)
		best := map[string]map[string]int{s.InitialState.CurrentChapter: {}}
		for range s.Chapters {
			for id, rep := range best {
				for _, c := range s.Chapters[id].Choices {
					if c.NextChapter == "" {
						continue
					}
					next := map[string]int{}
					for f, v := range rep {
						next[f] = v
					}
					for f, d := range c.Reputation {
						next[f] += d
					}
					if old, ok := best[c.NextChapter]; ok {
						for f, v := range old {
							if v > next[f] {
								next[f] = v
							}
						}
					}
					best[c.NextChapter] = next
				}
			}
		}
		max := map[string]int{}
		for _, rep := range best {
			for f, v := range rep {
				if v > max[f] {
					max[f] = v
				}
			}
		}
		for id, ch := range s.Chapters {
			for i, c := range ch.Choices {
				for f, need := range c.RequiresRep {
					if max[f] < need {
						t.Errorf("%s: %s choice %d needs %s %d, best reachable is %d", name, id, i, f, need, max[f])
					}
				}
			}
		}
	}
}

func TestEndingsHaveUniqueIDs(t *testing.T) {
	for _, name := range []string{"story.json", "story_de.json"} {
		ids := map[string]bool{}
		for id, ch := range load(t, name).Chapters {
			for i, c := range ch.Choices {
				if (c.Ending == "") != (c.EndingID == "") {
					t.Errorf("%s: %s choice %d: ending and ending_id must be set together", name, id, i)
				}
				if c.EndingID != "" && ids[c.EndingID] {
					t.Errorf("%s: duplicate ending_id %q", name, c.EndingID)
				}
				ids[c.EndingID] = true
			}
		}
	}
}
