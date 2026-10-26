package tui

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
)

const glitchChars = "▓▒░█▌▐▄▀#%&$"

// randomGlitchLine is a long run of noise; the glitch view cuts rows out of it.
func randomGlitchLine() string {
	r := []rune(glitchChars)
	b := make([]rune, 700)
	for i := range b {
		b[i] = r[rand.Intn(len(r))]
	}
	return string(b)
}

func heading(text string, c lipgloss.Color) string {
	return lipgloss.NewStyle().Bold(true).Foreground(c).Render(text)
}

func trunc(s string, w int) string { return lipgloss.NewStyle().MaxWidth(w).Render(s) }

func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.fatalErr != nil {
		return "ERROR: " + m.fatalErr.Error()
	}
	l := m.layout()
	if m.width > 0 && l.tooSmall() {
		return fmt.Sprintf(m.labels.TooSmallFmt, minW, minH, m.width, m.height)
	}

	switch m.phase {
	case phaseBoot:
		return m.viewBoot(l.full())
	case phaseChapter:
		return m.viewChapter(l)
	case phaseGlitch:
		return m.viewGlitch(l)
	case phaseTimelock:
		return m.viewTimelock(l.full())
	case phaseRealLock:
		return m.viewRealLock(l.full())
	case phaseGameOver:
		return m.viewEnd(l.full(), colorRed)
	case phaseEnding:
		return m.viewEnd(l.full(), colorMagenta)
	}
	return ""
}

// screen assembles the whole display: title bar, main panel (plus the
// instrument sidebar when there is room) and the key bar.
func (m Model) screen(l layout, c lipgloss.Color, body []string, footer string) string {
	if !l.side && (m.phase == phaseChapter || m.phase == phaseGlitch) {
		body = append([]string{statusBar(m.state.Hull, m.state.Bat, m.state.Scrap, len(m.state.Inventory)), ""}, body...)
	}
	for i := range body {
		body[i] = l.indent + body[i]
	}
	row := panel(c, l.mainW, l.bodyH, 2, body)
	if l.side {
		row = lipgloss.JoinHorizontal(lipgloss.Top, row, m.sidebar(l))
	}
	signal := "SIGNAL: LIVE"
	if m.phase == phaseGameOver {
		signal = "SIGNAL: LOST"
	}
	right := ""
	if idx := game.SpineIndex(m.story, m.state.CurrentChapter); idx >= 0 && m.phase != phaseBoot {
		right = fmt.Sprintf("CH %02d ", idx+1)
		if m.state.Flags["sat_data"] {
			right = fmt.Sprintf("CH %02d/%02d ", idx+1, len(game.ChapterOrder))
		}
	}
	head := bar(colorCyan, l.w, " ◢◤ S.T.E.V.E. TERMINAL // MK.IV", "SECTOR SCHRÖDINGER · "+signal+" ")
	foot := bar(colorDim, l.w, " "+footer, right)
	return head + "\n" + row + "\n" + foot
}

// sidebar is the instrument panel: gauges, cargo and the route so far.
func (m Model) sidebar(l layout) string {
	w := l.sideTxt
	s := m.state
	lines := []string{heading("◈ SUIT STATUS", colorCyan), "",
		sideRow("HULL", s.Hull, w), sideRow("BAT", s.Bat, w),
		boldStyle.Render(pad("SCRAP", 6)) + strconv.Itoa(s.Scrap), "",
		heading("◈ CARGO", colorCyan)}
	if len(s.Inventory) == 0 {
		lines = append(lines, dimStyle.Render("  — empty —"))
	}
	for i, item := range s.Inventory {
		if i == 5 {
			lines = append(lines, dimStyle.Render(fmt.Sprintf("  … +%d more", len(s.Inventory)-5)))
			break
		}
		lines = append(lines, trunc("  ◦ "+item, w))
	}

	cur := game.SpineIndex(m.story, s.CurrentChapter)
	lines = append(lines, "", heading("◈ ROUTE", colorCyan))
	if !s.Flags["sat_data"] {
		lines = append(lines, dimStyle.Render("  "+m.labels.RouteUnknown))
	} else if l.bodyH-len(lines) >= len(game.ChapterOrder)+1 {
		for i, id := range game.ChapterOrder {
			name := m.story.Chapters[id].Location
			mark := "·"
			switch {
			case i == cur:
				mark = lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render("▶")
			case i < cur:
				mark = dimStyle.Render("✓")
			}
			lines = append(lines, trunc(fmt.Sprintf(" %s %2d %s", mark, i+1, name), w))
		}
	} else {
		lines = append(lines, fmt.Sprintf(" %d/%d ", cur+1, len(game.ChapterOrder))+gauge((cur+1)*100/len(game.ChapterOrder), w-8))
	}
	return panel(colorDim, l.sideW, l.bodyH, 1, lines)
}

