# Chapter 9: The Long March Through the Sandstorms

> This file mirrors the shipped `story.json` content as of 2026-09-10 (49 choices across 12 chapters, flag-gated callbacks, denser intros) — regenerated to replace the earlier pre-mechanics draft. Edit here or in `story.json` directly; nothing keeps the two in sync automatically past this point.

**Location:** The Sandstorm Dunes

A blinding iron sandstorm rolls in on the trek back to the ship, grit scouring paint off every exposed surface. Static builds fast enough that S.T.E.V.E. starts flagging lightning risk every few seconds, with the enthusiasm of a smoke detector that's found a fresh battery. Visibility drops to arm's length; the dunes ahead are just suggestion and noise.

---

## Choices

### [1] Straight through the Crater — march directly into the howling gale.

**Outcome:** You cut the crossing time in half. Grit breaches the neck seal of your suit somewhere around the halfway mark and you feel every grain of it for the rest of the day.

- **Stats:** -3 HULL, -5 BAT
- **Leads to:** Chapter 10: The Repair Crisis

### [2] Cliffside Trail — scale a treacherous ledge out of the direct wind.

**Outcome:** Loose rock slides out from under you twice. You claw your way up by your fingertips both times, gear rattling, and make it over with your dignity in worse shape than your suit.

- **Stats:** -2 BAT
- **Leads to:** Chapter 10: The Repair Crisis

### [3] Hunker in the Dunes — wait out the storm under a steel plate.

**Outcome:** Zero physical damage from the storm itself. But sheltering in one place long enough lets factions on your trail catch up — you emerge to find your ship's camp already ransacked and hostile scouts waiting, forcing a retreat all the way back to the landing site.

- **Stats:** no change
- **Flags set:** `enemies_at_ship`
- **Time-lock:** 240 minutes
- **Leads to:** Chapter 1: The Landing Site

### [4] Cut Through Cartel Supply Lines — you're already a marked face; use it to bluff past a checkpoint hidden in the dunes. _(**unlocks after:** `wanted_by_cartel`)_

**Outcome:** Turns out being known to the cartel cuts both ways — you talk your way past their checkpoint by claiming you're delivering yourself for a bounty collection elsewhere. Bad idea in general, works exactly once. You're through the storm's worst stretch in half the time, adrenaline doing what shelter couldn't.

- **Stats:** -1 HULL, -1 BAT
- **Leads to:** Chapter 10: The Repair Crisis
