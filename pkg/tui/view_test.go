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
	for _, size := range [][2]int{{150, 42}, {100, 30}, {70, 26}, {60, 22}} {
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
