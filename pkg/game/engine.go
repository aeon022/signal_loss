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

// SpineIndex is the map position of a chapter: its place in ChapterOrder, or
// for a side chapter the place of the spine chapter it names. -1 if unknown.
func SpineIndex(st *story.StoryData, id string) int {
	if i := ChapterIndex(id); i >= 0 {
		return i
	}
	if ch, ok := st.Chapters[id]; ok {
		return ChapterIndex(ch.Spine)
	}
	return -1
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
		Rep:             map[string]int{},
		MaxChapterIndex: ChapterIndex(st.InitialState.CurrentChapter),
		Seed:            time.Now().UnixNano() & 0x7fffffff,
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

// AvailableChoices filters a chapter's choices down to the ones this run
// can actually see: everything with no RequiresFlag, plus anything gated
// on a flag or standing the player already has. Always use this instead of reading
// Chapter.Choices directly — an earlier decision unlocking a later option
// only works if every choice-list consumer (rendering, digit-key bounds,
// cursor bounds, selection) agrees on the same filtered list.
func AvailableChoices(state *GameState, ch story.Chapter) []story.Choice {
	available := make([]story.Choice, 0, len(ch.Choices))
	for _, c := range ch.Choices {
		if (c.RequiresFlag == "" || state.Flags[c.RequiresFlag]) && hasRep(state, c.RequiresRep) && (c.RequiresItem == "" || hasItem(state, c.RequiresItem)) {
			available = append(available, c)
		}
	}
	return available
}

func hasItem(state *GameState, item string) bool {
	for _, have := range state.Inventory {
		if have == item {
			return true
		}
	}
	return false
}

// luckRoll is a reproducible 0-99 for this run at this step (a splitmix-style hash).
func luckRoll(seed int64, steps int) int {
	x := uint64(seed)*0x9E3779B97F4A7C15 + uint64(steps)*0xBF58476D1CE4E5B9
	x = (x ^ x>>30) * 0xBF58476D1CE4E5B9
	x = (x ^ x>>27) * 0x94D049BB133111EB
	return int((x ^ x>>31) % 100)
}

// Risky reports whether picking c would kill this run: a fatal choice the
// player holds no saving item for. S.T.E.V.E. flags these once he has a route.
func Risky(state *GameState, c story.Choice) bool {
	return c.Fatal && !(c.SavedByItem != "" && hasItem(state, c.SavedByItem))
}

func hasRep(state *GameState, need map[string]int) bool {
	for faction, min := range need {
		if state.Rep[faction] < min {
			return false
		}
	}
	return true
}

// Resolution is what happens after a choice's mutations are applied: the
// run continues, time-locks, ends in death, or ends in one of the endings.
type Resolution struct {
	Outcome         string // the text to show: the choice's, its saved variant, plus any luck
	Fatal           bool
	GameOver        bool
	DeathMessage    string
	Ending          string
	EndingID        string
	TimeLockMinutes int
}

// ResolveChoice mutates state per the choice and reports what happens next.
// State.CurrentChapter is already updated by the time this returns. st is
// only needed for CheckGameOver's localized death messages.
func ResolveChoice(state *GameState, choice story.Choice, st *story.StoryData) Resolution {
	state.Steps++
	text := choice.Outcome
	if choice.Fatal && choice.SavedByItem != "" && hasItem(state, choice.SavedByItem) {
		state.Inventory = removeItem(state.Inventory, choice.SavedByItem)
		text, choice.Fatal, choice.NextChapter = choice.SavedOutcome, false, choice.SavedNext
	}

	state.ApplyMutations(
		choice.StatMutations.Hull,
		choice.StatMutations.Bat,
		choice.StatMutations.Scrap,
		choice.ItemsAdded,
		choice.ItemsRemoved,
		choice.Flags,
	)
	if state.Rep == nil { // saves from before reputation existed
		state.Rep = map[string]int{}
	}
	for faction, d := range choice.Reputation {
		state.Rep[faction] += d
	}
	if l := choice.Luck; l != nil && state.Seed >= 0 && luckRoll(state.Seed, state.Steps) < l.Chance {
		state.ApplyMutations(l.StatMutations.Hull, l.StatMutations.Bat, l.StatMutations.Scrap, nil, nil, nil)
		text += " " + l.Outcome
	}
	state.LastUpdated = time.Now().Unix()

	if choice.Fatal {
		return Resolution{Outcome: text, Fatal: true, DeathMessage: text}
	}

	if choice.NextChapter != "" {
		state.CurrentChapter = choice.NextChapter
		if idx := ChapterIndex(state.CurrentChapter); idx > state.MaxChapterIndex {
			state.MaxChapterIndex = idx
		}
	}

	if over, msg := CheckGameOver(state, st); over {
		return Resolution{Outcome: text, GameOver: true, DeathMessage: msg}
	}

	if choice.Ending != "" {
		return Resolution{Outcome: text, Ending: choice.Ending, EndingID: choice.EndingID}
	}

	return Resolution{Outcome: text, TimeLockMinutes: choice.TimeLockMinutes}
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
