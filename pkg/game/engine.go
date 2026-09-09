package game

import (
	"math/rand"
	"time"

	"github.com/aeon022/signal_loss/pkg/story"
)

// ChapterOrder is the main spine, used to draw the map. Story branches can
// send the player back to an earlier chapter (a "setback"), but the spine
// itself is a fixed sequence.
var ChapterOrder = []string{
	"chapter_1", "chapter_2", "chapter_3", "chapter_4", "chapter_5", "chapter_6",
	"chapter_7", "chapter_8", "chapter_9", "chapter_10", "chapter_11", "chapter_12",
}

// NewGameState builds a fresh GameState from the story's initial_state block.
func NewGameState(st *story.StoryData) *GameState {
	return &GameState{
		Hull:            st.InitialState.Hull,
		Bat:             st.InitialState.Bat,
		Scrap:           st.InitialState.Scrap,
		CurrentChapter:  st.InitialState.CurrentChapter,
		Inventory:       append([]string{}, st.InitialState.Inventory...),
		Flags:           map[string]bool{},
		MaxChapterIndex: ChapterIndex(st.InitialState.CurrentChapter),
	}
}

// ChapterIndex returns id's position in ChapterOrder, or -1 if unknown.
func ChapterIndex(id string) int {
	for i, c := range ChapterOrder {
		if c == id {
			return i
		}
	}
	return -1
}

// Resolution is what happens after a choice's mutations are applied: the
// run continues, time-locks, ends in death, or ends in one of the endings.
type Resolution struct {
	Fatal           bool
	GameOver        bool
	DeathMessage    string
	Ending          string
	TimeLockMinutes int
}

// ResolveChoice mutates state per the choice and reports what happens next.
// State.CurrentChapter is already updated by the time this returns. st is
// only needed for CheckGameOver's localized death messages.
func ResolveChoice(state *GameState, choice story.Choice, st *story.StoryData) Resolution {
	state.ApplyMutations(
		choice.StatMutations.Hull,
		choice.StatMutations.Bat,
		choice.StatMutations.Scrap,
		choice.ItemsAdded,
		choice.ItemsRemoved,
		choice.Flags,
	)
	state.LastUpdated = time.Now().Unix()

	if choice.Fatal {
		return Resolution{Fatal: true, DeathMessage: choice.Outcome}
	}

	if choice.NextChapter != "" {
		state.CurrentChapter = choice.NextChapter
		if idx := ChapterIndex(state.CurrentChapter); idx > state.MaxChapterIndex {
			state.MaxChapterIndex = idx
		}
	}

	if over, msg := CheckGameOver(state, st); over {
		return Resolution{GameOver: true, DeathMessage: msg}
	}

	if choice.Ending != "" {
		return Resolution{Ending: choice.Ending}
	}

	return Resolution{TimeLockMinutes: choice.TimeLockMinutes}
}

// CheckGameOver reports whether the run has ended in death from the stats
// alone (as opposed to a choice explicitly marked Fatal). Messages come
// from st.SystemMessages so they follow the story file's own language;
// the fallback here only fires for a story file that omits that section.
func CheckGameOver(s *GameState, st *story.StoryData) (bool, string) {
	switch {
	case s.Hull <= 0:
		if msg := st.SystemMessages.HullBreach; msg != "" {
			return true, msg
		}
		return true, "HULL BREACH. Decompression is instant. Sector Schrödinger claims another cargo engineer."
	case s.Bat <= 0:
		if msg := st.SystemMessages.BatteryDepleted; msg != "" {
			return true, msg
		}
		return true, "BAT depleted. S.T.E.V.E. shuts down. Life support fails in the dark."
	default:
		return false, ""
	}
}

// RandomInvalidResponse picks a sarcastic fallback line for bad input.
func RandomInvalidResponse(st *story.StoryData) string {
	responses := st.InvalidInputResponses
	if len(responses) == 0 {
		return "S.T.E.V.E.: That's not a valid option, genius."
	}
	return responses[rand.Intn(len(responses))]
}
