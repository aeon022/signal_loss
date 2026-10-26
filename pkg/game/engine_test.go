package game

import (
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

// One real winning line per ending, found once by a search over every legal
// choice (hull, battery, flags and standing included) and pinned here. Choice
// texts are matched by prefix, so a rewrite of a line shows up as a failure.
func TestEveryEndingHasAWinningLine(t *testing.T) {
	spine := []string{`"Right, afternoon`, "Honest Barter", "Barricade the Hatch", "Debug the Nomad", "Disguise as a Pilgrim", "Discharge Capacitor"}
	finale := []string{"Trade the Hyper-Capacitor", "Straight through the Crater", "Standard Reset Code", "Blow the Airlock Hatch"}
	line := func(mid string, route ...string) []string {
		l := append(append([]string{}, spine...), "Patch the Leaks", mid)
		return append(append(l, route...), finale...)
	}
	lines := map[string][]string{
		"lucky_castaway": append(line("Take the Direct Road"), "Blind Full-Throttle"),
		"ace":            append(line("Take the Direct Road"), "Manual Course Correction"),
		"chaotic_hero":   append(line("Take the Direct Road"), "Overclock"),
		"loud_one":       append(line("Take the Direct Road"), "Broadcast Everything"),
		"apostle":        append(line("Join the Pilgrim Road", "Carry the Toaster Shrine", "Deliver a Sermon"), "Broadcast the Gospel"),
		"franchisee":     append(line("Hitch a Ride", "Work the Fryer", "Order the Combo Meal"), "Sign the Franchise"),
		"ascension_file": append(line("Slip Into the Glitch", "Join the Card Game", "Fix the Sky"), "Upload Yourself"),
		"loop_breaker":   {`"Right, afternoon`, "Covert Theft", "Run for It", "Open the Maintenance Hatch", "Walk Through the Exit"},
	}
	st := loadStory(t)
	for want, choices := range lines {
		s := NewGameState(st)
		s.Seed = -1               // no luck: these lines are pinned
		s.Flags["ng_plus"] = true // the hidden chapter is for people who already finished three runs
		var r Resolution
		for _, prefix := range choices {
			r = choose(t, st, s, prefix)
		}
		if r.EndingID != want {
			t.Errorf("line for %s ended with %q", want, r.EndingID)
		}
	}
}

func TestLuckIsReproducibleAndCanBeSwitchedOff(t *testing.T) {
	hits := 0
	for steps := 1; steps <= 1000; steps++ {
		if luckRoll(7, steps) != luckRoll(7, steps) {
			t.Fatal("roll is not reproducible")
		}
		if luckRoll(7, steps) < 30 {
			hits++
		}
	}
	if hits < 250 || hits > 350 {
		t.Errorf("30%% chance hit %d of 1000 times", hits)
	}
	st := &story.StoryData{Chapters: map[string]story.Chapter{}}
	s := &GameState{Hull: 10, Bat: 10, Seed: -1, Flags: map[string]bool{}}
	for range 50 {
		ResolveChoice(s, story.Choice{Luck: &story.Luck{Chance: 100, StatMutations: story.StatMutations{Hull: -1}}}, st)
	}
	if s.Hull != 10 {
		t.Errorf("luck fired with seed -1: hull %d", s.Hull)
	}
}

func TestItemSavesFromDeathOnce(t *testing.T) {
	st := loadStory(t)
	run := func(items ...string) (Resolution, *GameState) {
		s := NewGameState(st)
		s.Seed = -1
		s.CurrentChapter, s.Inventory = "chapter_10", items
		return choose(t, st, s, "Crowbar Bypass"), s
	}
	if r, _ := run(); !r.Fatal {
		t.Error("crowbar bypass should kill without the fuse")
	}
	r, s := run("Blessed Fuse")
	if r.Fatal || s.CurrentChapter != "chapter_11" || len(s.Inventory) != 0 {
		t.Errorf("fuse should save and be used up: fatal=%v chapter=%s items=%v", r.Fatal, s.CurrentChapter, s.Inventory)
	}
}

func TestItemGatedChoiceNeedsTheItem(t *testing.T) {
	st := loadStory(t)
	s := NewGameState(st)
	s.CurrentChapter = "chapter_10"
	count := func() int { return len(AvailableChoices(s, st.Chapters["chapter_10"])) }
	before := count()
	s.Inventory = []string{"Spare Fuse"}
	if count() != before+1 {
		t.Errorf("spare fuse should unlock one more choice: %d -> %d", before, count())
	}
}
