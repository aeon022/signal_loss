# signal_loss – Erweiterungsvorschlag v0.2 (Entwurf, 26.10.2026)

Ausgangslage (geprüft am Code und an `story.json`): 12 Kapitel, 49 Entscheidungen, 4 Enden, EN + DE mechanisch identisch (Test `pkg/story/story_test.go`). Die Handlung ist eine **gerade Kette**: Jedes Kapitel hat drei brauchbare Optionen, die alle ins selbe nächste Kapitel führen und sich nur in den Werten unterscheiden; pro Kapitel gibt es eine Sackgasse (Tod oder Rückwurf). Alle fünf Rückwürfe (Kap. 2, 4, 6, 8, 9) landen in Kapitel 1. Jede der sieben Flags hat genau einen Nutzen.

Reihenfolge: **1 Story → 2 UI/UX → 3 Browser**. Jede Richtung wird für sich gebaut, getestet, committet und gepusht; erst danach die nächste.

---

## 1. Story

**Ziel:** Entscheidungen sollen den Weg verändern, nicht nur die Werte. Mehr Spielzeit, mehr Wiederspielwert, ohne den Ton (sarkastische S.T.E.V.E., Müll-Sci-Fi) zu ändern.

| Baustein | Beispiel | Aufwand |
|---|---|---|
| **A. Gabelung in der Mitte** | Nach Kapitel 5 zwei Routen, die in Kapitel 9 wieder zusammenlaufen. Route *Mönche*: „Der Hyper-Kondensator" (Tempel stehlen oder überzeugen). Route *Kartell*: „Der Schwarzmarkt-Konvoi" (mitfahren oder überfallen). In Kapitel 6 wird „Die Platine an Vex verkaufen" zum Eintritt in Kapitel 6b statt 7. | ca. 6 neue Kapitel, ca. 24 Entscheidungen |
| **B. Rückwürfe mit Folgen** | Der gescheiterte Diebstahl in Kapitel 2 führt nicht mehr nach Kapitel 1, sondern in ein Mini-Kapitel „Die Cousins, Schuldenabarbeiten" und danach zum Basar zurück, mit Kerbe im Anzug (−1 HULL, Item „Schuldschein"). | ca. 5 Mini-Kapitel, ca. 10 Entscheidungen; **keine Engine-Änderung**, `next_chapter` reicht |
| **C. Ruf bei den Fraktionen** | Zähler für Kartell, Mönche, Glitch-Nomaden und Vex statt einzelner Flags. Schwellen schalten Entscheidungen und Enden frei. Neue Enden: *Der Apostel* (Mönche), *Der Franchisenehmer* (Kartell), *Die Aufstiegs-Datei* (Glitch). Von 4 auf 7 Enden. | Engine: Feld `reputation` in `StatMutations`, `requires_rep` an `Choice` |
| **D. S.T.E.V.E.-Vertrauen (versteckt)** | Wer S.T.E.V.E. ständig übergeht, hat im Finale keinen Notfall-Reboot; wer höflich war (Flag `steve_polite`), bekommt einen Extrasatz im Epilog. | klein, nutzt Ruf-Mechanik |
| **E. Neues Spiel+** | Nach einem Ende bleiben Codex-Einträge und gesehene Enden erhalten; ein verstecktes Kapitel 13 erscheint erst nach drei verschiedenen Enden. | klein, braucht Statistik-Datei aus Teil 2 |

**Technische Vorbedingung:** `ChapterOrder` ist heute eine lineare Liste, die den Codex steuert (`MaxChapterIndex`). Verzweigte Kapitel brauchen ein Feld `order` oder `act` im Kapitel, sonst schaltet eine Route Einträge der anderen frei oder nicht.

**Absicherung (Tests, vor dem Schreiben):** Jedes Kapitel ist vom Start erreichbar, jedes Ende ist erreichbar, kein Kapitel ohne Ausgang, jede `requires_rep`-Schwelle ist erreichbar (Summe möglicher Ruf-Punkte), EN und DE bleiben mechanisch identisch.

**Vorschlag zur Reihenfolge:** zuerst B (günstig, keine Engine), dann A, dann C. D und E danach, wenn der Ton stimmt. Geschrieben wird zuerst Englisch, dann Deutsch; den Ton prüfst du an einem Probekapitel, bevor ich alle schreibe.

---

## 2. UI/UX

