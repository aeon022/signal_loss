// Package tui is a Bubble Tea front end for the signal_loss game: a
// full-screen, retro-futuristic terminal UI with typewriter text reveal,
// a persistent HUD/keybar, and codex/map overlays.
package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
)

type phase int

const (
	phaseBoot phase = iota
	phaseChapter
	phaseGlitch
	phaseTimelock
	phaseRealLock
	phaseGameOver
	phaseEnding
)

type chapterStep int

const (
	stepNarrative chapterStep = iota
	stepOutcome
)

const (
	revealInterval = 12 * time.Millisecond
	bootInterval   = 450 * time.Millisecond
	timelockTick   = 100 * time.Millisecond
	glitchInterval = 45 * time.Millisecond
	glitchFrames   = 3
	codexPageSize  = 5

	// realLockTick is how often a pending real-time lock rechecks the
	// clock — both to redraw the live countdown and, if left running, to
	// auto-continue the moment real time actually passes.
	realLockTick = 1 * time.Second
)

type tickMsg struct{}

func tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return tickMsg{} })
}

// Model is the Bubble Tea model driving the whole game.
type Model struct {
	story    *story.StoryData
	state    *game.GameState
	save     func(*game.GameState) error
	fastMode bool
	realTime bool
	labels   Labels

	phase       phase
	overlay     string // "", "codex", "map"
	codexScroll int

	bootIdx        int
	awaitingUplink bool

	step             chapterStep
	revealText       []rune
	revealCount      int
	resolution       game.Resolution
	awaitingContinue bool
	cursor           int
	hint             string

	// outcomeTitle pins the frame title to the chapter a choice was made
	// in while its outcome is still revealing, since ResolveChoice already
	// advances state.CurrentChapter before the outcome text finishes.
	outcomeTitle string

	glitchLeft int
	glitchLine string

	timelockTotal     time.Duration
	timelockRemaining time.Duration

	endTitle   string
	endMessage string

	quitting bool
	fatalErr error
}

// New builds the initial model. Pass an already-loaded state (fresh or
// resumed), story data, and a save func — file-backed on the native CLI,
// localStorage-backed in the WASM build, the model doesn't care which.
func New(st *story.StoryData, state *game.GameState, save func(*game.GameState) error, fastMode, realTime bool, lang string) Model {
	m := Model{
		story:    st,
		state:    state,
		save:     save,
		fastMode: fastMode,
		realTime: realTime,
		labels:   labelsFor(lang),
		phase:    phaseBoot,
	}

	// A pending real-time lock from a previous session always wins, no
	// matter what flags this run was launched with — that's the whole
	// point: -fast can't be used to skip past it.
	if state.TimeLockUntil > 0 {
		if time.Now().Unix() < state.TimeLockUntil {
			m.phase = phaseRealLock
			return m
		}
		state.TimeLockUntil = 0
	}

	if fastMode {
		// Fast mode skips the boot screen entirely — no keypress needed.
		m.phase = phaseChapter
		m.startChapterReveal()
	}
	return m
}

func (m Model) Init() tea.Cmd {
	if m.phase == phaseRealLock {
		return tick(realLockTick)
	}
	if m.fastMode {
		return nil
	}
	return tick(bootInterval)
}

func (m Model) currentChapter() (story.Chapter, bool) {
	ch, ok := m.story.Chapters[m.state.CurrentChapter]
	return ch, ok
}

func (m *Model) startChapterReveal() {
	ch, ok := m.currentChapter()
	if !ok {
		m.fatalErr = errUnknownChapter(m.state.CurrentChapter)
		return
	}
	m.step = stepNarrative
	m.cursor = 0
	m.awaitingContinue = false
	m.hint = ""
	m.setReveal(ch.Content)
}

func (m *Model) setReveal(text string) {
	m.revealText = []rune(text)
	if m.fastMode {
		m.revealCount = len(m.revealText)
		return
	}
	m.revealCount = 0
}

func (m Model) revealDone() bool {
	return m.revealCount >= len(m.revealText)
}

