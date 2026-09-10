# Kapitel 10: Die Reparatur-Krise

> Diese Datei spiegelt den aktuell live geschalteten Inhalt von `story_de.json` Stand 2026-09-10 (49 Entscheidungen über 12 Kapitel, Flag-gesteuerte Callbacks, dichtere Einleitungen) — neu generiert, ersetzt den früheren Entwurf von vor der eigentlichen Spielmechanik. Hier oder direkt in `story_de.json` weiterbearbeiten — danach synchronisiert sich nichts mehr automatisch.

**Ort:** Das Wrack — Reaktorraum

Teile zusammengebaut, Rumpf so gut repariert, wie es Klebeband und Trotz erlauben, verlangt der Hauptreaktor eine administrative Override-Passphrase, bevor er zündet. Drei Versuche vor der permanenten Sperre. Irgendwo hinter dem Grat nähern sich Fraktionen dem Lärm, und der Cursor der Konsole blinkt mit etwas, das sich wie persönliches Urteil anfühlt.

---

## Entscheidungen

### [1] Standard-Reset-Code (ADMIN_0000) — das werkseitige Standardpasswort.

**Ergebnis:** Angenommen, technisch gesehen. Es sendet dabei auch lautlos ein Signal an eine Konzern-Sicherheitsdrohne irgendwo im Orbit, die laut Warnprotokoll jetzt „im Anflug“ ist.

- **Werte:** keine Änderung
- **Flags gesetzt:** `drones_inbound`
- **Führt zu:** Kapitel 11: Der finale Countdown

### [2] S.T.E.V.E.s Eitelkeitspasswort (STEVE_IST_EIN_GENIE_42).

**Ergebnis:** Das Easter Egg ist real. Der Reaktor schnurrt zum Leben, als hätte er sein ganzes Diensteben darauf gewartet, dass jemand das eintippt, und S.T.E.V.E. ist, ausnahmsweise, zu selbstgefällig, um sarkastisch zu sein.

- **Werte:** +2 HULL, +10 BAT
- **Führt zu:** Kapitel 11: Der finale Countdown

### [3] Brecheisen-Bypass-Relais — eine Eisenstange zwischen die Hochspannungskontakte klemmen.

**Ergebnis:** Funken sprühen über die Konsole. Der Reaktor zündet nicht einfach — er zündet, während du noch über das offene Panel gebeugt bist. Davon gibt es keine überlebbare Version.

- **Werte:** keine Änderung
- 💀 **TÖDLICH — Lauf endet hier.**

### [4] Formaler Override-Antrag — den neu höflichen S.T.E.V.E. den Zugriff über offizielle Firmenkanäle beantragen lassen, statt ein Passwort zu raten. _(**schaltet frei nach:** `steve_polite`)_

**Ergebnis:** S.T.E.V.E.s unerbittlich höfliche neue Stimme trägt einen makellosen, öde formellen Zugriffsantrag vor — genau die Art, der Firmen-Sicherheitssysteme eingebaut vertrauen. Der Reaktor gewährt den Override sauber, keine roten Flaggen, keine Drohne entsendet. Höflichkeit ist ausnahmsweise der Exploit.

- **Werte:** +1 HULL, +5 BAT
- **Führt zu:** Kapitel 11: Der finale Countdown

### [5] Override unter Beschuss — die Kopfgeldjäger auf deiner Spur holen dich mitten in der Passphrase ein; den Standardcode durchjagen und das Firmensignal als kleineres Risiko akzeptieren. _(**schaltet frei nach:** `underworld_hunted`)_

**Ergebnis:** Die Jäger kommen zwei Sekunden nach ADMIN_0000 durch die Tür. Du bist schon in Bewegung, als sie die Schwelle erreichen, der Reaktor summt hinter dir. Das Signal an die Firmensicherheit ist jetzt offensichtlich dein geringstes Problem.

- **Werte:** -2 HULL
- **Flags gesetzt:** `drones_inbound`
- **Führt zu:** Kapitel 11: Der finale Countdown