var logoSignal = []string{
	"███████╗██╗ ██████╗ ███╗   ██╗ █████╗ ██╗     ",
	"██╔════╝██║██╔════╝ ████╗  ██║██╔══██╗██║     ",
	"███████╗██║██║  ███╗██╔██╗ ██║███████║██║     ",
	"╚════██║██║██║   ██║██║╚██╗██║██╔══██║██║     ",
	"███████║██║╚██████╔╝██║ ╚████║██║  ██║███████╗",
	"╚══════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝",
}

var logoLoss = []string{
	"██╗      ██████╗ ███████╗███████╗",
	"██║     ██╔═══██╗██╔════╝██╔════╝",
	"██║     ██║   ██║███████╗███████╗",
	"██║     ██║   ██║╚════██║╚════██║",
	"███████╗╚██████╔╝███████║███████║",
	"╚══════╝ ╚═════╝ ╚══════╝╚══════╝",
}

// logo is the title in big letters where they fit, plain text where not.
func logo(w, rows int) []string {
	var out []string
	switch {
	case w >= 84 && rows >= 16:
		for i := range logoSignal {
			out = append(out, lipgloss.NewStyle().Foreground(colorCyan).Render(logoSignal[i])+"  "+lipgloss.NewStyle().Foreground(colorMagenta).Render(logoLoss[i]))
		}
	case w >= 50 && rows >= 24:
		for _, part := range []struct {
			art []string
			c   lipgloss.Color
		}{{logoSignal, colorCyan}, {logoLoss, colorMagenta}} {
			for _, ln := range part.art {
				out = append(out, lipgloss.NewStyle().Foreground(part.c).Render(ln))
			}
		}
	default:
		out = []string{heading("S I G N A L   L O S S", colorCyan)}
	}
	return out
}

func (m Model) viewBoot(l layout) string {
	log := []string{" "}
	for i := 0; i < m.bootIdx; i++ {
		log = append(log, dimStyle.Render(m.labels.BootLines[i]))
	}
	prompt := []string{" "}
	if m.awaitingUplink {
		prompt = []string{lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render(m.labels.PressUplink)}
	}
	return m.screen(l, colorCyan, centered(l, logo(l.textW, l.rows-12), []string{"", m.labels.BootSubtitle}, []string{m.labels.BootLink}, log, []string{""}, prompt), "")
}

// centered places a block in the middle of the main panel, line by line.
// Each part is kept left-aligned as a unit and the units are centered, so
// lines of different length (a boot log) don't drift apart.
func centered(l layout, parts ...[]string) []string {
	units := make([]string, len(parts))
	for i, p := range parts {
		units[i] = lipgloss.JoinVertical(lipgloss.Left, p...)
	}
	return strings.Split(lipgloss.Place(l.textW, l.rows, lipgloss.Center, lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center, units...)), "\n")
}

func (m Model) viewGlitch(l layout) string {
	noise := []rune(m.glitchLine)
	body := make([]string, l.rows)
	for i := range body {
		off := (i * 37) % (len(noise) - l.textW)
		body[i] = dimStyle.Render(string(noise[off : off+l.textW]))
	}
	return m.screen(l, colorDim, body, "")
}

