package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
)

// The route stays hidden until a choice has found the satellite data, and the
// screen must fit the terminal exactly at every size we claim to support.
func TestRouteHiddenUntilSatelliteData(t *testing.T) {
	st, err := story.LoadStory("../../story.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, size := range [][2]int{{150, 42}, {100, 30}, {70, 26}, {60, 22}, {82, 20}, {40, 16}} {
		for _, known := range []bool{false, true} {
			s := game.NewGameState(st)
			s.Flags["sat_data"] = known
			m := New(st, s, func(*game.GameState) error { return nil }, &game.Stats{}, func(*game.Stats) error { return nil }, true, false, "en")
			next, _ := m.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
			out := next.(Model).View()
			if rows := strings.Count(out, "\n") + 1; rows != size[1] {
				t.Errorf("%v: %d rows rendered", size, rows)
			}
			if has := strings.Contains(out, "Scrap Bazaar"); has != (known && size[0] >= sideMinW) {
				t.Errorf("%v known=%v: route visible=%v", size, known, has)
			}
		}
	}
}

func TestEveryScreenFitsTheTerminal(t *testing.T) {
	st, err := story.LoadStory("../../story_de.json") // German runs longest
	if err != nil {
		t.Fatal(err)
	}
	long := strings.Repeat("Ein langer Satz über das Ende dieser Geschichte. ", 6)
	for _, size := range [][2]int{{150, 42}, {100, 30}, {70, 26}, {60, 22}, {82, 20}, {45, 18}, {40, 16}} {
		for _, phase := range []phase{phaseBoot, phaseGameOver, phaseEnding, phaseTimelock, phaseRealLock, phaseGlitch} {
			for _, overlay := range []string{"", "codex", "map", "help", "endings"} {
				if overlay != "" && phase != phaseChapter {
					continue
				}
				s := game.NewGameState(st)
				s.Flags["sat_data"] = true
				m := New(st, s, func(*game.GameState) error { return nil }, &game.Stats{Endings: map[string]int{"ace": 1}}, func(*game.Stats) error { return nil }, true, false, "de")
				m.phase, m.endMessage, m.endTitle, m.bootIdx, m.glitchLine = phase, long, "💀", 5, randomGlitchLine()
				m.overlay, m.awaitingUplink = overlay, true
				next, _ := m.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
				if rows := strings.Count(next.(Model).View(), "\n") + 1; rows != size[1] {
					t.Errorf("%v phase %d overlay %q: %d rows", size, phase, overlay, rows)
				}
			}
		}
	}
}

// The longest choice list (chapter 11 with everything unlocked) must fit even a tiny terminal.
func TestLongChoiceListFitsSmallTerminals(t *testing.T) {
	st, err := story.LoadStory("../../story_de.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, size := range [][2]int{{100, 30}, {82, 20}, {40, 16}} {
		s := game.NewGameState(st)
		s.CurrentChapter = "chapter_11"
		s.Flags = map[string]bool{"drones_inbound": true, "vex_grudge": true, "cartel_owns_you": true, "monks_forgave": true, "sat_data": true}
		s.Rep = map[string]int{"steve": 5}
		m := New(st, s, func(*game.GameState) error { return nil }, &game.Stats{}, func(*game.Stats) error { return nil }, true, false, "de")
		m.cursor = 6
		next, _ := m.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
		out := next.(Model).View()
		if rows := strings.Count(out, "\n") + 1; rows != size[1] {
			t.Errorf("%v: %d rows", size, rows)
		}
	}
}
