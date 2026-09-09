package tui

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"signal_loss/pkg/game"
	"signal_loss/pkg/story"
)

const glitchChars = "▓▒░#%&$"

func randomGlitchLine() string {
	b := make([]byte, TextWidth)
	for i := range b {
		b[i] = glitchChars[rand.Intn(len(glitchChars))]
	}
	return string(b)
}

func heading(text string, c lipgloss.Color) string {
	return lipgloss.NewStyle().Bold(true).Foreground(c).Render(text)
}

// hudLine is the instrument-panel readout shown below the screen on every
// in-game view. Empty before there's a meaningful run to report on.
func (m Model) hudLine() string {
	return statusBar(m.state.Hull, m.state.Bat, m.state.Scrap, len(m.state.Inventory))
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.fatalErr != nil {
		return chrome(frameStyle(colorRed).Render("ERROR: "+m.fatalErr.Error()), "")
	}

	switch m.phase {
	case phaseBoot:
		return m.viewBoot()
	case phaseChapter:
		return m.viewChapter()
	case phaseGlitch:
		glitch := dimStyle.Render(m.glitchLine) + "\n" + dimStyle.Render(m.glitchLine)
		return chrome(frameStyle(colorDim).Render(glitch), m.hudLine())
	case phaseTimelock:
		return m.viewTimelock()
	case phaseRealLock:
		return m.viewRealLock()
	case phaseGameOver:
		return m.viewEnd(colorRed)
	case phaseEnding:
		return m.viewEnd(colorMagenta)
	}
	return ""
}

// viewBoot renders the title banner and boot log as one fixed-size panel.
func (m Model) viewBoot() string {
	lines := []string{
		heading("S I G N A L   L O S S", colorCyan),
		"// SECTOR SCHRÖDINGER //",
		"CSS RUST-404 — EMERGENCY TERMINAL LINK",
		"",
	}
	for i := 0; i < m.bootIdx; i++ {
		lines = append(lines, dimStyle.Render(bootLines[i]))
	}
	if m.awaitingUplink {
		lines = append(lines, "", dimStyle.Render("Press any key to establish uplink..."))
	}
	return chrome(frameStyle(colorCyan).Render(strings.Join(lines, "\n")), "")
}

// viewChapter renders the title and narrative/outcome text up top, with the
// choices (or continue prompt) anchored near the bottom of the fixed-height
// screen — a consistent "command line" position instead of floating right
// after however much narrative text precedes it. HULL/BAT/SCRAP/INV lives
// below the screen now, on the instrument panel (see hudLine).
func (m Model) viewChapter() string {
	ch, ok := m.currentChapter()
	if !ok {
		return chrome(frameStyle(colorRed).Render("ERROR: unknown chapter"), "")
	}
	if m.overlay == "codex" {
		return m.viewCodex()
	}
	if m.overlay == "map" {
		return m.viewMap()
	}

	title := ch.Title
	if m.step == stepOutcome {
		title = m.outcomeTitle
	}

	top := []string{heading(title, colorCyan), ""}
	top = append(top, wrap(string(m.revealText[:m.revealCount]), TextWidth)...)

	var bottom []string
	switch {
	case m.step == stepOutcome:
		if m.awaitingContinue {
			bottom = []string{"", dimStyle.Render("Press any key to continue...")}
		}
	case !m.revealDone():
		// still typing the narrative — choices appear once it's done.
	default:
		bottom = append(bottom, "")
		bottom = append(bottom, renderChoices(ch.Choices, m.cursor, TextWidth)...)
		if m.hint != "" {
			bottom = append(bottom, "", dimStyle.Render(m.hint))
		}
		bottom = append(bottom, "", dimStyle.Render(fmt.Sprintf(
			"[1-%d] choose   ↑↓ + enter   [c]odex   [m]ap   [q]uit", len(ch.Choices))))
	}

	lines := anchorBottom(top, bottom)
	return chrome(frameStyle(colorCyan).Render(strings.Join(lines, "\n")), m.hudLine())
}

// renderChoices word-wraps each choice under a hanging indent and
// highlights the cursor-selected one. Returns individual rows, not a
// single joined block, so callers can count/pad them line-for-line.
func renderChoices(choices []story.Choice, cursor, width int) []string {
	var rows []string
	for i, c := range choices {
		marker, style := "▸", choiceStyle
		if i == cursor {
			marker, style = "▶", cursorStyle
		}
		prefix := fmt.Sprintf("[%d] ", i+1)
		indent := strings.Repeat(" ", len(prefix)+2)
		avail := width - len(indent)
		if avail < 10 {
			avail = 10
		}
		wrapped := wrap(c.Text, avail)
		block := make([]string, len(wrapped))
		for j, ln := range wrapped {
			if j == 0 {
				block[j] = marker + " " + prefix + ln
			} else {
				block[j] = indent + ln
			}
		}
		styled := style.Render(strings.Join(block, "\n"))
		rows = append(rows, strings.Split(styled, "\n")...)
	}
	return rows
}

// codexOverhead is the heading, its blank line, and the footer row —
// everything in the codex frame that isn't part of the scrollable body.
const codexOverhead = 3