// viewChapter renders the title and narrative/outcome text up top, with the
// choices (or continue prompt) anchored at the bottom of the panel — a
// consistent "command line" position however much text precedes it.
func (m Model) viewChapter(l layout) string {
	ch, ok := m.currentChapter()
	if !ok {
		return m.screen(l, colorRed, []string{m.labels.ErrUnknownChap}, "")
	}
	switch m.overlay {
	case "codex":
		return m.viewCodex(l)
	case "map":
		return m.viewMap(l)
	case "help":
		return m.viewHelp(l)
	}

	title := ch.Title
	if m.step == stepOutcome {
		title = m.outcomeTitle
	}

	top := []string{heading(title, colorCyan), ""}
	top = append(top, wrap(string(m.revealText[:m.revealCount]), l.textW)...)

	var bottom []string
	footer := ""
	switch {
	case m.step == stepOutcome:
		if m.awaitingContinue {
			bottom = []string{"", lipgloss.NewStyle().Foreground(colorGreen).Render(m.labels.PressContinue)}
		}
	case !m.revealDone():
		// still typing the narrative — choices appear once it's done.
	default:
		choices := game.AvailableChoices(m.state, ch)
		bottom = append([]string{""}, renderChoices(choices, m.cursor, l.textW)...)
		if m.hint != "" {
			bottom = append(bottom, "", dimStyle.Render(m.hint))
		}
		footer = fmt.Sprintf(m.labels.ChoiceFooterFmt, len(choices))
		if n := m.unseenCodex(); n > 0 {
			footer += fmt.Sprintf(m.labels.CodexNewFmt, n)
		}
	}

	top = window(top, l.rows-len(bottom), m.textScroll, !m.revealDone())
	return m.screen(l, colorCyan, anchorBottom(top, bottom, l.rows), footer)
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
		wrapped := wrap(c.Text, max(width-len(indent), 10))
		block := make([]string, len(wrapped))
		for j, ln := range wrapped {
			if j == 0 {
				block[j] = marker + " " + prefix + ln
			} else {
				block[j] = indent + ln
			}
		}
		rows = append(rows, strings.Split(style.Render(strings.Join(block, "\n")), "\n")...)
	}
	return rows
}

// knownCodexTerms are the entries the story has introduced so far, sorted.
func (m Model) knownCodexTerms() []string {
	terms := make([]string, 0, len(m.story.Codex))
	for term, entry := range m.story.Codex {
		if game.ChapterIndex(entry.UnlockChapter) <= m.state.MaxChapterIndex {
			terms = append(terms, term)
		}
	}
	sort.Strings(terms)
	return terms
}

func (m Model) unseenCodex() int {
	n := 0
	for _, term := range m.knownCodexTerms() {
		if !m.state.CodexSeen[term] {
			n++
		}
	}
	return n
}

func (m Model) viewHelp(l layout) string {
	lines := append([]string{heading(m.labels.HelpHeading, colorCyan), ""}, m.labels.HelpLines...)
	return m.screen(l, colorCyan, lines, m.labels.MapFooter)
}

// codexOverhead is the heading, its blank line and the footer row — everything
// in the codex that isn't part of the scrollable body.
const codexOverhead = 3

