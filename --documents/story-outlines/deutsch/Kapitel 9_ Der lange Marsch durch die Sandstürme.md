# Kapitel 9: Der lange Marsch durch die Sandstürme

> Diese Datei spiegelt den aktuell live geschalteten Inhalt von `story_de.json` Stand 2026-09-10 (49 Entscheidungen über 12 Kapitel, Flag-gesteuerte Callbacks, dichtere Einleitungen) — neu generiert, ersetzt den früheren Entwurf von vor der eigentlichen Spielmechanik. Hier oder direkt in `story_de.json` weiterbearbeiten — danach synchronisiert sich nichts mehr automatisch.

**Ort:** Die Sandsturm-Dünen

Ein blendender eiserner Sandsturm zieht auf dem Rückweg zum Schiff auf, Sandkörner schleifen die Farbe von jeder exponierten Oberfläche. Die Elektrostatik baut sich so schnell auf, dass S.T.E.V.E. alle paar Sekunden Blitzschlag-Risiko meldet, mit dem Enthusiasmus eines Rauchmelders, der eine frische Batterie gefunden hat. Die Sicht sinkt auf Armlänge; die Dünen voraus sind nur noch Andeutung und Rauschen.

---

## Entscheidungen

### [1] Direkt durch den Krater — geradewegs in den heulenden Sturm marschieren.

**Ergebnis:** Du halbierst die Überquerungszeit. Irgendwo auf halber Strecke bricht Sand durch die Halsdichtung deines Anzugs, und du spürst für den Rest des Tages jedes einzelne Körnchen davon.

- **Werte:** -3 HULL, -5 BAT
- **Führt zu:** Kapitel 10: Die Reparatur-Krise

### [2] Der Klippenpfad — eine tückische Felskante außerhalb des direkten Winds erklimmen.

**Ergebnis:** Loses Gestein rutscht dir zweimal unter den Füßen weg. Beide Male klammerst du dich mit den Fingerspitzen hoch, Ausrüstung klappernd, und schaffst es hinüber — deine Würde in schlechterem Zustand als dein Anzug.

- **Werte:** -2 BAT
- **Führt zu:** Kapitel 10: Die Reparatur-Krise

### [3] In den Dünen ausharren — den Sturm unter einer Stahlplatte abwarten.

**Ergebnis:** Null physischer Schaden durch den Sturm selbst. Aber wer so lange an einem Ort ausharrt, lässt die Fraktionen auf seiner Spur aufholen — du tauchst auf und findest dein Schiffslager bereits geplündert vor, mit feindlichen Kundschaftern, die auf dich warten, und musst den ganzen Weg zurück zum Landepunkt fliehen.

- **Werte:** keine Änderung
- **Flags gesetzt:** `enemies_at_ship`
- **Zeitsperre:** 240 Minuten
- **Führt zu:** Kapitel 1: Der Landepunkt

### [4] Durch die Versorgungslinien des Kartells schneiden — du bist bereits ein markiertes Gesicht; nutze es, um an einem versteckten Kontrollposten in den Dünen vorbeizubluffen. _(**schaltet frei nach:** `wanted_by_cartel`)_

**Ergebnis:** Es stellt sich heraus, dass dem Kartell bekannt zu sein in beide Richtungen wirkt — du redest dich am Kontrollposten vorbei, indem du behauptest, dich selbst woanders zur Kopfgeld-Abholung zu liefern. Generell eine schlechte Idee, funktioniert genau einmal. Du bist durch den schlimmsten Abschnitt des Sturms in halber Zeit, das Adrenalin leistet, was Unterschlupf nicht konnte.

- **Werte:** -1 HULL, -1 BAT
- **Führt zu:** Kapitel 10: Die Reparatur-Krise
