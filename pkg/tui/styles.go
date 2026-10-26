package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// The screen fills whatever terminal it runs in. Below minW×minH there is no
// sensible layout, so the game says so instead (see View). Text never gets
// wider than maxTextW: very long lines are hard to read on a big monitor.
const (
	minW     = 40
	minH     = 16
	maxTextW = 104
	sideMinW = 96 // from this width on, the instrument sidebar appears
)

var (
	colorCyan    = lipgloss.Color("14")
	colorAmber   = lipgloss.Color("3")
	colorGreen   = lipgloss.Color("10")
	colorRed     = lipgloss.Color("9")
	colorMagenta = lipgloss.Color("13")
	colorDim     = lipgloss.Color("8")
)

var (
	boldStyle   = lipgloss.NewStyle().Bold(true)
	dimStyle    = lipgloss.NewStyle().Foreground(colorDim)
	choiceStyle = lipgloss.NewStyle().Foreground(colorGreen)
	cursorStyle = lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Reverse(true)
)

// layout is derived from the terminal size on every render, never stored.
type layout struct {
	w, h    int
	side    bool // sidebar shown
	sideW   int  // outer width of the sidebar, border included
	mainW   int  // outer width of the main panel
	bodyH   int  // rows inside the panels
	head    bool // title bar shown (hidden on short terminals to save a row)
	textW   int  // widest text line inside the main panel
	indent  string
	sideTxt int // widest text line inside the sidebar
	rows    int // rows a view may fill inside the main panel (the one-line status takes two without a sidebar)
}

func (m Model) layout() layout {
	w, h := m.width, m.height
	if w == 0 || h == 0 { // no size reported yet (or not a tty)
		w, h = 100, 36
	}
	l := layout{w: w, h: h, head: h >= 24}
	l.bodyH = h - 3 // key bar and the panel borders
	if l.head {
		l.bodyH-- // title bar
	}
	l.mainW = w
	if w >= sideMinW {
		l.side = true
		l.sideW = 34
		if w >= 140 {
			l.sideW = 42
		}
		l.mainW = w - l.sideW
	}
	l.textW = min(l.mainW-6, maxTextW)
	l.indent = strings.Repeat(" ", (l.mainW-6-l.textW)/2)
	l.sideTxt = l.sideW - 4
	l.rows = l.bodyH
	if !l.side {
		l.rows -= 2
	}
	return l
}

// full is the layout for screens that use the whole width (boot, endings).
func (l layout) full() layout {
	l.side, l.sideW, l.mainW = false, 0, l.w
	l.textW = min(l.mainW-6, maxTextW)
	l.indent = strings.Repeat(" ", (l.mainW-6-l.textW)/2)
	l.rows = l.bodyH
	return l
}

func (l layout) tooSmall() bool { return l.w < minW || l.h < minH }

// panel draws a double-line box of the given outer width and inner height.
func panel(c lipgloss.Color, outerW, innerH, padX int, lines []string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).BorderForeground(c).
		Padding(0, padX).
		Width(outerW - 2).Height(innerH).
		Render(strings.Join(lines, "\n"))
}

// bar fills a whole row with text on the left and right, in reverse video.
func bar(c lipgloss.Color, w int, left, right string) string {
	left = lipgloss.NewStyle().MaxWidth(w).Render(left) // never wider than the screen
	gap := w - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 1 {
		right, gap = "", max(w-lipgloss.Width(left), 0)
	}
	return lipgloss.NewStyle().Bold(true).Foreground(c).Reverse(true).Render(left + strings.Repeat(" ", gap) + right)
}

// anchorBottom pads between top and bottom so bottom sits at the foot of an
// area height rows tall: choices stay in one spot however long the text is.
func anchorBottom(top, bottom []string, height int) []string {
	lines := append([]string{}, top...)
	for i := len(top) + len(bottom); i < height; i++ {
		lines = append(lines, "")
	}
	return append(lines, bottom...)
}

// window shows rows of lines starting at scroll, marking what is cut off.
// With follow set it sticks to the end instead (text still being typed).
func window(lines []string, rows, scroll int, follow bool) []string {
	if len(lines) <= rows || rows < 3 {
		return lines
	}
	maxScroll := len(lines) - rows
	switch {
	case follow:
		scroll = maxScroll
	case scroll > maxScroll:
		scroll = maxScroll
	}
	out := append([]string{}, lines[scroll:scroll+rows]...)
	if scroll > 0 {
		out[0] = dimStyle.Render("▲ …")
	}
	if scroll < maxScroll {
		out[rows-1] = dimStyle.Render("▼ …")
	}
	return out
}

// wrap breaks plain text into lines no wider than width. Apply color/bold
// styling AFTER wrapping, not before — styling first would feed ANSI codes
// into the width math.
func wrap(text string, width int) []string {
	if text == "" {
		return []string{""}
	}
	var lines []string
	var line string
	for _, word := range strings.Fields(text) {
		switch {
		case line == "":
			line = word
		case lipgloss.Width(line)+1+lipgloss.Width(word) > width:
			lines = append(lines, line)
			line = word
		default:
			line += " " + word
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

// gauge renders a segmented meter, colored by how bad it is; a nearly empty
// one blinks where the terminal supports it.
func gauge(pct, width int) string {
	pct = max(0, min(pct, 100))
	filled := pct * width / 100
	if pct > 0 && filled == 0 {
		filled = 1
	}
	color := colorGreen
	switch {
	case pct <= 20:
		color = colorRed
	case pct <= 50:
		color = colorAmber
	}
	style := lipgloss.NewStyle().Foreground(color).Blink(pct <= 10)
	return style.Render(strings.Repeat("▰", filled)) + dimStyle.Render(strings.Repeat("▱", width-filled))
}

func pad(s string, n int) string { return lipgloss.NewStyle().Width(n).Render(s) }

// statusBar is the one-line readout used when the sidebar does not fit;
// it gets plainer as the width shrinks.
func statusBar(width, hull, bat, scrap, inventoryCount int) string {
	switch {
	case width >= 56:
		return boldStyle.Render("HULL") + " " + gauge(hull, 8) + " " + pad(strconv.Itoa(hull)+"%", 4) +
			boldStyle.Render("BAT") + " " + gauge(bat, 8) + " " + pad(strconv.Itoa(bat)+"%", 4) +
			boldStyle.Render("SCRAP") + " " + pad(strconv.Itoa(scrap), 3) +
			boldStyle.Render("INV") + " " + strconv.Itoa(inventoryCount)
	case width >= 36:
		return fmt.Sprintf("%s %d%%  %s %d%%  %s %d", boldStyle.Render("HULL"), hull, boldStyle.Render("BAT"), bat, boldStyle.Render("SCRAP"), scrap)
	default:
		return fmt.Sprintf("H%d B%d S%d", hull, bat, scrap)
	}
}

func sideRow(label string, pct, width int) string {
	return boldStyle.Render(pad(label, 6)) + gauge(pct, width-6-5) + " " + fmt.Sprintf("%3d%%", pct)
}
