# Teil 1: Projekt-Outline ("SIGNAL LOSS: SECTOR SCHRODINGER")

### 1. Architektur & Tech-Stack

* **Sprache:** Go (pure Standardbibliothek + ein leichtgewichtiges TUI-Framework wie `[github.com/rivo/tview](https://github.com/rivo/tview)` oder `charmbracelet/bubbletea` für den modernen Look).
* **Look & Feel (Retro-Sci-Fi Terminal UI):**
* Farbpalette: Dunkler Hintergrund, Neon-Akzente (Matrix-Grün, Terminal-Cyan, Warn-Pink/Orange).
* Layout: Aufgeteilt in Sektionen (oben Status-Balken für `HULL`, `BAT`, `SCRAP`, mittig die erzählende Story im "Typewriter"-Effekt, unten die interaktiven Antwort-Optionen als nummeriertes Menü).


* **Daten- und Zustandsverwaltung:**
* Ein einfacher `GameState`-Struct in JSON, der lokal im User-Verzeichnis (`~/.config/signal_loss/save.json`) gespeichert wird, damit man das Spiel über mehrere Tage spielen und pausieren kann.
* Echte Zeit-Timer für Warte-Kapitel (z. B. Kapitel 3: "Reaktor kühlt ab – bitte in 2 Stunden wiederkommen").



### 2. Modul-Struktur (`/cmd/signal_loss`)

* `main.go`: Einstiegspunkt, Initialisierung, Laden des Spielstands.
* `game/engine.go`: Steuert den Ablauf der 12 Kapitel, prüft Bedingungen (z. B. ob man Gegenstände wie den Kondensator hat).
* `ui/terminal.go`: Handhabt die TUI-Darstellung, den langsamen Text-Druck (`printSlow`) und saubere Screen-Clears.
* `story/chapters.go`: Enthält alle ausformulierten Kapitel-Texte, Entscheidungen und deren Auswirkungen auf den `GameState`.

---

# Teil 2: Agenten-Anweisung (Prompt für die Umsetzung)

Wenn du dieses Projekt von einem Coding-Agenten (oder einer IDE) in Go schreiben lassen willst, kannst du ihm diesen exakten Befehl geben:

```text
Du bist ein erfahrener Go-Entwickler und Retro-Sci-Fi-Fan. Erstelle ein vollwertiges, interaktives Terminal-Text-Adventure in Go mit dem Titel "SIGNAL LOSS: SECTOR SCHRODINGER".

Technische Anforderungen:
1. Programmiersprache: Go (nutze die Standardbibliothek, optional charmbracelet/bubbletea oder tview für eine saubere Terminal-UI, wenn passend).
2. UI-Design: Das Terminal soll sich wie das Interface einer kaputten Schiffs-KI (S.T.E.V.E.) anfühlen. Nutze ANSI-Farben (Cyan, Grün, Pink/Rot) für Statusleisten und Text.
3. Typewriter-Effekt: Story-Texte und Dialoge von S.T.E.V.E. sollen fließend mit einem leichten Delay ausgegeben werden.
4. Spielmechanik & State: 
   - Ein zentraler State verwaltet: HULL (Hülle in %), BAT (Energie in %), SCRAP (Schrott/Währung) und das aktuelle Kapitel (1 bis 12).
   - Der Spielstand muss in einer lokalen JSON-Datei gespeichert werden, damit das Spiel über Tage hinweg fortgesetzt werden kann.
   - Integriere echte Zeit-Sperren (z. B. in Kapitel 3 muss der Reaktor real abkühlen, bevor es weitergeht).
5. Story-Inhalt:
   - 12 epische, detailreiche Kapitel im Stil von "Lifeline" und "Perry Rhodan" (Absturz auf Gryps-4, Food-Truck-Kartell, Glitch-Nomaden, Recycling-Mönche, Piratenangriff, Säure-Regen, Schwarzmarkt, Sandsturm, Reaktor-Reparatur, Finale Flucht).
   - Jedes Kapitel präsentiert eine Situation mit 3 distinkten Antwortmöglichkeiten, die den State und den weiteren Weg direkt beeinflussen.

Schreibe sauberen, modularen Go-Code mit einer main.go und sprechenden Unterpaketen. Beginne mit der Implementierung der Grundstruktur und Kapitel 1 bis 3.

```