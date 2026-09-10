# Kapitel 5: Der Tempel der Recycling-Mönche

> Diese Datei spiegelt den aktuell live geschalteten Inhalt von `story_de.json` Stand 2026-09-10 (49 Entscheidungen über 12 Kapitel, Flag-gesteuerte Callbacks, dichtere Einleitungen) — neu generiert, ersetzt den früheren Entwurf von vor der eigentlichen Spielmechanik. Hier oder direkt in `story_de.json` weiterbearbeiten — danach synchronisiert sich nichts mehr automatisch.

**Ort:** Tempel von Sankt-Plastik

Eine Kathedrale aus Solarpaneelen und geborgenen Hauptplatinen erhebt sich aus dem Schluchtboden, ihre Türme behängt mit Gebetsfahnen aus antistatischen Tüten. Drinnen rezitieren die Recycling-Mönche von Sankt-Plastik BIOS-Fehlercodes vor einem goldbeschichteten Altar, auf dem der legendäre Hyperkondensator liegt — genau die Art von Hochleistungs-Energiezelle, die deinen Reaktor wieder in Gang bringen könnte. Der Weihrauchrauch riecht unverkennbar nach brennendem Lötzinn.

---

## Entscheidungen

### [1] Als Pilger verkleiden — sich in eine Plastikplane wickeln und Fehler 404 rezitieren.

**Ergebnis:** Der Scanner-Mönch stuft dich als harmlosen Fanatiker ein. Du schlüpfst hinein, hebst den Kondensator auf und schlüpfst wieder hinaus, ohne dass sich auch nur eine Augenbraue hinter einem Visier hebt.

- **Werte:** keine Änderung
- **Gegenstände erhalten:** Hyperkondensator
- **Führt zu:** Kapitel 6: Der Akt der Sabotage

### [2] Schnapp-und-weg-Sprint — die Türen aufbrechen und zum Altar rennen.

**Ergebnis:** Alarme kreischen. Du greifst den Kondensator im Laufschritt und fängst dafür einen Elektroschlagstock-Treffer an der Schulter ein, aber du bist schon durch die Türen, bevor die Mönche sich organisieren.

- **Werte:** -2 HULL, -8 BAT
- **Gegenstände erhalten:** Hyperkondensator
- **Führt zu:** Kapitel 6: Der Akt der Sabotage

### [3] Das Solarnetz sabotieren — die Hauptleitung durchtrennen und den Tempel in Dunkelheit tauchen.

**Ergebnis:** Der Tempel wird schwarz, und die Mönche, wie sich herausstellt, sind ohne Licht keineswegs harmlos. Der Gesang wird zu Geschrei, dann zu etwas Organisierterem. Du bekommst keine Gelegenheit mehr, dich zu erklären, bevor die Dunkelheit sich endgültig schließt.

- **Werte:** keine Änderung
- 💀 **TÖDLICH — Lauf endet hier.**

### [4] Den glitchenden S.T.E.V.E. den Gesang anführen lassen — sein zerfetztes Pseudo-Latein klingt unheimlich nah an der eigenen BIOS-Liturgie der Mönche. _(**schaltet frei nach:** `steve_glitched`)_

**Ergebnis:** S.T.E.V.E.s verstümmeltes latein-ähnliches Gebrabbel passt so gut zum Fehlercode-Gesang der Mönche, dass drei von ihnen mitzunicken beginnen. Niemand hinterfragt einen Glaubensbruder, der in Zungen spricht. Du hebst den Kondensator mitten in der Zeremonie ab, und ein Mönch drückt dir auf dem Weg hinaus eine Ersatzsicherung in die Hand, weil er dich für Klerus hält.

- **Werte:** keine Änderung
- **Gegenstände erhalten:** Hyperkondensator, Ersatzsicherung
- **Führt zu:** Kapitel 6: Der Akt der Sabotage