type errUnknownChapter string

func (e errUnknownChapter) Error() string { return "unknown chapter: " + string(e) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case tickMsg:
		return m.handleTick()
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	if key == "ctrl+c" {
		m.quitting = true
		return m, tea.Quit
	}

	switch m.phase {
	case phaseBoot:
		return m.handleBootKey()
	case phaseChapter:
		return m.handleChapterKey(key)
	case phaseGlitch:
		return m, nil // ignore input mid-transition, it's brief
	case phaseTimelock:
		return m.skipTimelock()
	case phaseRealLock:
		if key == "q" {
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil
	case phaseGameOver, phaseEnding:
		m.quitting = true
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) handleBootKey() (tea.Model, tea.Cmd) {
	if !m.awaitingUplink {
		m.bootIdx = len(m.labels.BootLines)
		m.awaitingUplink = true
		return m, nil
	}
	m.phase = phaseChapter
	m.startChapterReveal()
	if m.fatalErr != nil {
		return m, tea.Quit
	}
	if m.fastMode {
		return m, nil
	}
	return m, tick(revealInterval)
}

func (m Model) handleChapterKey(key string) (tea.Model, tea.Cmd) {
	if m.overlay != "" {
		switch key {
		case "c", "m", "esc":
			m.overlay = ""
		case "up", "k":
			if m.codexScroll > 0 {
				m.codexScroll--
			}
		case "down", "j":
			m.codexScroll++ // clamped against actual content at render time
		case "pgup":
			m.codexScroll -= codexPageSize
			if m.codexScroll < 0 {
				m.codexScroll = 0
			}
		case "pgdown":
			m.codexScroll += codexPageSize
		}
		return m, nil
	}

	if !m.revealDone() {
		m.revealCount = len(m.revealText)
		return m, nil
	}

	if m.step == stepOutcome {
		if m.awaitingContinue {
			return m.applyResolution()
		}
		return m, nil
	}

	// stepNarrative, choices visible.
	ch, _ := m.currentChapter()
	choices := game.AvailableChoices(m.state, ch)
	switch key {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
		return m, nil
	case "down", "j":
		if m.cursor < len(choices)-1 {
			m.cursor++
		}
		return m, nil
	case "enter":
		return m.selectChoice(m.cursor)
	case "c":
		m.overlay = "codex"
		m.codexScroll = 0
		return m, nil
	case "m":
		m.overlay = "map"
		return m, nil
	case "q":
		m.quitting = true
		return m, tea.Quit
	}

	if idx, ok := digitIndex(key); ok && idx < len(choices) {
		return m.selectChoice(idx)
	}

	m.hint = game.RandomInvalidResponse(m.story)
	return m, nil
}

func digitIndex(key string) (int, bool) {
	if len(key) != 1 || key[0] < '1' || key[0] > '9' {
		return 0, false
	}
	return int(key[0] - '1'), true
}

func (m Model) selectChoice(idx int) (tea.Model, tea.Cmd) {
	ch, _ := m.currentChapter()
	choices := game.AvailableChoices(m.state, ch)
	if idx < 0 || idx >= len(choices) {
		m.hint = game.RandomInvalidResponse(m.story)
		return m, nil
	}
	choice := choices[idx]
	m.outcomeTitle = ch.Title
	m.resolution = game.ResolveChoice(m.state, choice, m.story)
	m.save(m.state)

	m.step = stepOutcome
	m.awaitingContinue = false
	m.hint = ""
	m.setReveal(choice.Outcome)

	if m.fastMode {
		return m.finishOutcome()
	}
	return m, tick(revealInterval)
}

// finishOutcome is called once the outcome text is fully revealed (or
// immediately, in fast mode). It never changes phase itself — whatever
// happens next (death, a time-lock, an ending, or just the next chapter)
// always waits for an explicit keypress first, so the screen never gets
// yanked away the instant the last character of an outcome lands.
func (m Model) finishOutcome() (tea.Model, tea.Cmd) {
	m.awaitingContinue = true
	return m, nil
}

// applyResolution acts on the choice's already-computed Resolution once
// the player confirms they've read the outcome (see finishOutcome above).
func (m Model) applyResolution() (tea.Model, tea.Cmd) {
	r := m.resolution
	switch {
	case r.Fatal || r.GameOver:
		m.phase = phaseGameOver
		m.endTitle = m.labels.GameOverTitle
		m.endMessage = r.DeathMessage

		return m, nil
	case r.TimeLockMinutes > 0 && m.realTime:
		m.state.TimeLockUntil = time.Now().Add(time.Duration(r.TimeLockMinutes) * time.Minute).Unix()
		m.save(m.state)
		m.phase = phaseRealLock
		return m, tick(realLockTick)
	case r.TimeLockMinutes > 0 && !m.fastMode:
		m.phase = phaseTimelock
		m.timelockTotal = time.Duration(r.TimeLockMinutes) * secondsPerGameMinute
		m.timelockRemaining = m.timelockTotal
		return m, tick(timelockTick)
	case r.Ending != "":
		m.phase = phaseEnding
		m.endMessage = r.Ending

		return m, nil
	default:
		return m.advanceChapter()
	}
}

// secondsPerGameMinute keeps the same compression ratio the original
// hours-based time-locks used (2 real seconds per game-hour = 8h -> 16s),
// now expressed per game-minute so shorter locks (15m, 30m, ...) scale
// down proportionally instead of needing their own constant.
const secondsPerGameMinute = 2 * time.Second / 60

func (m Model) skipTimelock() (tea.Model, tea.Cmd) {
	m.timelockRemaining = 0
	return m.afterTimelock()
}

func (m Model) afterTimelock() (tea.Model, tea.Cmd) {
	if m.resolution.Ending != "" {
		m.phase = phaseEnding
		m.endMessage = m.resolution.Ending

		return m, nil
	}
	m.phase = phaseGlitch
	m.glitchLeft = glitchFrames
	m.glitchLine = randomGlitchLine()
	return m, tick(glitchInterval)
}

func (m Model) advanceChapter() (tea.Model, tea.Cmd) {
	m.phase = phaseGlitch
	m.glitchLeft = glitchFrames
	m.glitchLine = randomGlitchLine()
	if m.fastMode {
		return m.finishGlitch()
	}
	return m, tick(glitchInterval)
}

func (m Model) finishGlitch() (tea.Model, tea.Cmd) {
	m.phase = phaseChapter
	m.startChapterReveal()
	if m.fatalErr != nil {
		return m, tea.Quit
	}
	if m.fastMode {
		return m, nil
	}
	return m, tick(revealInterval)
}

func (m Model) handleTick() (tea.Model, tea.Cmd) {
	switch m.phase {
	case phaseBoot:
		if m.bootIdx < len(m.labels.BootLines) {
			m.bootIdx++
			if m.bootIdx >= len(m.labels.BootLines) {
				m.awaitingUplink = true
				return m, nil
			}
			return m, tick(bootInterval)
		}
		return m, nil

	case phaseChapter:
		if m.overlay != "" {
			return m, nil
		}
		if !m.revealDone() {
			m.revealCount++
			if !m.revealDone() {
				return m, tick(revealInterval)
			}
		}
		if m.step == stepOutcome && !m.awaitingContinue {
			return m.finishOutcome()
		}
		return m, nil

	case phaseTimelock:
		m.timelockRemaining -= timelockTick
		if m.timelockRemaining <= 0 {
			return m.afterTimelock()
		}
		return m, tick(timelockTick)

	case phaseRealLock:
		if time.Now().Unix() >= m.state.TimeLockUntil {
			m.state.TimeLockUntil = 0
			m.save(m.state)
			return m.afterTimelock()
		}
		return m, tick(realLockTick)

	case phaseGlitch:
		m.glitchLeft--
		if m.glitchLeft <= 0 {
			return m.finishGlitch()
		}
		m.glitchLine = randomGlitchLine()
		return m, tick(glitchInterval)
	}
	return m, nil
}
