# Kapitel 8: Der Schwarzmarkt-Deal

> Diese Datei spiegelt den aktuell live geschalteten Inhalt von `story_de.json` Stand 2026-09-10 (49 Entscheidungen über 12 Kapitel, Flag-gesteuerte Callbacks, dichtere Einleitungen) — neu generiert, ersetzt den früheren Entwurf von vor der eigentlichen Spielmechanik. Hier oder direkt in `story_de.json` weiterbearbeiten — danach synchronisiert sich nichts mehr automatisch.

**Ort:** Rust-Bar

Die Rust-Bar-Kneipe riecht nach Lötzinn und Reue, beleuchtet von einem einzigen flackernden Streifen geborgenen Neonlichts, das sich nicht entscheiden kann, welche Farbe es haben will. Ein zweiköpfiger Vermittler schiebt einen orbitalen Navigationskern über die Theke — genau das, was in deinem Sprungplan fehlt — und nennt seinen Preis: den Hyperkondensator, oder eine vollständige Löschung von S.T.E.V.E.s Persönlichkeitspartition. Beide Köpfe beobachten, wie du dich entscheidest, was irgendwie schlimmer ist als nur einer.

---

## Entscheidungen

### [1] Den Hyperkondensator eintauschen — ihn für den Navigationskern übergeben.

**Ergebnis:** Die Navigation ist gesichert. Deine Unterlichttriebwerke müssen ohne Hochleistungs-Booster auskommen, aber wenigstens weißt du jetzt, in welche Richtung es nach Hause geht.

- **Werte:** keine Änderung
- **Gegenstände erhalten:** Navigationskern
- **Gegenstände verloren:** Hyperkondensator
- **Führt zu:** Kapitel 9: Der lange Marsch durch die Sandstürme

### [2] S.T.E.V.E.s Erinnerungen löschen — den Vermittler die Sarkasmus-Partition entfernen lassen.

**Ergebnis:** Du behältst den Kondensator und bekommst den Kern. S.T.E.V.E. dankt dir in einem Ton, der so höflich und steril ist, dass er irgendwie schlimmer wirkt als der Sarkasmus je war.

- **Werte:** keine Änderung
- **Gegenstände erhalten:** Navigationskern
- **Flags gesetzt:** `steve_polite`
- **Führt zu:** Kapitel 9: Der lange Marsch durch die Sandstürme

### [3] Rauchbomben-Coup — eine Rauchbombe werfen, den Kern von der Theke greifen und abhauen.

**Ergebnis:** Du verlässt die Tür mit dem Kern und ohne etwas eingetauscht zu haben — für etwa neunzig Sekunden, bis jede Kopfgeldtafel im Viertel mit deinem Gesicht aufleuchtet. Angeheuerte Kopfgeldjäger jagen dich sauber aus dem Viertel und den ganzen Weg zurück zum Wrack, und du verlierst den Kern, während du es beweist.

- **Werte:** -8 BAT
- **Gegenstände verloren:** Navigationskern
- **Flags gesetzt:** `underworld_hunted`
- **Führt zu:** Kapitel 1: Der Landepunkt

### [4] Die Differenz teilen — anbieten, die Hälfte der Kondensatorladung abzuzapfen, statt ihn komplett aufzugeben.

**Ergebnis:** Der Gesichtsausdruck des Vermittlers — schwer zu lesen bei einem Gesicht ohne Augenbrauen — registriert so etwas wie Respekt vor der Dreistigkeit. Er nimmt die halbe Ladung und lässt einen Teil vom Preis nach. Du behältst einen stark entladenen Kondensator und bekommst trotzdem den Kern.

- **Werte:** -15 BAT
- **Gegenstände erhalten:** Navigationskern
- **Führt zu:** Kapitel 9: Der lange Marsch durch die Sandstürme