// viewCodex shows only the lexicon entries the story has actually
// introduced so far (by furthest chapter reached), not the full database —
// S.T.E.V.E. hasn't got intel on Vex before you've met him.
func (m Model) viewCodex(l layout) string {
	terms := m.knownCodexTerms()

	var body []string
	for _, term := range terms {
		title := boldStyle.Render(term)
		if !m.state.CodexSeen[term] {
			title += " " + lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render(m.labels.CodexNewTag)
		}
		body = append(body, title)
		body = append(body, wrap(m.story.Codex[term].Text, l.textW)...)
		body = append(body, "")
	}
	if len(terms) == 0 {
		body = []string{dimStyle.Render(m.labels.CodexEmpty)}
	}

	visibleRows := l.rows - codexOverhead
	maxScroll := max(len(body)-visibleRows, 0)
	scroll := min(m.codexScroll, maxScroll)
	end := min(scroll+visibleRows, len(body))
	visible := append([]string{}, body[scroll:end]...)
	for len(visible) < visibleRows {
		visible = append(visible, "")
	}

	footer := fmt.Sprintf(m.labels.CodexFooterFmt, len(terms))
	if maxScroll > 0 {
		footer = fmt.Sprintf(m.labels.CodexScrollFmt, scroll+1, maxScroll+1, footer)
	}
	lines := append([]string{heading(m.labels.CodexHeading, colorAmber), ""}, visible...)
	lines = append(lines, dimStyle.Render(footer))
	return m.screen(l, colorAmber, lines, "")
}

func (m Model) viewMap(l layout) string {
	currentIdx := game.SpineIndex(m.story, m.state.CurrentChapter)

	lines := []string{heading(m.labels.MapHeading, colorCyan), ""}
	if !m.state.Flags["sat_data"] {
		lines = append(lines, dimStyle.Render(m.labels.RouteUnknown), "", dimStyle.Render(m.labels.MapFooter))
		return m.screen(l, colorCyan, lines, "")
	}
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
		if cur := m.story.Chapters[m.state.CurrentChapter]; i == currentIdx && cur.Spine != "" {
			lines = append(lines, dimStyle.Render("       ↳ "+cur.Location))
		}
	}
	lines = window(lines, l.rows-2, 0, false)
	lines = append(lines, "", dimStyle.Render(m.labels.MapFooter))
	return m.screen(l, colorCyan, lines, "")
}

func (m Model) viewTimelock(l layout) string {
	pct := 0
	if m.timelockTotal > 0 {
		pct = 100 - int(m.timelockRemaining*100/m.timelockTotal)
	}
	block := []string{dimStyle.Render(m.labels.TimelockStandby), "", gauge(pct, min(l.textW-8, 60)), "", dimStyle.Render(m.labels.TimelockSkip)}
	return m.screen(l, colorDim, centered(l, block), "")
}

// viewRealLock shows a genuine real-time wait (-realtime): the run is
// saved and it's safe to close the terminal — S.T.E.V.E. will still be
// there when the clock catches up, whether that's from a fresh launch or
// this same session left running.
func (m Model) viewRealLock(l layout) string {
	remaining := max(time.Until(time.Unix(m.state.TimeLockUntil, 0)), 0)
	unlockAt := time.Unix(m.state.TimeLockUntil, 0).Local().Format("Mon 15:04")

	block := []string{heading(m.labels.RealLockHeading, colorDim), ""}
	block = append(block, wrap(m.labels.RealLockBodyPre+formatDuration(remaining)+".", min(l.textW, 70))...)
	block = append(block,
		"",
		dimStyle.Render(m.labels.RealLockBackOnlinePre+unlockAt),
		"",
		dimStyle.Render(m.labels.RealLockSaved),
		"",
		dimStyle.Render(m.labels.RealLockQuit),
	)
	return m.screen(l, colorDim, centered(l, block), "")
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

func (m Model) viewEnd(l layout, c lipgloss.Color) string {
	var block []string
	if m.endTitle != "" {
		block = append(block, heading(m.endTitle, c), "")
	}
	block = append(block, wrap(m.endMessage, min(l.textW, 78))...)
	block = append(block, "", dimStyle.Render(fmt.Sprintf(m.labels.EndStatsFmt, m.stats.Runs, m.stats.Deaths, len(m.stats.Endings), m.endingTotal())), "", dimStyle.Render(m.labels.EndQuit))
	return m.screen(l, c, centered(l, block), "")
}

func (m Model) endingTotal() int {
	n := 0
	for _, ch := range m.story.Chapters {
		for _, c := range ch.Choices {
			if c.EndingID != "" {
				n++
			}
		}
	}
	return n
}
