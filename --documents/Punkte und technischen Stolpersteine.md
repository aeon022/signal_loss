### 1. Cross-Platform Terminal-Verhalten (Windows vs. macOS/Linux)

* **ANSI-Escape-Codes:** Go gibt ANSI-Farben (`\033[36m` etc.) auf Linux und macOS nativ aus. Unter Windows (cmd.exe / PowerShell) kann das ohne das Paket `golang.org/x/term` oder `[github.com/mattn/go-colorable](https://github.com/mattn/go-colorable)` zu unschönen Farb-Müllzeichen im Terminal führen. Da du auf macOS entwickelst, läuft es bei dir out-of-the-box, aber für den Build sollte man das sauber absichern.
* **Bildschirm-Clear:** Befehle wie `clear` oder `cls` plattformübergreifend sauber aufzurufen, ohne dass der Screen flackert oder den Scrollback-Buffer zerstört.

### 2. State-Persistence & Cheat-Sicherheit (JSON)

* Der Spielstand (`save.json`) sollte in einem OS-spezifischen Konfigurationsordner liegen (unter macOS z. B. `~/Library/Application Support/SignalLoss/` oder einfach im aktuellen Verzeichnis als `.signal_loss_save.json`).
* **Wichtig bei echten Timern (Kapitel 3 etc.):** Speichere im JSON nicht nur das Kapitel, sondern einen echten Unix-Timestamp (`time.Now().Unix()`). Wenn der Spieler das Spiel schließt und erst am nächsten Tag wieder öffnet, rechnet Go beim Start aus: `now - saved_time`. Wenn die Zeit vergangen ist, schaltet das Spiel automatisch Kapitel 4 frei. Das verhindert, dass man den Timer durch simples Neustarten des Programms austrickst.

### 3. UX & Typewriter-Effekt abbrechen können

* Der langsame Text-Druck (`time.Sleep` pro Buchstabe) sieht extrem gut aus, nervt aber gewaltig, wenn man das Spiel zum dritten Mal spielt oder schnell weiter will.
* **Best Practice:** Implementiere einen Mechanismus (z. B. wenn man `Enter` drückt), der den aktuellen Typewriter-Effekt abbricht und den Text sofort komplett ausgibt.

### 4. Input-Validierung im Terminal

* Spieler tippen Mist ein (Buchstaben statt Zahlen, Leerzeichen, leere Enters). Der Read-Loop darf wegen einer falschen Taste niemals abstürzen (`panic`).
* Nutze einen robusten Reader (wie `bufio.Reader`), trimme Leerzeichen (`strings.TrimSpace`) und fange ungültige Eingaben mit einer freundlichen S.T.E.V.E.-Fehlermeldung ab ("Boss, die Taste gab es nicht, versuch's noch mal").

### 5. Code-Modularisierung (Keine 2000-Zeilen-`main.go`)

* Da wir 12 ausführliche Kapitel mit je 3 Entscheidungen und Textwüsten haben, wird die `main.go` sonst unlesbar.
* Lagere die Kapitel-Texte und Logiken in ein eigenes Paket (`/story`) aus, sodass die Kapitel saubere Go-Structs sind:
```go
type Choice struct {
    Text        string
    NextChapter int
    Effect      func(s *GameState)
}
type Chapter struct {
    Title   string
    Text    string
    Choices []Choice
}

```



So lässt sich die Story später kinderleicht erweitern oder anpassen, ohne den Programmcode anzufassen.