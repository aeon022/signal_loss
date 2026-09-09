package tui

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

var ansiCode = regexp.MustCompile("\x1b\\[[0-9;]*m")

// visibleWidth is the rune count once ANSI color codes are stripped — used
// to pad an already-styled line (like the status readout) to a fixed width.
func visibleWidth(s string) int {
	return utf8.RuneCountInString(ansiCode.ReplaceAllString(s, ""))
}

// MaxWidth and FixedHeight are the screen's fixed columns/rows — like a
// monitor, not a scrolling terminal: they never change with what's on
// screen. FixedHeight (30) comfortably covers the tallest real screen
// (chapter 1's narrative + 3 wrapped choices = 23 rows at this width);
// content taller than that (the codex) scrolls inside the frame instead
// of growing it.
const (
	MaxWidth    = 76
	FixedHeight = 30
)

// FrameWidth is the total rendered width of the inner bordered box.
// Width(MaxWidth) already includes the 1-column padding on each side (see
// TextWidth above); the border then adds 1 more column on each side.
const FrameWidth = MaxWidth + 2

// TextWidth is how wide text may actually get before Render()'s own
// wrapping kicks in a second time: frameStyle's Width(MaxWidth) counts the
// 1-column padding on each side as part of that budget, so the real text
// budget is 2 columns narrower. Wrap everything that goes inside a frame
// to this, not to MaxWidth directly.
const TextWidth = MaxWidth - 2

// BezelWidth is the outer bezel's total width: the inner box plus a
// 1-space, 1-border-char margin on each side.
const BezelWidth = FrameWidth + 4

var (
	colorCyan    = lipgloss.Color("14")
	colorAmber   = lipgloss.Color("3")
	colorGreen   = lipgloss.Color("10")
	colorRed     = lipgloss.Color("9")
	colorMagenta = lipgloss.Color("13")
	colorDim     = lipgloss.Color("8")
)

// retroBorder is a single-line terminal box for the inner "screen" — kept
// plain since the outer bezel (see chrome) now carries the decoration.
var retroBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
}

// frameStyle borders and pads a block of text to the fixed MaxWidth — every
// screen is the same size, CRT-monitor style, regardless of how much text
// is actually on it.
func frameStyle(c lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(retroBorder).
		BorderForeground(c).
		Padding(0, 1).
		Width(MaxWidth).
		Height(FixedHeight)
}

// chrome wraps a bordered box in an outer monitor bezel: corner screws,
// a single-line housing, and the terminal's label/status plates embedded
// directly into the top and bottom edges instead of floating above/below.
// statusLine, if non-empty, renders as an instrument-panel readout between
// the screen and the bottom bezel edge — a spaceship console gauge strip,
// not part of the "screen" itself. body must already be exactly FrameWidth
// columns per line (frameStyle guarantees this).
func chrome(body, statusLine string) string {
	top := dimStyle.Render(bezelEdge("S.T.E.V.E. TERMINAL — MK.IV"))
	bottom := dimStyle.Render(bezelEdge("SIGNAL: LIVE — SECTOR SCHRÖDINGER"))

	var out []string
	out = append(out, top)
	for _, line := range strings.Split(body, "\n") {
		out = append(out, dimStyle.Render("│ ")+line+dimStyle.Render(" │"))
	}
	if statusLine != "" {
		pad := FrameWidth - visibleWidth(statusLine)
		if pad < 0 {
			pad = 0
		}
		blank := dimStyle.Render("│ ") + strings.Repeat(" ", FrameWidth) + dimStyle.Render(" │")
		out = append(out, blank)
		out = append(out, dimStyle.Render("│ ")+statusLine+strings.Repeat(" ", pad)+dimStyle.Render(" │"))
		out = append(out, blank)
	}
	out = append(out, bottom)
	return strings.Join(out, "\n") + "\n"
}

// bezelEdge draws one horizontal bezel edge: ●──[ label ]──●, centered,
// exactly BezelWidth columns wide.
func bezelEdge(label string) string {
	text := "[ " + label + " ]"
	textWidth := utf8.RuneCountInString(text)
	inner := BezelWidth - 2 // the two corner screws
	if textWidth > inner {
		textWidth = inner
		text = text[:inner]
	}
	padTotal := inner - textWidth
	left := padTotal / 2
	right := padTotal - left
	return "●" + strings.Repeat("─", left) + text + strings.Repeat("─", right) + "●"
}

var (
	boldStyle   = lipgloss.NewStyle().Bold(true)
	dimStyle    = lipgloss.NewStyle().Foreground(colorDim)
	choiceStyle = lipgloss.NewStyle().Foreground(colorGreen)
	cursorStyle = lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Reverse(true)
)

// anchorBottom lays out top followed by bottom, padded with enough blank
// lines to reach FixedHeight — so bottom (e.g. the choice list) sits at a
// consistent position near the bottom of the screen instead of right after
// however much narrative text happens to precede it.
func anchorBottom(top, bottom []string) []string {
	filler := FixedHeight - len(top) - len(bottom)
	if filler < 0 {
		filler = 0
	}
	lines := make([]string, 0, FixedHeight)
	lines = append(lines, top...)
	for i := 0; i < filler; i++ {
		lines = append(lines, "")
	}
	lines = append(lines, bottom...)
	return lines
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
		case len(line)+1+len(word) > width:
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

// bar renders a fixed-width block-character gauge, colored by threshold.
func bar(pct, width int) string {
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}
	filled := pct * width / 100
	color := colorGreen
	switch {
	case pct <= 20:
		color = colorRed
	case pct <= 50:
		color = colorAmber
	}
	style := lipgloss.NewStyle().Foreground(color)
	return style.Render(strings.Repeat("█", filled) + strings.Repeat("░", width-filled))
}

func statusBar(hull, bat, scrap, inventoryCount int) string {
	return boldStyle.Render("HULL") + " " + bar(hull, 10) + boldStyle.Render(padPct(hull)) +
		"   " + boldStyle.Render("BAT") + " " + bar(bat, 10) + boldStyle.Render(padPct(bat)) +
		"   " + boldStyle.Render("SCRAP") + " " + boldStyle.Render(pad3(scrap)) +
		"   " + boldStyle.Render("INV") + " " + boldStyle.Render(pad3(inventoryCount))
}

func padPct(n int) string {
	return lipgloss.NewStyle().Width(4).Align(lipgloss.Right).Render(strconv.Itoa(n) + "%")
}

func pad3(n int) string {
	return lipgloss.NewStyle().Width(3).Render(strconv.Itoa(n))
}
