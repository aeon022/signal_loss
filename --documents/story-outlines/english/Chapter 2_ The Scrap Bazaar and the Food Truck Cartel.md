# Chapter 2: The Scrap Bazaar

> This file mirrors the shipped `story.json` content as of 2026-09-10 (49 choices across 12 chapters, flag-gated callbacks, denser intros) — regenerated to replace the earlier pre-mechanics draft. Edit here or in `story.json` directly; nothing keeps the two in sync automatically past this point.

**Location:** Scrap Bazaar

Bat reserves dwindling, you follow a hand-painted sign reading "B.B.Q. & BATTERIES" to a grease-slicked mobile fortress run by a four-armed chef of the Intergalactic Food Truck Cartel. Something is frying that probably shouldn't be, and the smoke smells faintly of ozone and regret. He is simultaneously flipping something unidentifiable and eyeing your suit's fusion cell socket like a man appraising a used car. Behind him, a cooler marked "NOT FOOD (PROBABLY)" hums ominously. "New meat," he rumbles, not unkindly. "Buy, steal, or entertain me."

---

## Choices

### [1] Honest Barter — pay scavenged scrap for a fusion battery.

**Outcome:** Scrap changes hands. The battery needs a slow trickle-charge before it's safe to pull, so you wait it out with, unrequested, a radioactive glowing burger "on the house." S.T.E.V.E. logs it as a biohazard and a snack, in that order.

- **Stats:** +25 BAT, ALL SCRAP LOST
- **Items gained:** Radioactive Burger
- **Time-lock:** 15 minutes
- **Leads to:** Chapter 3: Night Falls

### [2] S.T.E.V.E. Commercial Jingle — blast an obnoxious 1980s battery ad over your suit speakers.

**Outcome:** The chef stops mid-flip, genuinely delighted. He hands over a battery as a "promotional fan gift" and requests an encore. S.T.E.V.E.'s ego, already unmanageable, reaches new heights.

- **Stats:** +20 BAT
- **Flags set:** `steve_ego_maxed`
- **Leads to:** Chapter 3: Night Falls

### [3] Covert Theft — slither behind the truck to steal a battery from the generator.

**Outcome:** You get the battery. You also trip over a fuel can, faceplant into a folding table, and get identified on sight. The chef's cousins drag you back to the wreck as a warning to other customers, dumping you — and none of your gear — back where you started.

- **Stats:** ALL SCRAP LOST
- **Flags set:** `wanted_by_cartel`
- **Leads to:** Chapter 1: The Landing Site

### [4] The Reasonable Haggle — offer half your scrap and a straight face, and see how far that gets you.

**Outcome:** The chef considers your offer for exactly as long as it takes to flip his mystery meat, then names a price roughly double yours. You settle somewhere uncomfortable in the middle — less scrap gone than honest barter, less charge gained too, but nobody has to sing anything.

- **Stats:** +15 BAT, -1 SCRAP
- **Leads to:** Chapter 3: Night Falls
