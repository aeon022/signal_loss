# Chapter 10: The Repair Crisis

> This file mirrors the shipped `story.json` content as of 2026-09-10 (49 choices across 12 chapters, flag-gated callbacks, denser intros) — regenerated to replace the earlier pre-mechanics draft. Edit here or in `story.json` directly; nothing keeps the two in sync automatically past this point.

**Location:** The Wreck — Reactor Bay

Parts assembled, hull patched as well as duct tape and spite allow, the main reactor demands an administrative override passphrase before it will fire. Three attempts before permanent lockout. Somewhere out past the ridge, factions are closing in on the noise, and the console's cursor blinks with what feels like personal judgment.

---

## Choices

### [1] Standard Reset Code (ADMIN_0000) — the factory default password.

**Outcome:** Accepted, technically. It also silently pings a corporate security drone somewhere in orbit, which is now, per the alert log, "inbound."

- **Stats:** no change
- **Flags set:** `drones_inbound`
- **Leads to:** Chapter 11: The Final Countdown

### [2] S.T.E.V.E.'s Vanity Password (STEVE_IS_A_GENIUS_42).

**Outcome:** The Easter egg is real. The reactor purrs to life like it's been waiting its whole service life for someone to type that in, and S.T.E.V.E. is, for once, too smug to be sarcastic about it.

- **Stats:** +2 HULL, +10 BAT
- **Leads to:** Chapter 11: The Final Countdown

### [3] Crowbar Bypass Relay — jam an iron bar between the high-voltage contacts.

**Outcome:** Sparks fountain across the console. The reactor doesn't just ignite — it ignites while you're still leaning over the open panel. There's no version of that with a survivable ending.

- **Stats:** no change
- 💀 **FATAL — run ends here.**

### [4] Formal Override Request — let the newly-polite S.T.E.V.E. request access through proper corporate channels instead of guessing a password. _(**unlocks after:** `steve_polite`)_

**Outcome:** S.T.E.V.E.'s relentlessly courteous new voice recites a flawless, dully formal access request — exactly the kind corporate security systems are built to trust. The reactor grants the override cleanly, no red flags, no drone dispatch. For once, politeness is the exploit.

- **Stats:** +1 HULL, +5 BAT
- **Leads to:** Chapter 11: The Final Countdown

### [5] Override Under Fire — the bounty hunters on your tail catch up mid-passphrase; slam the standard code through and accept the corporate ping as the lesser risk. _(**unlocks after:** `underworld_hunted`)_

**Outcome:** The hunters get through the door two seconds after ADMIN_0000 goes in. You're already moving by the time they clear the threshold, reactor humming behind you. The ping to corporate security is the least of your problems now, apparently.

- **Stats:** -2 HULL
- **Flags set:** `drones_inbound`
- **Leads to:** Chapter 11: The Final Countdown
