package tui

// Labels holds every UI-chrome string rendered by pkg/tui outside the
// story content itself — headings, hints, footers, boot log lines. Set
// once at Model construction via labelsFor(lang) and read as a plain
// struct field (m.labels.X), never as mutable package state, so multiple
// Models (tests, future multi-instance use) can't stomp on each other.
//
// HULL/BAT/SCRAP/INV stay English acronyms in every language — the status
// bar's fixed-width bar math (see styles.go's statusBar/padPct/pad3) is
// tuned to their exact lengths, and compact stat acronyms staying in
// English is common practice even in fully localized games.
type Labels struct {
	BootSubtitle   string
	BootLink       string
	BootLines      []string
	PressUplink    string
	ErrUnknownChap string
	PressContinue  string
	ChoiceFooterFmt string // %d = choice count

	CodexHeading      string
	CodexEmpty        string
	CodexFooterFmt    string // %d = known count
	CodexScrollFmt    string // %d, %d, %s = pos, max, footer

	MapHeading string
	MapFooter  string

	TimelockStandby string
	TimelockSkip    string

	RealLockHeading   string
	RealLockBodyPre   string // + battery-standby explanation, + duration + "."
	RealLockBackOnlinePre string // + timestamp
	RealLockSaved     string
	RealLockQuit      string

	EndQuit       string
	GameOverTitle string
}

var labelsEN = Labels{
	BootSubtitle:    "// SECTOR SCHRÖDINGER //",
	BootLink:        "CSS RUST-404 — EMERGENCY TERMINAL LINK",
	BootLines: []string{
		"> INITIALIZING S.T.E.V.E. CORE......... OK",
		"> LIFE SUPPORT DIAGNOSTIC............... OK",
		"> HULL INTEGRITY SENSORS................ OK",
		"> BATTERY TELEMETRY..................... OK",
		"> HOME SIGNAL............................ LOST",
	},
	PressUplink:     "Press any key to establish uplink...",
	ErrUnknownChap:  "ERROR: unknown chapter",
	PressContinue:   "Press any key to continue...",
	ChoiceFooterFmt: "[1-%d] choose   ↑↓ + enter   [c]odex   [m]ap   [q]uit",

	CodexHeading:   "S.T.E.V.E. DATABASE — CODEX",
	CodexEmpty:     "No data on file yet. Keep moving.",
	CodexFooterFmt: "%d known — [c/esc] back to chapter",
	CodexScrollFmt: "[↑↓ scroll %d/%d]  %s",

	MapHeading: "NAVIGATION — SECTOR SCHRÖDINGER",
	MapFooter:  "[m/esc] back to chapter",

	TimelockStandby: "SIGNAL LOST — standing by...",
	TimelockSkip:    "[any key] skip wait",

	RealLockHeading:       "SIGNAL LOST — TRANSMISSION PAUSED",
	RealLockBodyPre:       "S.T.E.V.E. has gone into standby to conserve what's left of the battery. Reconnect in ",
	RealLockBackOnlinePre: "Back online: ",
	RealLockSaved:         "Your run is saved — safe to close this window and come back.",
	RealLockQuit:          "[q]uit",

	EndQuit:       "[any key] quit",
	GameOverTitle: "💀 GAME OVER",
}

var labelsDE = Labels{
	BootSubtitle:    "// SEKTOR SCHRÖDINGER //",
	BootLink:        "CSS RUST-404 — NOTFALL-TERMINALVERBINDUNG",
	BootLines: []string{
		"> INITIALISIERE S.T.E.V.E.-KERN......... OK",
		"> LEBENSERHALTUNGS-DIAGNOSE.............. OK",
		"> RUMPFINTEGRITÄTS-SENSOREN.............. OK",
		"> BATTERIE-TELEMETRIE.................... OK",
		"> HEIMATSIGNAL............................ VERLOREN",
	},
	PressUplink:     "Beliebige Taste drücken, um den Uplink herzustellen...",
	ErrUnknownChap:  "FEHLER: unbekanntes Kapitel",
	PressContinue:   "Beliebige Taste drücken, um fortzufahren...",
	ChoiceFooterFmt: "[1-%d] wählen   ↑↓ + Enter   [c]odex   [m]ap   [q] beenden",

	CodexHeading:   "S.T.E.V.E.-DATENBANK — CODEX",
	CodexEmpty:     "Noch keine Daten vorhanden. Weiter vorankommen.",
	CodexFooterFmt: "%d bekannt — [c/esc] zurück zum Kapitel",
	CodexScrollFmt: "[↑↓ scrollen %d/%d]  %s",

	MapHeading: "NAVIGATION — SEKTOR SCHRÖDINGER",
	MapFooter:  "[m/esc] zurück zum Kapitel",

	TimelockStandby: "SIGNAL VERLOREN — Standby...",
	TimelockSkip:    "[beliebige Taste] Wartezeit überspringen",

	RealLockHeading:       "SIGNAL VERLOREN — ÜBERTRAGUNG PAUSIERT",
	RealLockBodyPre:       "S.T.E.V.E. ist in den Standby-Modus gegangen, um den Rest der Batterie zu schonen. Wiederverbindung in ",
	RealLockBackOnlinePre: "Wieder online: ",
	RealLockSaved:         "Dein Spielstand ist gespeichert — du kannst dieses Fenster sicher schließen und später zurückkommen.",
	RealLockQuit:          "[q] beenden",

	EndQuit:       "[beliebige Taste] beenden",
	GameOverTitle: "💀 SPIEL VORBEI",
}

// labelsFor resolves a lang code ("en", "de") to its Labels set, defaulting
// to English for anything else (including "").
func labelsFor(lang string) Labels {
	if lang == "de" {
		return labelsDE
	}
	return labelsEN
}
