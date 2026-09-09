You are an expert Systems Engineer and Senior Go Developer with a deep love for retro sci-fi interactive fiction. Your task is to build a fully functional, highly immersive, terminal-based text adventure game in Go (using pure standard library) titled "SIGNAL LOSS: SECTOR SCHRODINGER".

================================================================================
1. GAME OVERVIEW & THEME
================================================================================
- Title: SIGNAL LOSS: SECTOR SCHRODINGER
- Genre: Interactive Sci-Fi Terminal Adventure (inspired by "Lifeline" and "Perry Rhodan").
- Protagonist: Arthur Pendelton, a low-tier cargo ship engineer crashed on the junk planet Gryps-4.
- Companion: S.T.E.V.E. v4.2 (Shipboard Technical Emergency Verbal Entity), a panicked, highly sarcastic suit AI with a pink-flickering HUD.
- Visual Style: Retro Shipboard CLI UI using ANSI escape codes:
  * Cyan: S.T.E.V.E. dialogue & system status
  * Yellow/Orange: Warnings, resource adjustments (-5 BAT, +3 SCRAP)
  * Pink/Red: Critical alarms, alien threats, hull damage
  * Green: Successes, items acquired, engine start

================================================================================
2. TECHNICAL REQUIREMENTS & GO ARCHITECTURE
================================================================================
Write clean, idiomatic, modular Go code. Stick strictly to the Go standard library (no external TUI dependencies required to guarantee zero-dependency building).

Project Package Structure:
cmd/signal_loss/
├── main.go               # Entry point, flag parsing, game loop, save loading/saving
pkg/
├── game/
│   └── state.go          # GameState struct, persistence (JSON), time calculation
├── ui/
│   └── terminal.go       # ANSI colors, typewriter effect, input reader, ASCII banner
└── story/
    └── chapters.go       # Chapter data structure, 12 chapters, choice logic & effects

Key Features:
1. CLI Flags:
   - Provide a `--fast` or `-f` flag to bypass the typewriter effect for rapid testing/debugging.
2. Typewriter Effect & Input Handling:
   - Implement `printSlow(text string, delay time.Duration)`.
   - Implement robust input parsing using `bufio.Reader`.
   - If the player inputs invalid options (not 1, 2, or 3), trigger S.T.E.V.E. to output a random, funny sarcastic insult before re-prompting.
3. State Management & Persistence:
   - `GameState` struct:
     * Hull (int, default 12%)
     * Bat (int, default 45%)
     * Scrap (int, default 2)
     * CurrentChapter (int, 1 through 12)
     * LastTimestamp (int64, Unix time for real-world wait times)
     * Inventory (map[string]bool or struct flags for items like HyperCapacitor, NavModule)
   - Save automatically after every chapter to `.signal_loss_save.json` in the current working directory.
4. Real-World Time Locks (Timer Feature):
   - In Chapter 3 (Overnight Cooldown), record `LastTimestamp`. If the player launches the game before 2 hours (or a configurable short duration for testing), show a standby screen where S.T.E.V.E. is asleep and force the user to wait or return later.

================================================================================
3. STORY & CHAPTER OUTLINE (12 CHAPTERS)
================================================================================
Implement 12 detailed chapters with rich prose, 3 choices each, and stat mutations:

Chapter 1: The Crash Site (Sector Schrodinger)
- Context: Waking up in a burning pod embedded in a giant junk mountain. Laser-pitchfork aliens approaching.
- Choices: [1] Bluff "Hello boys!", [2] Play dead in the mud, [3] Fire emergency EMP (drains battery/damages hull).

Chapter 2: The Scrap Bazaar & Food Truck Cartel
- Context: Encountering a heavy-armored food truck run by a 4-armed alien chef selling a fusion battery.
- Choices: [1] Trade scrap for battery + radioactive burger, [2] Sneak and steal battery, [3] Make S.T.E.V.E. sing an embarrassing radio jingle.

Chapter 3: The Cold Night (Time-Lock Chapter)
- Context: Temperatures plunge to -40°C. S.T.E.V.E. goes into low-power standby mode.
- Choices: [1] Barricade the hatch, [2] Use torch beam against a stalker, [3] Attack blind with an iron bar.
- Mechanic: Enforces a real-time wait check upon resumption.

Chapter 4: The Canyon of Glitching Pixels
- Context: Traversing a canyon populated by Glitch Nomads made of holographic digital noise.
- Choices: [1] Debug the nomad's code via terminal, [2] Melee attack through holograms, [3] Sacrificial firewall route using S.T.E.V.E.

Chapter 5: The Temple of the Recycling Monks
- Context: A cathedral of solar panels housing the Hyper-Capacitor on an altar.
- Choices: [1] Disguise as a trash pilgrim & chant BIOS error codes, [2] Smash and grab sprint, [3] Sabotage main solar power supply.

Chapter 6: The Pirate Interception
- Context: Pirate captain "Vex" lands his heavy freighter to steal your newly acquired gear.
- Choices: [1] Overload capacitor blast at Vex, [2] Frequency noise hack on Vex's helmet, [3] Self-destruct bluff.

Chapter 7: Acid Rain Emergency
- Context: Green toxic rain falls, eating through hull metal.
- Choices: [1] Seal leaks with tarps and suit pieces, [2] Grab capacitor and run to a nearby cavern, [3] Overclock S.T.E.V.E.'s emergency shield.

Chapter 8: The Black Market Deal
- Context: Meeting a two-headed broker in Engine Alley who holds the FTL Navigation Module.
- Choices: [1] Trade away the Hyper-Capacitor, [2] Trade away S.T.E.V.E.'s joke logs & personality data, [3] Throw smoke grenade and steal the module.

Chapter 9: The Sandstorm Trek
- Context: Iron-dust storm threatens to destroy suit seals on the trek back to the ship.
- Choices: [1] Sprint through the crater wind-tunnel, [2] Climb narrow cliff edge, [3] Huddle under metal plate and wait it out.

Chapter 10: The Repair Crisis
- Context: Wiring components together; reactor asks for an unknown admin password.
- Choices: [1] Type default code "ADMIN_0000", [2] Manual bypass with iron rod (causes fire/sparks), [3] Enter S.T.E.V.E.'s joke password ("STEVE_IS_A_GENIUS_42").

Chapter 11: The Final Countdown
- Context: Hostile factions rush the launch pad as engines prime.
- Choices: [1] Explosive hatch jettison, [2] Dangerous early engine throttle up, [3] Imperial cleaner emergency beacon ping.

Chapter 12: FTL Leap into the Unknown (The Finale)
- Context: Breaking atmosphere into an unmapped asteroid belt.
- Choices: [1] Blind full-throttle FTL leap, [2] Manual joystick navigation through asteroids, [3] Overcharge capacitor quantum jump.
- Endings: Provide 3 distinct epilogue screens based on Choice [1], [2], or [3].

================================================================================
4. INSTRUCTIONS FOR CODE GENERATION
================================================================================
- Generate full, executable Go code.
- Ensure proper string formatting and multi-line string handling.
- Do not leave missing function stubs. Implement state loading/saving, terminal rendering, typewriter delay, dynamic status bar printing (`[HULL: 12% | BAT: 45% | SCRAP: 2]`), and chapter transition handling.

Begin implementation now!