// viewCodex shows only the lexicon entries the story has actually
// introduced so far (by furthest chapter reached), not the full database —
// S.T.E.V.E. hasn't got intel on Vex before you've met him.
func (m Model) viewCodex() string {
	terms := make([]string, 0, len(m.story.Codex))
	for term, entry := range m.story.Codex {
		if game.ChapterIndex(entry.UnlockChapter) <= m.state.MaxChapterIndex {
			terms = append(terms, term)
		}
	}
	sort.Strings(terms)

	var body []string
	for _, term := range terms {
		body = append(body, boldStyle.Render(term))
		body = append(body, wrap(m.story.Codex[term].Text, TextWidth)...)
		body = append(body, "")
	}
	if len(terms) == 0 {
		body = []string{dimStyle.Render("No data on file yet. Keep moving.")}
	}

	visibleRows := FixedHeight - codexOverhead
	maxScroll := len(body) - visibleRows
	if maxScroll < 0 {
		maxScroll = 0
	}
	scroll := m.codexScroll
	if scroll > maxScroll {
		scroll = maxScroll
	}

	end := scroll + visibleRows
	if end > len(body) {
		end = len(body)
	}
	visible := append([]string{}, body[scroll:end]...)
	for len(visible) < visibleRows {
		visible = append(visible, "")
	}

	footer := fmt.Sprintf("%d known — [c/esc] back to chapter", len(terms))
	if maxScroll > 0 {
		footer = fmt.Sprintf("[↑↓ scroll %d/%d]  %s", scroll+1, maxScroll+1, footer)
	}

	lines := append([]string{heading("S.T.E.V.E. DATABASE — CODEX", colorAmber), ""}, visible...)
	lines = append(lines, dimStyle.Render(footer))
	return chrome(frameStyle(colorAmber).Render(strings.Join(lines, "\n")), m.hudLine())
}

func (m Model) viewMap() string {
	currentIdx := -1
	for i, id := range game.ChapterOrder {
		if id == m.state.CurrentChapter {
			currentIdx = i
			break
		}
	}

	lines := []string{heading("NAVIGATION — SECTOR SCHRÖDINGER", colorCyan), ""}
	for i, id := range game.ChapterOrder {
		ch, ok := m.story.Chapters[id]
		if !ok {
			continue
		}
		name := ch.Location
		if name == "" {
			name = ch.Title
		}

		var mark string
		switch {
		case i == currentIdx:
			mark = lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render("▶")
		case i < currentIdx || currentIdx == -1:
			mark = dimStyle.Render("✓")
		default:
			mark = "·"
		}
		lines = append(lines, fmt.Sprintf("%s %2d. %s", mark, i+1, name))
	}
	lines = append(lines, "", dimStyle.Render("[m/esc] back to chapter"))
	return chrome(frameStyle(colorCyan).Render(strings.Join(lines, "\n")), m.hudLine())
}

func (m Model) viewTimelock() string {
	pct := 0
	if m.timelockTotal > 0 {
		pct = 100 - int(m.timelockRemaining*100/m.timelockTotal)
	}
	lines := []string{
		dimStyle.Render("SIGNAL LOST — standing by..."),
		"",
		bar(pct, TextWidth-8),
		"",
		dimStyle.Render("[any key] skip wait"),
	}
	return chrome(frameStyle(colorDim).Render(strings.Join(lines, "\n")), m.hudLine())
}

// viewRealLock shows a genuine real-time wait (-realtime): the run is
// saved and it's safe to close the terminal — S.T.E.V.E. will still be
// there when the clock catches up, whether that's from a fresh launch or
// this same session left running.
func (m Model) viewRealLock() string {
	remaining := time.Until(time.Unix(m.state.TimeLockUntil, 0))
	if remaining < 0 {
		remaining = 0
	}
	unlockAt := time.Unix(m.state.TimeLockUntil, 0).Local().Format("Mon 15:04")

	lines := []string{heading("SIGNAL LOST — TRANSMISSION PAUSED", colorDim), ""}
	lines = append(lines, wrap("S.T.E.V.E. has gone into standby to conserve what's left of the "+
		"battery. Reconnect in "+formatDuration(remaining)+".", TextWidth)...)
	lines = append(lines,
		"",
		dimStyle.Render("Back online: "+unlockAt),
		"",
		dimStyle.Render("Your run is saved — safe to close this window and come back."),
		"",
		dimStyle.Render("[q]uit"),
	)
	return chrome(frameStyle(colorDim).Render(strings.Join(lines, "\n")), m.hudLine())
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	mn := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	switch {
	case h > 0:
		return fmt.Sprintf("%dh %dm", h, mn)
	case mn > 0:
		return fmt.Sprintf("%dm %ds", mn, s)
	default:
		return fmt.Sprintf("%ds", s)
	}
}

func (m Model) viewEnd(c lipgloss.Color) string {
	var lines []string
	if m.endTitle != "" {
		lines = append(lines, heading(m.endTitle, c), "")
	}
	lines = append(lines, wrap(m.endMessage, TextWidth)...)
	lines = append(lines, "", dimStyle.Render("[any key] quit"))
	return chrome(frameStyle(c).Render(strings.Join(lines, "\n")), m.hudLine())
}
