# Kapitel 2: Der Schrott-Basar

> Diese Datei spiegelt den aktuell live geschalteten Inhalt von `story_de.json` Stand 2026-09-10 (49 Entscheidungen über 12 Kapitel, Flag-gesteuerte Callbacks, dichtere Einleitungen) — neu generiert, ersetzt den früheren Entwurf von vor der eigentlichen Spielmechanik. Hier oder direkt in `story_de.json` weiterbearbeiten — danach synchronisiert sich nichts mehr automatisch.

**Ort:** Schrott-Basar

Die Batteriereserven schwinden, und du folgst einem handgemalten Schild mit der Aufschrift „GRILL & BATTERIEN“ zu einer fettverschmierten mobilen Festung, betrieben von einem vierarmigen Koch des Intergalaktischen Food-Truck-Kartells. Irgendetwas brutzelt, das vermutlich nicht sollte, und der Rauch riecht schwach nach Ozon und Reue. Er wendet gerade etwas nicht Identifizierbares und begutachtet gleichzeitig die Fusionszellen-Buchse deines Anzugs wie ein Gebrauchtwagenkäufer. Hinter ihm summt eine Kühlbox mit der Aufschrift „KEIN ESSEN (VERMUTLICH)“ bedrohlich vor sich hin. „Frisches Fleisch“, brummt er, nicht unfreundlich. „Kaufen, klauen oder mich amüsieren.“

---

## Entscheidungen

### [1] Ehrlicher Tausch — geborgenen Schrott gegen eine Fusionsbatterie zahlen.

**Ergebnis:** Schrott wechselt den Besitzer. Die Batterie braucht eine langsame Trickle-Ladung, bevor man sie sicher abziehen kann, also wartest du es ab — mit einem ungefragten, strahlend leuchtenden Burger „aufs Haus“. S.T.E.V.E. verbucht ihn als Biogefahr und Snack, in dieser Reihenfolge.

- **Werte:** +25 BAT, GESAMTER SCRAP VERLOREN
- **Gegenstände erhalten:** Strahlungs-Burger
- **Zeitsperre:** 15 Minuten
- **Führt zu:** Kapitel 3: Die Nacht bricht ein

### [2] S.T.E.V.E.-Werbejingle — einen nervigen 1980er-Batterie-Werbespot über die Anzuglautsprecher jagen.

**Ergebnis:** Der Koch hält mitten im Wenden inne, ehrlich begeistert. Er drückt dir eine Batterie als „Werbegeschenk für Fans“ in die Hand und verlangt eine Zugabe. S.T.E.V.E.s Ego, ohnehin schon unkontrollierbar, erreicht neue Höhen.

- **Werte:** +20 BAT
- **Flags gesetzt:** `steve_ego_maxed`
- **Führt zu:** Kapitel 3: Die Nacht bricht ein

### [3] Verdeckter Diebstahl — dich hinter den Truck schlängeln und eine Batterie vom Generator klauen.

**Ergebnis:** Du bekommst die Batterie. Du stolperst dabei auch über einen Kanister, fliegst mit dem Gesicht auf einen Klapptisch und wirst auf der Stelle erkannt. Die Cousins des Kochs schleifen dich als Warnung für andere Kunden zurück zum Wrack und lassen dich — samt deiner gesamten Ausrüstung — genau dort zurück, wo du angefangen hast.

- **Werte:** GESAMTER SCRAP VERLOREN
- **Flags gesetzt:** `wanted_by_cartel`
- **Führt zu:** Kapitel 1: Der Landepunkt

### [4] Der vernünftige Handel — die Hälfte deines Schrotts anbieten und dabei ernst bleiben, um zu sehen, wie weit das reicht.

**Ergebnis:** Der Koch überlegt dein Angebot genau so lange, wie er braucht, um sein Mysteriumsfleisch zu wenden, und nennt dann einen Preis, der etwa doppelt so hoch ist wie deiner. Ihr einigt euch irgendwo unangenehm in der Mitte — weniger Schrott weg als beim ehrlichen Tausch, aber auch weniger Ladung gewonnen. Niemand muss dafür singen.

- **Werte:** +15 BAT, -1 SCRAP
- **Führt zu:** Kapitel 3: Die Nacht bricht ein
