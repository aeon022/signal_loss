# Chapter 8: The Black Market Deal

> This file mirrors the shipped `story.json` content as of 2026-09-10 (49 choices across 12 chapters, flag-gated callbacks, denser intros) — regenerated to replace the earlier pre-mechanics draft. Edit here or in `story.json` directly; nothing keeps the two in sync automatically past this point.

**Location:** Rust-Bar

The Rust-Bar speakeasy smells like solder and regret, lit by a single flickering strip of salvaged neon that hasn't decided what color it wants to be. A two-headed broker slides an orbital Navigation Core across the counter — exactly what's missing from your jump plot — and names his price: the Hyper-Capacitor, or a full wipe of S.T.E.V.E.'s personality partition. Both heads watch you decide, which is somehow worse than one.

---

## Choices

### [1] Trade the Hyper-Capacitor — hand it over for the Navigation Core.

**Outcome:** Navigation is secured. Your sub-light engines will have to make do without a high-yield booster, but at least you know which way home is now.

- **Stats:** no change
- **Items gained:** Nav-Core
- **Items lost:** Hyper-Capacitor
- **Leads to:** Chapter 9: The Long March Through the Sandstorms

### [2] Wipe S.T.E.V.E.'s Memories — let the broker strip the sarcasm partition.

**Outcome:** You keep the capacitor and get the Core. S.T.E.V.E. thanks you in a tone so polite and sterile it's somehow worse than the sarcasm ever was.

- **Stats:** no change
- **Items gained:** Nav-Core
- **Flags set:** `steve_polite`
- **Leads to:** Chapter 9: The Long March Through the Sandstorms

### [3] Smoke Grenade Heist — drop a smoke bomb, grab the Core off the counter, and bolt.

**Outcome:** You clear the door with the Core and nothing traded away — for about ninety seconds, until every bounty board in the district lights up with your face. Hired trackers chase you clean out of the district and all the way back to the wreck, and you lose the Core proving it.

- **Stats:** -8 BAT
- **Items lost:** Nav-Core
- **Flags set:** `underworld_hunted`
- **Leads to:** Chapter 1: The Landing Site

### [4] Split the Difference — offer to siphon off half the capacitor's charge instead of surrendering it whole.

**Outcome:** The broker's expression — hard to read on a face with no eyebrows — registers something like respect at the audacity. He takes the partial charge and knocks a chunk off the price. You keep a badly depleted capacitor and get the Core anyway.

- **Stats:** -15 BAT
- **Items gained:** Nav-Core
- **Leads to:** Chapter 9: The Long March Through the Sandstorms