| Baustein | Beispiel | Aufwand |
|---|---|---|
| **Fenstergröße** | Statt fester 76 Spalten passt sich die Anzeige an (60–100 Spalten), `tea.WindowSizeMsg`. Sehr kleine Fenster zeigen einen Hinweis statt zu zerbrechen. | mittel |
| **Statusleiste als Balken** | `HULL ▰▰▱▱▱▱ 12%`; bei ≤ 3 % rot und blinkend, `BAT` ab 10 % gelb. Zahlen bleiben daneben stehen. | klein |
| **Tempo steuern** | Leertaste überspringt das Tippen eines Textes, `?` zeigt die Tastenhilfe, sichtbar im Fuß. | klein |
| **Codex** | Neue Einträge mit `NEU`-Marke, Suche mit `/`, Scrollen bleibt. | klein |
| **Karte** | Besuchte Kapitel mit Fraktionszeichen, noch nicht besuchte als `???`, aktuelle Route hervorgehoben (wichtig für die Gabelung aus Teil 1). | mittel |
| **Statistik nach dem Spiel** | Ende, Todesursache, gewählte Wege, Anzahl Durchläufe; im Terminal in `stats.json`, im Browser in `localStorage`. Grundlage für „Neues Spiel+". | mittel |
| **Ton (nur Browser)** | Tastenklicken, Alarm bei niedriger HULL, Boot-Brummen, Web Audio, aus per Taste `M`, standardmäßig aus bis zur ersten Eingabe. Im Terminal höchstens das Terminalsignal bei einem Tod. | mittel |
| **Barrierefreiheit** | Option „weniger Bewegung" (keine Glitch-Übergänge) und hoher Kontrast; im Browser über `prefers-reduced-motion`. | klein |

Beispiel für die Leiste (Entwurf):

```
HULL ▰▰▱▱▱▱▱▱ 12%   BAT ▰▰▰▰▱▱▱▱ 45%   SCRAP 2   INV 1
```

---

## 3. Browser

**Ziel:** Browser und Terminal bleiben gleich (Engine und Story sind es heute schon). Offen sind Funktionen, die nur auf einer Seite existieren, und die Seite um das Spiel.

| Baustein | Beispiel | Aufwand |
|---|---|---|
| **Spielstand tauschen** | „Export" und „Import" im Browser, damit man im Terminal weiterspielen kann (und umgekehrt). Heute sind Saves nach Sprache getrennt (`signal_loss_save_en/_de`) und nicht austauschbar. | klein |
| **Touch-Bedienung** | Auf dem Handy Tasten 1–4 und Pfeile als Knöpfe unter dem Terminal. | klein |
| **Seite auf Deutsch** | Die Seite um das Spiel ist nur Englisch. Sprache folgt `navigator.language`, Schalter oben, `hreflang`. | mittel |
| **404-Seite und Devlog** | Eigene 404; Devlog aus den „Kriegsgeschichten" in `roadmap.md`. | klein |
| **Teilen** | Nach einem Ende ein Knopf „Ergebnis kopieren": „Ich bin als *Der Laute* von Gryps-4 entkommen." Eigenes OG-Bild je Ende. | mittel |
| **Prüfung im CI** | Der Deploy-Lauf führt zuerst `go test` aus und danach einen Puppeteer-Rauchtest (Seite lädt, erste Entscheidung funktioniert, beide Sprachen). | mittel |
| **itch.io und eigene Adresse** | Eintrag auf itch.io; `abteilung83.at/…` leitet auf die Seite. | klein, Handarbeit von dir |

**Entscheidung nötig – Zähler:** Die Seite verspricht heute „kein Tracking, nur Spielstand im Browser". Jede Statistik bricht dieses Versprechen oder muss es ändern. Vorschlag: **keine Analytics**, Interesse an GitHub-Sternen und Polar-Spenden ablesen. Wenn du doch zählen willst, dann nur ein datenschutzfreundlicher Zähler ohne Cookies, und der Hinweistext wird angepasst.

---

## Reihenfolge und Version

1. Teil 1 in der Reihenfolge B, A, C, danach D, E. → Tag **v0.2.0**
2. Teil 2 (ohne Ton) → **v0.2.x**; Ton danach.
3. Teil 3; die Seite wird nach jedem Schritt neu ausgeliefert (`workflow_dispatch` auf `deploy/landing`).

Nach jedem Schritt: `go build ./... && go vet ./... && go test ./...`, Commit, Push, bei Browser-Änderungen Deploy und Prüfung gegen die Live-Seite.
