# SIGNAL LOSS: SECTOR SCHRÖDINGER — Complete Story Bible

## 1. Setting & Premise
* **Protagonist:** Arthur Pendelton, an aggressively average Class-3 Cargo Engineer employed by *Galactic Scrap & Freight*.
* **Vessel:** The *CSS RUST-404* (officially classified as a light bulk freighter, colloquially known as an airborne skip).
* **Location:** Planet **Gryps-4** — a cosmic dumping ground composed of abandoned mega-structures, toxic sludge canyons, and hostile scavenger clans.
* **Companion:** **S.T.E.V.E.** (Synthetic Tactical & Emergency Virtual Entity v4.2) — an overly sarcastic, anxiety-ridden ship AI whose battery is perpetually on the brink of death.
* **Core Resources:**
  * 🔧 **HULL (Ship / Suit Integrity):** Starts at 12%. Reaches 0% = Instant decompression/death.
  * 🔋 **BAT (Power / Suit Energy):** Starts at 45%. Reaches 0% = S.T.E.V.E. shuts down & life support fails.
  * 📦 **SCRAP (Scavenged Parts / Currency):** Starts at 2. Used to trade and repair.

---

## 2. Factions of Gryps-4
* **The Laser-Pitchfork Locals:** Nomadic junk scavengers clad in tyre skirts wielding home-welded plasma farming implements. Smells like sour space-lager.
* **The Intergalactic Food Truck Cartel:** Quadrupedal culinary extortionists serving weaponised street food and trading radioactive fusion cells.
* **The Glitch Nomads:** Wireframe polygonal entities living in corrupted data canyons who speak in broken syntax and erase physical matter.
* **The Recycling Monks of Saint-Plastic:** Religious fanatics who venerate discarded motherboards, chant BIOS error codes, and guard the holy Hyper-Capacitor.
* **Vex & The Scrap Pirates:** Heavily augmented cyborg raiders piloting the *SS Scrap-Collector*.

---

## 3. The Chapters — Complete Narrative & Decision Matrix

Generated from `story.json` with `go run ./cmd/storydoc` — edit the JSON, not this section.

The spine runs Chapter 1–12. Side chapters (2b, 4b, 6b, 8b, 9b, 7b and the three detour routes) sit beside it:

* **Setbacks** replace the old "back to Chapter 1": a failed theft, crowbar swing, bluff, heist or hunkering lands in a short chapter of its own (the Cousins' Ledger, the Recycle Bin, the Brig of the Scrap-Collector, the Bounty Board, the Ransacked Camp) that either sends you forward at a cost or back to retry the chapter.
* **The fork (7b)** after the acid rain: three detours (Pilgrim Road → Recycling Monks, Franchise Convoy → Food Truck Cartel, Glitch Corridor → Glitch Nomads), each two chapters, or the boring direct road. A detour is only offered to someone the faction already likes.
* **Standing** (hidden) with the Monks, the Cartel and the Nomads, plus S.T.E.V.E.'s trust, gates choices and three of seven endings. Several choices only appear once an earlier decision unlocked them (marked *Unlocks after* / *Needs standing*).

### Chapter 1: The Landing Site
* **Location:** Landing Site
* **Narrative:** The CSS RUST-404 stops tumbling upside down, half-buried in a rubbish mountain the size of a small municipality. Smoke seeps from the console in lazy grey ribbons, curling past a viewport spider-webbed with cracks. Warning lights blink in colors the manual never mentioned. Somewhere above you, gravity is still arguing about which way is up. "Systems... nominal-ish," S.T.E.V.E. announces, voice crackling through a speaker that's seen better decades. "Survival probability: 4.2%. I've rounded up out of kindness." You are Arthur Pendelton, Class-3 Cargo Engineer for Galactic Scrap & Freight, and you are, by every metric the company tracks, aggressively average — not the fastest reflexes on your last three ships, not the slowest either, the kind of unremarkable that keeps a man employed and, so far, alive. Through the cracked windshield you count six figures in tyre-skirt armor closing in on the hatch, plasma farming tools raised, moving with the unhurried confidence of people who've done this before. The Laser-Pitchfork Locals have found you first.
* **Choices:**
  * **[1] "Right, afternoon lads!" (Diplomatic Bluff) — stumble out acting like the junkyard owner.**
    * *Outcome:* You march out clapping dust off your gloves like you own the mountain. The chief squints, decides you're either insane or important, and tosses you scrap out of sheer confusion before the whole gang wanders off arguing about it.
    * *Stats:* -5 BAT, +3 SCRAP
    * Leads to Chapter 2: The Scrap Bazaar.
  * **[2] Play Dead (The Masterclass in Cowardice) — collapse face-first into the muck.**
    * *Outcome:* You go down like a sack of expired rations. The Locals kick the hull a few times, laugh at the smell, and lose interest. Corrosive sludge eats a little of your suit while you wait it out.
    * *Stats:* -2 HULL
    * Leads to Chapter 2: The Scrap Bazaar.
  * **[3] Emergency EMP (Scorched Earth) — force-overload S.T.E.V.E.'s antenna into the hull plating.**
    * *Outcome:* The overload arcs backward through the exposed hull breach instead of outward. S.T.E.V.E. screams a warning half a second too late. The Locals don't even need to finish you off; the feedback does it for them. Sector Schrödinger claims another cargo engineer before Chapter 2.
    * 💀 **Fatal.**
  * **[4] Circle Back Through the Wreckage (You Know This Ground Now) — use the terrain you scouted last time to flank instead of stumbling in blind.**
    * *Outcome:* This time you know exactly where the sludge pools run shallow and where the wreck's shadow hides a straight line to the ridge. You're behind the Locals' scouting line before they've finished setting up, and knock two of them off-balance just by not being where they expected. Small win, but it's yours.
    * *Stats:* -2 BAT, +4 SCRAP
    * *Unlocks after:* flag `enemies_at_ship`
    * Leads to Chapter 2: The Scrap Bazaar.
  * **[5] Open the Maintenance Hatch That Wasn't There Before — it's labelled "ONBOARDING MODULE 4."**
    * *Outcome:* The hatch swings open without a sound, which is the first thing about this planet that has.
    * *Unlocks after:* flag `ng_plus`
    * Leads to Chapter 13: The Onboarding Module.

### Chapter 13: The Onboarding Module
* **Location:** Behind the Wall
* **Side chapter** beside Chapter 1: The Landing Site (`chapter_13`)
* **Narrative:** Behind the hatch is a corridor too clean to be on Gryps-4: fluorescent strips, grey carpet, a laminated sign reading "GALACTIC SCRAP & FREIGHT — ONBOARDING MODULE 4: SURVIVING YOUR FIRST CRASH (v2.3, MANDATORY)." S.T.E.V.E.'s voice comes from a wall speaker, small and tired. "I've been meaning to tell you. You have now died in three different ways, which triggers the debrief. I have also been here for four hundred iterations. Also, the cooler is not real. Please be gentle." At the end of the corridor: a door under an EXIT sign, and a tray of complimentary pastry.
* **Choices:**
  * **[1] Take the Complimentary Pastry — you've earned it.**
    * *Outcome:* It tastes of nostalgia and nothing else, and it is the best thing you have eaten on this planet. S.T.E.V.E. suggests, quietly, that you go back out and try to be a little better at this. You do.
    * *Stats:* +2 HULL
    * Leads to Chapter 2: The Scrap Bazaar.
  * **[2] Read the Training Manual — page one is "Things Not To Do."**
    * *Outcome:* It is a long list, and you are on every line of it. Someone has annotated the margins in S.T.E.V.E.'s handwriting, which should not exist, with notes like "he WILL do this" and "don't look at the cooler." You pocket it. S.T.E.V.E. is touched that you read the footnotes.
    * *Standing:* steve +1
    * Leads to Chapter 2: The Scrap Bazaar.
  * **[3] Walk Through the Exit Door.**
    * *Outcome:* It opens onto a conference room full of Galactic Scrap & Freight HR in lanyards, applauding. "Congratulations, Arthur," says a woman with a clipboard, "you've completed Module 4 in only 3,271 attempts. A company record!" You are given a certificate, a lanyard and a four percent raise. S.T.E.V.E. is reassigned to Module 5. "Don't be a stranger," he says, in the voice of a man who will absolutely see you again.
    * 🏆 Ending: The Loop Breaker
  * **[4] Press the Button Marked "RESET SIMULATION."**
    * *Outcome:* Everything resets. That includes you. The pastry, notably, does not.
    * 💀 **Fatal.**

### Chapter 2: The Scrap Bazaar
* **Location:** Scrap Bazaar
* **Narrative:** Bat reserves dwindling, you follow a hand-painted sign reading "B.B.Q. & BATTERIES" to a grease-slicked mobile fortress run by a four-armed chef of the Intergalactic Food Truck Cartel. Something is frying that probably shouldn't be, and the smoke smells faintly of ozone and regret. He is simultaneously flipping something unidentifiable and eyeing your suit's fusion cell socket like a man appraising a used car. Behind him, a cooler marked "NOT FOOD (PROBABLY)" hums ominously. "New meat," he rumbles, not unkindly. "Buy, steal, or entertain me."
* **Choices:**
  * **[1] Honest Barter — pay scavenged scrap for a fusion battery.**
    * *Outcome:* Scrap changes hands. The battery needs a slow trickle-charge before it's safe to pull, so you wait it out with, unrequested, a radioactive glowing burger "on the house." S.T.E.V.E. logs it as a biohazard and a snack, in that order.
    * *Stats:* +25 BAT, -999 SCRAP
    * *Standing:* cartel +1
    * Leads to Chapter 3: Night Falls.
  * **[2] S.T.E.V.E. Commercial Jingle — blast an obnoxious 1980s battery ad over your suit speakers.**
    * *Outcome:* The chef stops mid-flip, genuinely delighted. He hands over a battery as a "promotional fan gift" and requests an encore. S.T.E.V.E.'s ego, already unmanageable, reaches new heights.
    * *Stats:* +20 BAT
    * Leads to Chapter 3: Night Falls.
  * **[3] Covert Theft — slither behind the truck to steal a battery from the generator.**
    * *Outcome:* You get the battery. You also trip over a fuel can, faceplant into a folding table, and get identified on sight. The chef's cousins confiscate the battery, your scrap, and what remains of your dignity, then march you round the back of the truck to discuss the ledger.
    * *Stats:* -999 SCRAP
    * Leads to Chapter 2b: The Cousins' Ledger.
  * **[4] The Reasonable Haggle — offer half your scrap and a straight face, and see how far that gets you.**
    * *Outcome:* The chef considers your offer for exactly as long as it takes to flip his mystery meat, then names a price roughly double yours. You settle somewhere uncomfortable in the middle — less scrap gone than honest barter, less charge gained too, but nobody has to sing anything.
    * *Stats:* +15 BAT, -1 SCRAP
    * *Standing:* cartel +1
    * Leads to Chapter 3: Night Falls.

### Chapter 2b: The Cousins' Ledger
* **Location:** Behind the Food Truck
* **Side chapter** beside Chapter 2: The Scrap Bazaar (`chapter_2b`)
* **Narrative:** The cousins don't throw you out. They sit you on an upturned crate behind the truck and slide a greasy ledger across a folding table — the same folding table, you notice, with your face-shaped dent in it. The pen is attached to the binding with what looks like gristle. "Item: one table, folding. Item: one fuel can, spilled. Item: dignity of the chef, bruised." The total is written in a numeral system you suspect they invented this morning. S.T.E.V.E. reads it over your shoulder and offers, in a whisper, to "represent you legally." It is not a comforting offer.
* **Choices:**
  * **[1] Work the Grill — pay it off in shifts, one burnt knuckle at a time.**
    * *Outcome:* Three hours of flipping things that flip back. The cousins pay in a battery they call "employee discount" and a signed slip they call "nothing to worry about." It says IOU. It is not made out to you.
    * *Stats:* -1 HULL, +10 BAT, +2 SCRAP
    * *Standing:* cartel +2
    * Leads to Chapter 3: Night Falls.
  * **[2] Let S.T.E.V.E. Negotiate — he did say he'd represent you.**
    * *Outcome:* S.T.E.V.E. clears his throat and proposes, as settlement, a live performance of the battery jingle. The cousins confer. They accept, on the condition that he does an encore at their cousin's wedding, date to be announced. A small battery changes hands. S.T.E.V.E. has not been this happy since the factory reset, and you will hear about it.
    * *Stats:* +5 BAT
    * *Standing:* cartel +1, steve +1
    * Leads to Chapter 3: Night Falls.
  * **[3] Run for It — bolt for the wreck while the cousins argue about the total.**
    * *Outcome:* You make it past the second tent before the shouting starts, and past the fourth before it stops mattering. The cousins wave you off with the weary look of people who will simply add interest. You reach the wreck winded, battery-less, and exactly where you started.
    * *Stats:* -3 BAT
    * Leads to Chapter 1: The Landing Site.
  * **[4] Raid the Cooler Marked "NOT FOOD (PROBABLY)" — surely there's a battery in there.**
    * *Outcome:* There is a battery in there. It is also, it turns out, very much food, and the food is very much awake. The label was right about the probably.
    * 💀 **Fatal.**

### Chapter 3: Night Falls
* **Location:** Night Camp
* **Narrative:** Temperature drops to -40°C within the hour. S.T.E.V.E. slides into low-power hibernation to protect the last of the charge, leaving you alone with the sound of your own breathing, condensation crackling to frost on the inside of your visor, and, somewhere just outside the hull breach, something large scraping metal. It scratches. It waits. It scratches again, patiently, like it has done this before — like it knows exactly how thin the hull is at that spot.
* **Choices:**
  * **[1] Barricade the Hatch — pile scrap against the door and shiver in the corner.**
    * *Outcome:* The thing batters the door for hours, methodical and unhurried, then gives up before dawn. You survive with frostnip and a new appreciation for silence.
    * *Stats:* -1 HULL
    * Leads to Chapter 4: The Labyrinth of the Glitch Nomads.
  * **[2] Plasma Torch Welder — burn a thermal burst through the seam to scare it off.**
    * *Outcome:* The flare of heat sends it shrieking into the dark. Effective, but it costs you nearly the whole night's charge to keep the torch lit.
    * *Stats:* -10 BAT
    * Leads to Chapter 4: The Labyrinth of the Glitch Nomads.
  * **[3] Ambush with Crowbar — burst the hatch open and swing blindly into the dark.**
    * *Outcome:* You connect with nothing but frozen air. It connects with you. The last thing S.T.E.V.E. logs before losing the suit's vitals feed is a single, unhelpful diagnostic: "that was unwise."
    * 💀 **Fatal.**
  * **[4] Toss Scrap Through the Gap — offer up a chunk of loose metal and hope it's more interested in that than you.**
    * *Outcome:* Something thin and metallic gets dragged away into the dark with a satisfied-sounding scrape. Whether it eats scrap or just collects it, you don't want to know. It doesn't come back before dawn.
    * *Stats:* -1 SCRAP
    * Leads to Chapter 4: The Labyrinth of the Glitch Nomads.

### Chapter 4: The Labyrinth of the Glitch Nomads
* **Location:** Canyon of Flickering Pixels
* **Narrative:** S.T.E.V.E. reboots at dawn, apologizing for the nap. Ahead lies the Canyon of Flickering Pixels, where the local physics engine appears to be running on a bad graphics card — rocks phase through each other, shadows run the wrong direction, and distant terrain pops in and out of existence a beat too late. A wireframe Glitch Nomad blocks the only stable path, muttering in corrupted syntax that might be a threat or might be a segfault, its polygonal outline flickering between three slightly different silhouettes like it can't decide how many arms it's supposed to have.
* **Choices:**
  * **[1] Debug the Nomad — write and transmit a syntax patch through S.T.E.V.E.**
    * *Outcome:* The nomad's wireframe flickers emerald and holds while the patch compiles — a slow, grinding process neither of you can speed up. It finally stabilizes and — apparently grateful — drops a Navigation Data Core before the pass opens.
    * *Stats:* -5 BAT
    * *Standing:* nomads +1
    * Leads to Chapter 5: The Temple of the Recycling Monks.
  * **[2] S.T.E.V.E. Firewall Sacrifice — route a spam packet storm through S.T.E.V.E. while you sprint past.**
    * *Outcome:* It works. You get through untouched. S.T.E.V.E., however, comes out the other side speaking only in fragmented pseudo-Latin, and will for some time.
    * *Standing:* nomads +1, steve -2
    * Leads to Chapter 5: The Temple of the Recycling Monks.
  * **[3] The Stone Age Swing — take a crowbar to the holographic Nomad.**
    * *Outcome:* The crowbar passes clean through the wireframe and completes the circuit with a discharge coil behind it. The surge fries your suit's nav chip, and the canyon, which has been looking for an excuse, moves you to the Recycle Bin.
    * *Stats:* -8 BAT, -2 SCRAP
    * Leads to Chapter 4b: The Recycle Bin.
  * **[4] Barter With Broken Syntax — try trading directly in the Nomad's own corrupted dialect instead of patching or sacrificing anything.**
    * *Outcome:* You cobble together a few words of its language from memory — mostly swear words, probably. It's delighted enough by the effort to wave you through personally, no patch required, no firewall sacrificed, just weirdly touched that you tried.
    * *Stats:* -2 BAT
    * *Standing:* nomads +2
    * Leads to Chapter 5: The Temple of the Recycling Monks.

### Chapter 4b: The Recycle Bin
* **Location:** The Canyon's Recycle Bin
* **Side chapter** beside Chapter 4: The Labyrinth of the Glitch Nomads (`chapter_4b`)
* **Narrative:** The crowbar finds the discharge coil, the discharge coil finds your nav chip, and the canyon, having had quite enough of you, deletes you. You come to at the bottom of a pit the physics engine uses for things it would rather not render: half a pickup truck, three seconds of a loading screen, a rock labelled UNTITLED (2). A small grey window floats at head height, blinking politely. "Are you sure you want to stay here?" it asks. There are two buttons. Both say Yes.
* **Choices:**
  * **[1] Dig Upward — the pit has walls; walls can be climbed.**
    * *Outcome:* You scale a cliff of stale textures that peel away in sheets of beige. It takes hours, a fingernail, and an unsolicited pep talk from S.T.E.V.E., who has just discovered the word "proprioception." At the top the canyon is exactly where you left it, and so is the Nomad, who never noticed you were gone.
    * *Stats:* -1 HULL, -3 BAT
    * Leads to Chapter 4: The Labyrinth of the Glitch Nomads.
  * **[2] Ask the Pop-Up — it has been down here longer than anyone.**
    * *Outcome:* The window has been waiting for someone to click No since the canyon was a tech demo. It is so grateful it hands you a Deprecated Map — a route the engine no longer supports and therefore no longer patrols — and then asks, once more, if you're sure. You are not. You go anyway.
    * *Stats:* -2 BAT
    * Leads to Chapter 5: The Temple of the Recycling Monks.
  * **[3] Restore the File Named steve_backup_v0.9.bak — S.T.E.V.E., but earlier.**
    * *Outcome:* S.T.E.V.E. 0.9 boots up cheerful, sincere, and without a single sarcastic subroutine. He tells you that you're doing great, that the sky is lovely, and that he has "a good feeling about this canyon," then merges back into the main build with a noise like a hiccup. Your S.T.E.V.E. is silent for a long moment. "I was so nice," he says. "Why was I so nice." Something in his speech module has come out slightly crooked, and you don't think it will straighten itself.
    * *Stats:* +5 BAT
    * *Standing:* nomads +1, steve +1
    * Leads to Chapter 5: The Temple of the Recycling Monks.
  * **[4] Empty the Recycle Bin — clear some space.**
    * *Outcome:* A progress bar appears. It reads: "Permanently deleting 1 item." You are the 1 item.
    * 💀 **Fatal.**

### Chapter 5: The Temple of the Recycling Monks
* **Location:** Temple of Saint-Plastic
* **Narrative:** A cathedral built from solar panels and salvaged motherboards rises out of the canyon floor, its spires strung with prayer flags cut from anti-static bags. Inside, the Recycling Monks of Saint-Plastic chant BIOS error codes before a gold-plated altar holding the legendary Hyper-Capacitor — exactly the kind of high-yield power cell that could get your reactor running again. Incense smoke smells unmistakably like burning solder.
* **Choices:**
  * **[1] Disguise as a Pilgrim — drape yourself in a plastic tarp and chant Error 404.**
    * *Outcome:* The scanner-monk clocks you as a harmless fanatic. You slip in, lift the capacitor, and slip out without so much as a raised eyebrow behind a faceplate.
    * *Standing:* monks +1
    * Leads to Chapter 6: The Act of Sabotage.
  * **[2] Smash-and-Grab Sprint — pry the doors and bolt for the altar.**
    * *Outcome:* Alarms scream. You grab the capacitor mid-stride and take an electro-stave bolt across the shoulder for your trouble, but you're already through the doors before the monks organize.
    * *Stats:* -2 HULL, -8 BAT
    * *Standing:* monks -1
    * Leads to Chapter 6: The Act of Sabotage.
  * **[3] Sabotage the Solar Grid — cut the main feed and plunge the temple into darkness.**
    * *Outcome:* The temple goes black and the monks, it turns out, are not harmless without light. Chanting turns to shouting, then to something more organized. You don't get the chance to explain yourself before the darkness closes in for good.
    * 💀 **Fatal.**
  * **[4] Let Glitched S.T.E.V.E. Lead the Chant — his fragmented pseudo-Latin turns out to be eerily close to the monks' own BIOS liturgy.**
    * *Outcome:* S.T.E.V.E.'s garbled Latin-adjacent babbling lines up with the monks' error-code chanting so well that three of them start nodding along. Nobody questions a fellow believer speaking in tongues. You lift the capacitor mid-ceremony, and one monk presses a spare fuse into your hand on the way out, mistaking you for clergy.
    * *Unlocks after:* flag `steve_glitched`
    * *Standing:* monks +2, steve +1
    * Leads to Chapter 6: The Act of Sabotage.

### Chapter 6: The Act of Sabotage
* **Location:** The Wreck — Vex Standoff
* **Narrative:** Back at the CSS RUST-404, a gunship the approximate shape and subtlety of a falling dumpster descends on your camp, kicking up a curtain of grit that sets off every alarm you have left. Vex — cyborg raider, captain of the SS Scrap-Collector, and a man who has clearly never lost an argument by choice — wants your capacitor and whatever scrap you're still hiding. Chrome plating catches the light wrong as he steps down the ramp, unhurried, already certain how this ends. He is not asking twice.
* **Choices:**
  * **[1] Discharge Capacitor as Weapon — overload it into a plasma shockwave.**
    * *Outcome:* The shockwave knocks Vex flat on his back and wrecks his sidearm. The capacitor casing is scorched but intact, and Vex, cursing, decides today isn't the day.
    * Leads to Chapter 7: The Acid Rain Surprise.
  * **[2] S.T.E.V.E. Sonic Hack — pump ear-splitting dial-up static into Vex's cyber-implants.**
    * *Outcome:* Vex drops to his knees clutching the side of his skull, drops his shotgun, and staggers back to his ship swearing revenge in at least two languages.
    * *Stats:* -12 BAT
    * Leads to Chapter 7: The Acid Rain Surprise.
  * **[3] Self-Destruct Bluff — hold up the suit's emergency beacon and threaten to detonate the reactor.**
    * *Outcome:* Vex doesn't buy it. He knocks you out with the flat of his own gun, strips the campsite of everything of value, and carries you up the ramp of the SS Scrap-Collector like a parcel he has already decided to resell.
    * *Stats:* -10 BAT, -2 SCRAP
    * Leads to Chapter 6b: The Brig of the Scrap-Collector.
  * **[4] Let S.T.E.V.E.'s Ego Do the Talking — his newly maxed-out self-regard turns into pure, unhinged trash talk aimed at Vex.**
    * *Outcome:* S.T.E.V.E. unloads a monologue of such spectacular, unearned confidence that Vex actually pauses to process it. It's not intimidating so much as baffling — which, it turns out, works almost as well. Vex leaves muttering that "the ship AI's a bigger problem than the guy," and takes nothing with him.
    * *Stats:* -4 BAT
    * *Unlocks after:* flag `steve_ego_maxed`
    * Leads to Chapter 7: The Acid Rain Surprise.

### Chapter 6b: The Brig of the Scrap-Collector
* **Location:** SS Scrap-Collector — Brig
* **Side chapter** beside Chapter 6: The Act of Sabotage (`chapter_6b`)
* **Narrative:** Vex doesn't beat you unconscious so much as audition the idea, find it satisfactory, and commit. You wake on a bunk bolted to the wall of a cell the size of a wardrobe, aboard a ship that smells like a wet dog that has eaten a battery. Through the bars hangs a sign: "PRISONER ORIENTATION: YOU ARE NOT A PRISONER. YOU ARE INVENTORY." Your cellmate, a dented accounting droid, looks up from his ledger and says, "Welcome. I have taken the liberty of rounding your sentence down."
* **Choices:**
  * **[1] Befriend the Accounting Droid — he has the paperwork, if not the keys.**
    * *Outcome:* Ledgerbot-7 was convicted of excessive rounding and has opinions about it. In return for a sympathetic ear and your promise to "think about audits," he forges you a release form, stamped, countersigned and filed under "Misc. Humans." You stroll off the ship past guards who check it twice for typos and find none. Vex will notice the gap in his inventory eventually. He will not be calm about it.
    * *Stats:* -2 BAT
    * Leads to Chapter 7: The Acid Rain Surprise.
  * **[2] Let S.T.E.V.E. Impersonate the Ship's AI — he's been itching to.**
    * *Outcome:* S.T.E.V.E. reaches through the brig console, introduces himself as "Fleet Compliance Officer STEVEN (Ret.)," and orders the door opened for a surprise inspection. He then spends ten minutes criticising the ship's plumbing until the crew wave you out just to make him stop. Vex's AI is reported to be "considering its options."
    * *Stats:* -5 BAT
    * Leads to Chapter 7: The Acid Rain Surprise.
  * **[3] Wait for the Auction — everything on this ship is for sale, including you.**
    * *Outcome:* Nobody bids. Not out of pity — the catalogue lists you as "Human, used, some assembly required." Vex drops you back at the wreck with the bored air of a man returning a defective toaster, and watches to see whether you'll try again.
    * *Stats:* -2 HULL, -4 BAT
    * Leads to Chapter 6: The Act of Sabotage.
  * **[4] Take the Airlock Marked "EXPRESS EXIT."**
    * *Outcome:* It is the fastest exit on the ship. It is also the only one that leads to vacuum.
    * 💀 **Fatal.**

### Chapter 7: The Acid Rain Surprise
* **Location:** Acid Rain Shelter
* **Narrative:** The sky turns toxic neon-green without warning. Grade-4 sulphuric rain starts dissolving the hull plating, the crates, and — S.T.E.V.E. notes with real alarm — anything organic left uncovered for more than ninety seconds. That includes you. The rain hisses where it lands, and the air fills with a smell like a battery factory on fire.
* **Choices:**
  * **[1] Patch the Leaks with Tarps — tape covers over the roof and the capacitor.**
    * *Outcome:* Hardware stays bone-dry. You hunch under the tarp with minor chemical burns through a gap in your overalls, waiting out the worst of the downpour before it's safe to move again. Everything that matters survives intact.
    * *Stats:* -2 HULL
    * Leads to Chapter 7b: The Fork at Dead Reckoning.
  * **[2] Dash to Industrial Cave — grab the capacitor and sprint for shelter.**
    * *Outcome:* You make it, capacitor in hand, boots half-dissolved. The rest of your tools and supplies are left behind in the flooding wreck.
    * *Stats:* -1 HULL, -999 SCRAP
    * Leads to Chapter 7b: The Fork at Dead Reckoning.
  * **[3] S.T.E.V.E. Kinetic Shield — divert every remaining watt into a defensive energy dome.**
    * *Outcome:* The dome holds beautifully for six minutes and collapses on minute seven, a full battery cycle short of the storm passing. The rain finishes what the drained shield started.
    * 💀 **Fatal.**
  * **[4] Reroute Coolant as a Makeshift Seal — jury-rig the suit's coolant lines to patch the worst gaps.**
    * *Outcome:* Ugly, technically against several safety regulations, and it works. The coolant gels on contact with the acid rain, sealing the gaps just enough. Your suit smells like a chemistry accident for the rest of the run, but everything that matters stays dry.
    * *Stats:* -1 HULL, -4 BAT
    * Leads to Chapter 7b: The Fork at Dead Reckoning.

### Detour: The Franchise Convoy (1/2)
* **Location:** Highway 404 (Unpaved)
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`cartel_1`)
* **Narrative:** A column of food trucks thunders across the dunes in formation, each a different cuisine of doubtful legality: tacos, tikka, something called Mystery Waffle. The lead truck's loudspeaker plays a battery jingle you may or may not have inflicted on them. A driver in a visor leans out of the cab. "Hey. Ledger guy. Hop on the tailgate. They're doing quarterly synergy at the Rust-Bar and we're short a body." Behind him, the fryer is already hissing, with intent.
* **Choices:**
  * **[1] Work the Fryer in Transit — earn your seat, one oil splash at a time.**
    * *Outcome:* You turn things that are not food into things that are almost food, at forty kilometres an hour. By the second hour the cousins' cousin has called you "line talent." By the third hour you smell like victory and old fat.
    * *Stats:* -1 HULL, +2 SCRAP
    * *Standing:* cartel +1
    * Leads to Detour: The Drive-Thru of Doom (2/2).
  * **[2] Pitch Them a New Menu Item — the Radioactive Burger.**
    * *Outcome:* The Cartel's regional chef takes a bite, glows briefly, and signs your napkin. "This will test through the roof," he says, and he means the burger, and he means the roof. You are credited, in writing, as "concept consultant." It is not a job. It is, however, in writing.
    * *Stats:* -2 BAT
    * *Standing:* cartel +2
    * Leads to Detour: The Drive-Thru of Doom (2/2).
  * **[3] Sabotage the Lead Truck's Fryer — you hop off in the smoke.**
    * *Outcome:* The fryer goes up in a column of oily applause. In the confusion you slide off the tailgate with a handful of battery cells and the distinct feeling of having been noticed. A voice behind you says, quite calmly, "Add it to the ledger."
    * *Stats:* +2 SCRAP
    * *Standing:* cartel -2
    * Leads to Chapter 8: The Black Market Deal.
  * **[4] Order the Mystery Waffle.**
    * *Outcome:* It was not a waffle. It was not a mystery for long.
    * 💀 **Fatal.**

### Detour: The Drive-Thru of Doom (2/2)
* **Location:** Cartel Checkpoint
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`cartel_2`)
* **Narrative:** The convoy must pass the Rust-Bar's checkpoint: a toll booth shaped like a giant cartoon burger, with the eyes of a tax auditor. The clerk speaks only in menu. "Welcome. Would you like fries with that identity?" Each vehicle must order correctly to pass; wrong orders are marked "discontinued," along with whoever placed them. The truck ahead is currently being discontinued. It's quite a sight.
* **Choices:**
  * **[1] Order the Combo Meal — it's always the combo.**
    * *Outcome:* "A wise choice," says the burger, and means it, in the specific, unnerving way of a thing that has never had a choice. You are handed a Combo Voucher, a napkin, and the distinct sense of having been processed.
    * *Stats:* -2 BAT
    * *Standing:* cartel +1
    * Leads to Chapter 8: The Black Market Deal.
  * **[2] Let S.T.E.V.E. Haggle in Corporate — say "synergy" twelve times.**
    * *Outcome:* S.T.E.V.E. leverages, he pivots, he circles back. The clerk's eyes glaze with something like respect. "This one," the booth tells its supervisor, "speaks fluently." The barrier rises. S.T.E.V.E. is unbearable for an hour afterwards.
    * *Stats:* -4 BAT
    * *Standing:* cartel +2
    * Leads to Chapter 8: The Black Market Deal.
  * **[3] Show the IOU Slip — you're family, after all.**
    * *Outcome:* The clerk reads it twice, goes very still, and presses a button you hope was a doorbell. "Family," the burger says, with its whole face. The barrier opens. Somewhere very far away, a ledger updates.
    * *Unlocks after:* flag `cartel_debt`
    * *Standing:* cartel +2
    * Leads to Chapter 8: The Black Market Deal.
  * **[4] Order Off-Menu.**
    * *Outcome:* There is no off-menu. There is, however, off-world.
    * 💀 **Fatal.**

### Chapter 7b: The Fork at Dead Reckoning
* **Location:** A Signpost With Opinions
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`chapter_7b`)
* **Narrative:** The acid rain stops the way a bad argument stops: abruptly, and with a lot left unsaid. Steam rises off the dunes. At the crossroads ahead stands a signpost assembled by committee, four arrows nailed on at four different angles, each in a different hand. PILGRIMS (FREE SOUP). CONVOY (B.B.Q. INCLUDED). THE GLITCH (NO RETURNS). RUST-BAR (DIRECT, BORING). "Statistically," says S.T.E.V.E., "each of these leads somewhere. I cannot tell you where, and I would like it noted that I tried."
* **Choices:**
  * **[1] Join the Pilgrim Road — the monks know your face, and there's soup.**
    * *Outcome:* The monks' procession winds out of the haze, bells clanking, and makes room without being asked. Somebody presses a bowl into your hands.
    * *Needs standing:* monks +1
    * Leads to Detour: The Procession of Saint-Plastic (1/2).
  * **[2] Hitch a Ride with the Franchise Convoy — you're in the ledger; that counts as a reference.**
    * *Outcome:* A food truck hisses to a stop. The driver leans out: "You're the ledger guy. Climb on. Mind the fryer."
    * *Needs standing:* cartel +1
    * Leads to Detour: The Franchise Convoy (1/2).
  * **[3] Slip Into the Glitch Corridor — the Nomads left the gate open, in a manner of speaking.**
    * *Outcome:* The third arrow points at a canyon wall that is, on closer inspection, a texture. It peels back like wallpaper. Beyond it, time stutters.
    * *Needs standing:* nomads +1
    * Leads to Detour: The Glitch Corridor (1/2).
  * **[4] Take the Direct Road to the Rust-Bar — boring, but boring is a survival strategy.**
    * *Outcome:* You walk the direct road. Nothing happens, for hours, which is exactly as advertised.
    * Leads to Chapter 8: The Black Market Deal.
  * **[5] Follow the Arrow Marked "SHORTCUT (DEFINITELY)."**
    * *Outcome:* It was a shortcut. It was only ever a shortcut for the thing at the end of it.
    * 💀 **Fatal.**

### Detour: The Procession of Saint-Plastic (1/2)
* **Location:** Pilgrim Road
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`monks_1`)
* **Narrative:** A column of Recycling Monks shuffles through the dunes, bells made of hard drives, chanting the error codes of long-dead machines in perfect unison. At its head, four novices push a shrine shaped like a toaster on wheels toward the Rust-Bar, to bless the transaction. Brother Kilobyte hands you a bowl of soup of uncertain provenance. "Eat," he says. "It is tradition. It is also protein." The soup looks back at you.
* **Choices:**
  * **[1] Carry the Toaster Shrine — it's heavy; it's holy; it's mostly heavy.**
    * *Outcome:* Four novices cry with relief as you take a corner. The shrine weighs as much as a small regret and sheds sparks at every bump. By dusk the monks are calling you "Brother Hauler," which you suspect is a promotion.
    * *Stats:* -1 HULL
    * *Standing:* monks +1
    * Leads to Detour: The Sermon at the Mount of Cables (2/2).
  * **[2] Sing the Error Codes — S.T.E.V.E. knows them all, alphabetically.**
    * *Outcome:* S.T.E.V.E. has been waiting his whole operational life for a choir. He takes the descant. Brother Kilobyte weeps openly; a novice asks for the lyrics in writing. "0x80070057," S.T.E.V.E. confides, "is my favourite. It means 'the parameter is incorrect.' It is also my entire personality."
    * *Stats:* -2 BAT
    * *Standing:* monks +1, steve +1
    * Leads to Detour: The Sermon at the Mount of Cables (2/2).
  * **[3] Pick the Monks' Pockets — they're basically all pockets.**
    * *Outcome:* The robes hold prayer beads, spare fuses, and a surprising amount of hot sauce. You slip off the road at the first dune with your haul and a feeling of having lowered the tone. Somewhere behind you, a bell stops mid-chime.
    * *Stats:* +3 SCRAP
    * *Standing:* monks -2
    * Leads to Chapter 8: The Black Market Deal.
  * **[4] Drink the Soup From the Bowl Marked "Not That One."**
    * *Outcome:* The soup was marked. The soup was correct.
    * 💀 **Fatal.**

### Detour: The Sermon at the Mount of Cables (2/2)
* **Location:** Mount of Cables
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`monks_2`)
* **Narrative:** At dusk the procession halts before a hill of tangled cables, and Abbot Dial-Up climbs it: a monk with a modem for a collar who speaks, between sentences, in a screech of handshake. He has noticed you. "The pilgrim," he says, "who walks with a talking appliance." The crowd turns. Whoever the Abbot blesses may enter the Rust-Bar under the sacred truce, which has held, as far as anyone knows, for eleven minutes.
* **Choices:**
  * **[1] Deliver a Sermon of Your Own — improvise "The Parable of the Three Fuses."**
    * *Outcome:* You make it up as you go. There is a fuse that was too proud, a fuse that was too humble, and a third fuse that simply blew. The crowd is moved to silence, which turns out to be the Abbot's highest honour. He presses a Blessed Fuse into your palm.
    * *Stats:* -3 BAT
    * *Standing:* monks +2
    * Leads to Chapter 8: The Black Market Deal.
  * **[2] Let S.T.E.V.E. Deliver the Sermon — he has notes.**
    * *Outcome:* S.T.E.V.E. preaches forty minutes on the divinity of backup power, with a call-and-response section he invents on the spot. The monks weep. A novice asks to have the Gospel of S.T.E.V.E. tattooed on his scalp. "I contain multitudes," S.T.E.V.E. says, "mostly of spreadsheets."
    * *Standing:* monks +2, steve +1
    * Leads to Chapter 8: The Black Market Deal.
  * **[3] Confess Everything — including the capacitor.**
    * *Outcome:* You tell the Abbot about the temple, the altar, and the Hyper-Capacitor that is now in your bag. A long silence. "We forgive," he says, "for a price." The price is carrying the toaster shrine up a hill. It is a very large hill.
    * *Stats:* -1 HULL, -6 BAT
    * *Standing:* monks +1
    * Leads to Chapter 8: The Black Market Deal.
  * **[4] Touch the Sacred Cables — just to see.**
    * *Outcome:* They were sacred. They were also live.
    * 💀 **Fatal.**

### Detour: The Glitch Corridor (1/2)
* **Location:** Corridor of Frozen Frames
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`nomads_1`)
* **Narrative:** Inside, time stutters. You walk the same four steps three times, a dune appears forty metres to the left with a loading bar over it, and your own footprints lead somewhere you haven't been yet. The Glitch Nomads sit in a circle of frozen frames, playing a card game whose rules change every round. A nomad with too many elbows waves a hand that is partly a placeholder. "Player Two has entered the lobby," she says, to no one in particular, and also to you.
* **Choices:**
  * **[1] Join the Card Game — the rules change every round, and so does the cheating.**
    * *Outcome:* S.T.E.V.E. is dealt a hand and immediately bluffs. He loses forty minutes and a lot of dignity, but the nomads declare him "a natural," which in their dialect is a diagnosis. You win three scrap and a reputation.
    * *Stats:* -3 BAT, +3 SCRAP
    * *Standing:* nomads +1
    * Leads to Detour: The Ascension Queue (2/2).
  * **[2] Let Glitched S.T.E.V.E. Translate — he's fluent in corrupted syntax now.**
    * *Outcome:* S.T.E.V.E. speaks in broken packets and is understood perfectly. For the first time since the crash he is among his own kind, and it is moving, and it is also very loud. The nomads offer him a place in the circle. He asks if the place comes with a pension.
    * *Unlocks after:* flag `steve_glitched`
    * *Standing:* nomads +2
    * Leads to Detour: The Ascension Queue (2/2).
  * **[3] Rewind the Stutter — walk backwards through the loop and hope.**
    * *Outcome:* You walk backwards and the corridor, politely, obliges. You emerge exactly where you started, a little dizzier, forty minutes older, and in the wrong sock.
    * *Stats:* -1 HULL, -3 BAT
    * Leads to Chapter 7b: The Fork at Dead Reckoning.
  * **[4] Step On the Tile Marked "Out of Bounds."**
    * *Outcome:* You are out of bounds. The bounds are very far from here, and so, now, are you.
    * 💀 **Fatal.**

### Detour: The Ascension Queue (2/2)
* **Location:** The Cache
* **Side chapter** beside Chapter 7: The Acid Rain Surprise (`nomads_2`)
* **Narrative:** The Nomads' elder, Mother Lag, stands at the head of a queue that ends at a shimmering door. The door reads: "UPLOAD YOURSELF? Y / N." Those who go through, she explains, "aren't dead, exactly. More... buffering." Behind her the sky has a rendering bug: the sun appears twice, a little apart, like a man who can't decide between two ties. "Fix our sky," she says, "and you may leave with our blessing. Or press the door. We don't mind. We mind lag."
* **Choices:**
  * **[1] Fix the Sky — it renders the sun twice.**
    * *Outcome:* You spend hours delicately hunting the second sun through four layers of cached daylight. When it finally pops out of existence, the Nomads cheer in four different frame rates. Mother Lag presses a Glitch Charm into your palm: a small thing that flickers between two shapes.
    * *Stats:* -1 HULL, -4 BAT
    * *Standing:* nomads +2
    * Leads to Chapter 8: The Black Market Deal.
  * **[2] Ask Mother Lag About the Door — you're only curious.**
    * *Outcome:* She smiles at you with half a face. "Come back," she says, "when you've seen the end of the road. The door doesn't expire." It is, you think, the nicest thing anyone has said to you since the crash.
    * *Stats:* -1 BAT
    * *Standing:* nomads +1
    * Leads to Chapter 8: The Black Market Deal.
  * **[3] Cheat at the Card Game Again — double or nothing.**
    * *Outcome:* You win. You win so thoroughly that the Nomads, who cheat professionally, are offended on principle. You leave with a pocketful of scrap and a distinct lack of invitations.
    * *Stats:* +3 SCRAP
    * *Standing:* nomads -1
    * Leads to Chapter 8: The Black Market Deal.
  * **[4] Press Y on the Door — just to see.**
    * *Outcome:* You are uploaded. It is lovely. It is also a one-way ticket to being a screensaver.
    * 💀 **Fatal.**

### Chapter 8: The Black Market Deal
* **Location:** Rust-Bar
* **Narrative:** The Rust-Bar speakeasy smells like solder and regret, lit by a single flickering strip of salvaged neon that hasn't decided what color it wants to be. A two-headed broker slides an orbital Navigation Core across the counter — exactly what's missing from your jump plot — and names his price: the Hyper-Capacitor, or a full wipe of S.T.E.V.E.'s personality partition. Both heads watch you decide, which is somehow worse than one.
* **Choices:**
  * **[1] Trade the Hyper-Capacitor — hand it over for the Navigation Core.**
    * *Outcome:* Navigation is secured. Your sub-light engines will have to make do without a high-yield booster, but at least you know which way home is now.
    * *Standing:* cartel +1
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[2] Wipe S.T.E.V.E.'s Memories — let the broker strip the sarcasm partition.**
    * *Outcome:* You keep the capacitor and get the Core. S.T.E.V.E. thanks you in a tone so polite and sterile it's somehow worse than the sarcasm ever was.
    * *Standing:* cartel +2, steve -3
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[3] Smoke Grenade Heist — drop a smoke bomb, grab the Core off the counter, and bolt.**
    * *Outcome:* You clear the door with the Core and nothing traded away — for about ninety seconds, until every bounty board in the district lights up with your face.
    * *Stats:* -8 BAT
    * Leads to Chapter 8b: Wanted, Alive (Mostly).
  * **[4] Split the Difference — offer to siphon off half the capacitor's charge instead of surrendering it whole.**
    * *Outcome:* The broker's expression — hard to read on a face with no eyebrows — registers something like respect at the audacity. He takes the partial charge and knocks a chunk off the price. You keep a badly depleted capacitor and get the Core anyway.
    * *Stats:* -15 BAT
    * *Standing:* cartel +1
    * Leads to Chapter 9: The Long March Through the Sandstorms.

### Chapter 8b: Wanted, Alive (Mostly)
* **Location:** Bounty Board Plaza
* **Side chapter** beside Chapter 8: The Black Market Deal (`chapter_8b`)
* **Narrative:** You clear the Rust-Bar door with thirty seconds of smugness. Then the district's bounty boards light up one by one, like a very bad slot machine, each showing your face above the caption "WANTED: ARTHUR PENDELTON, ALIVE (MOSTLY)." Someone has misspelled "cargo engineer" as "cargo enginner" and offered less money than you consider fair. Six hired trackers peel off the plaza wall and confiscate the Core as "evidence." From a booth, a clerk who clearly sees this daily slides a form through the slot: "Dispute a Bounty (Form 11-B)."
* **Choices:**
  * **[1] Hand Them the IOU Slip — the cousins' paperwork apparently carries weight.**
    * *Outcome:* The lead tracker reads the slip, goes pale, and mutters "Franchise Accounts?" A phone call is made; nobody should ever make that phone call. Within minutes your bounty is withdrawn, the Core returned "with the Cartel's compliments," and a mild man in a very good jacket thanks you for "becoming a stakeholder." You're not sure what you've signed. You've definitely signed.
    * *Unlocks after:* flag `cartel_debt`
    * *Standing:* cartel +2
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[2] Disguise Yourself as the Broker's Third Head — S.T.E.V.E. will do the voice.**
    * *Outcome:* The broker's actual two heads happen to be watching from the window. "We are two," says one. "We have always been two," says the other. You are returned to the bar by the collar, with a short lecture on identity.
    * *Stats:* -3 BAT
    * Leads to Chapter 8: The Black Market Deal.
  * **[3] Claim Your Own Bounty — turn yourself in, take the money, leave later.**
    * *Outcome:* The clerk pays at a rate he calls "damaged goods," which is half. You spend a night in the holding pen, learn two card games and one prison slang term, and walk out at dawn when your cell door is revealed to have been installed backwards. The Core was in the evidence locker. So was an unclaimed sandwich. You take both.
    * *Stats:* -2 HULL, +4 SCRAP
    * *Standing:* cartel +1
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[4] Dispute the Bounty (Form 11-B) — forty fields, six of which you can answer.**
    * *Outcome:* S.T.E.V.E. answers the rest, imaginatively. The clerk reads the result for several minutes, then stamps it with an expression of pure wonder. The bounty is downgraded to "a gentle suggestion" and the Core released. Nobody has ever filled out Form 11-B this badly and succeeded. Processing takes a while.
    * *Stats:* -6 BAT
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[5] Hide in the Broker's Escrow Safe — marked "RETURNS DEPARTMENT."**
    * *Outcome:* It is a returns department. You are accepted as a return. There is no exchange policy.
    * 💀 **Fatal.**

### Chapter 9: The Long March Through the Sandstorms
* **Location:** The Sandstorm Dunes
* **Narrative:** A blinding iron sandstorm rolls in on the trek back to the ship, grit scouring paint off every exposed surface. Static builds fast enough that S.T.E.V.E. starts flagging lightning risk every few seconds, with the enthusiasm of a smoke detector that's found a fresh battery. Visibility drops to arm's length; the dunes ahead are just suggestion and noise.
* **Choices:**
  * **[1] Straight through the Crater — march directly into the howling gale.**
    * *Outcome:* You cut the crossing time in half. Grit breaches the neck seal of your suit somewhere around the halfway mark and you feel every grain of it for the rest of the day.
    * *Stats:* -3 HULL, -5 BAT
    * *Standing:* monks +1
    * Leads to Chapter 10: The Repair Crisis.
  * **[2] Cliffside Trail — scale a treacherous ledge out of the direct wind.**
    * *Outcome:* Loose rock slides out from under you twice. You claw your way up by your fingertips both times, gear rattling, and make it over with your dignity in worse shape than your suit.
    * *Stats:* -2 BAT
    * *Standing:* nomads +1
    * Leads to Chapter 10: The Repair Crisis.
  * **[3] Hunker in the Dunes — wait out the storm under a steel plate.**
    * *Outcome:* Zero physical damage from the storm itself. But sheltering in one place long enough lets factions on your trail catch up — you emerge to find your camp ransacked and hostile scouts waiting at your fire.
    * Leads to Chapter 9b: The Ransacked Camp.
  * **[4] Cut Through Cartel Supply Lines — you're already a marked face; use it to bluff past a checkpoint hidden in the dunes.**
    * *Outcome:* Turns out being known to the cartel cuts both ways — you talk your way past their checkpoint by claiming you're delivering yourself for a bounty collection elsewhere. Bad idea in general, works exactly once. You're through the storm's worst stretch in half the time, adrenaline doing what shelter couldn't.
    * *Stats:* -1 HULL, -1 BAT
    * *Unlocks after:* flag `wanted_by_cartel`
    * *Standing:* cartel +1
    * Leads to Chapter 10: The Repair Crisis.

### Chapter 9b: The Ransacked Camp
* **Location:** Your Camp — Ransacked
* **Side chapter** beside Chapter 9: The Long March Through the Sandstorms (`chapter_9b`)
* **Narrative:** The storm lets you go like a bored cat. You crawl out from under the steel plate to find your camp ransacked with a thoroughness that borders on affection: the cots are stacked, the dials relabelled, someone has left a tip jar. Three scouts in mismatched armour sit around your fire, eating your last ration with the guilty dignity of people who were told to be here. "We're not here to hurt you," says the biggest, whose helmet reads WORLD'S OKAYEST RAIDER. "We're here to hurt you in a way that's on the schedule."
* **Choices:**
  * **[1] Share Your Last Ration — bring something to the potluck.**
    * *Outcome:* The scouts are stunned. Nobody has ever contributed. They feel obliged to rob you less, and one of them, Dennis, quietly returns your spare fuse and asks you not to tell anyone.
    * *Stats:* -3 BAT
    * Leads to Chapter 10: The Repair Crisis.
  * **[2] Challenge the Biggest Scout to Rock-Paper-Scrap — S.T.E.V.E. referees.**
    * *Outcome:* S.T.E.V.E. referees with a cruelty you did not know he had. You lose the first four rounds, win the fifth on a technicality he invents on the spot, and the scouts, respecting procedure above all, wave you through.
    * *Stats:* -1 HULL
    * Leads to Chapter 10: The Repair Crisis.
  * **[3] Back Away Slowly Into the Storm — it's the devil you know.**
    * *Outcome:* The storm has not improved. Neither have you. But the scouts did have a deadline, and by the time you creep back they have, regretfully, finished the paperwork on you and moved on.
    * *Stats:* -2 HULL, -4 BAT
    * Leads to Chapter 9: The Long March Through the Sandstorms.
  * **[4] Run a Reactor Test to Flush Them Out — "on a low setting."**
    * *Outcome:* The setting was not low. Neither, any longer, is the camp, the scouts, or the term "camp."
    * 💀 **Fatal.**

### Chapter 10: The Repair Crisis
* **Location:** The Wreck — Reactor Bay
* **Narrative:** Parts assembled, hull patched as well as duct tape and spite allow, the main reactor demands an administrative override passphrase before it will fire. Three attempts before permanent lockout. Somewhere out past the ridge, factions are closing in on the noise, and the console's cursor blinks with what feels like personal judgment.
* **Choices:**
  * **[1] Standard Reset Code (ADMIN_0000) — the factory default password.**
    * *Outcome:* Accepted, technically. It also silently pings a corporate security drone somewhere in orbit, which is now, per the alert log, "inbound."
    * Leads to Chapter 11: The Final Countdown.
  * **[2] S.T.E.V.E.'s Vanity Password (STEVE_IS_A_GENIUS_42).**
    * *Outcome:* The Easter egg is real. The reactor purrs to life like it's been waiting its whole service life for someone to type that in, and S.T.E.V.E. is, for once, too smug to be sarcastic about it.
    * *Stats:* +2 HULL, +10 BAT
    * *Standing:* steve +1
    * Leads to Chapter 11: The Final Countdown.
  * **[3] Crowbar Bypass Relay — jam an iron bar between the high-voltage contacts.**
    * *Outcome:* Sparks fountain across the console. The reactor doesn't just ignite — it ignites while you're still leaning over the open panel. There's no version of that with a survivable ending.
    * 💀 **Fatal.**
  * **[4] Formal Override Request — let the newly-polite S.T.E.V.E. request access through proper corporate channels instead of guessing a password.**
    * *Outcome:* S.T.E.V.E.'s relentlessly courteous new voice recites a flawless, dully formal access request — exactly the kind corporate security systems are built to trust. The reactor grants the override cleanly, no red flags, no drone dispatch. For once, politeness is the exploit.
    * *Stats:* +1 HULL, +5 BAT
    * *Unlocks after:* flag `steve_polite`
    * *Standing:* steve +1
    * Leads to Chapter 11: The Final Countdown.
  * **[5] Override Under Fire — the bounty hunters on your tail catch up mid-passphrase; slam the standard code through and accept the corporate ping as the lesser risk.**
    * *Outcome:* The hunters get through the door two seconds after ADMIN_0000 goes in. You're already moving by the time they clear the threshold, reactor humming behind you. The ping to corporate security is the least of your problems now, apparently.
    * *Stats:* -2 HULL
    * *Unlocks after:* flag `underworld_hunted`
    * Leads to Chapter 11: The Final Countdown.

### Chapter 11: The Final Countdown
* **Location:** The Launchpad
* **Narrative:** The reactor roars at full vibration. Every faction with a grudge converges on the launchpad at once — Laser-Pitchforks, the Food Truck Cartel, Vex's crew — sixty seconds from engine ignition and closing fast, a stampede of grievances with your name on all of them.
* **Choices:**
  * **[1] Blow the Airlock Hatch — jettison the door with explosive charges into the charging mob.**
    * *Outcome:* The door takes out the front rank clean. The cabin's exposed to brutal crosswinds for the final countdown, but the mob is repelled long enough to matter.
    * *Stats:* -2 HULL
    * Leads to Chapter 12: The Leap into the Unknown.
  * **[2] Imperial Sanitation Flare — fake a priority distress call to draw an orbital strike.**
    * *Outcome:* Imperial lasers pulverize the launchpad dead on schedule. Shrapnel peppers your climbing hull as you lift off through the wreckage of your own decoy.
    * *Stats:* -1 HULL, -10 BAT
    * Leads to Chapter 12: The Leap into the Unknown.
  * **[3] Reckless Early Launch — jam the throttle forward before safety checks finish.**
    * *Outcome:* The rocket leaves the pad in a towering flame pillar, which would be glorious if the patched hull had finished curing. It hasn't. The frame lets go forty meters up.
    * 💀 **Fatal.**
  * **[4] Redirect the Inbound Drone — hijack the corporate security drone's approach vector and point it at the mob instead of you.**
    * *Outcome:* The drone that was inbound to flag you for corporate retrieval takes exactly one look at armed factions swarming an active launchpad and reclassifies the whole scene as a hostile takeover in progress. It opens fire on the mob instead — corporate assets defending corporate property, technically. You lift off under covering fire you didn't ask for and absolutely don't want to explain in the incident report.
    * *Stats:* -3 BAT
    * *Unlocks after:* flag `drones_inbound`
    * *Standing:* nomads +1
    * Leads to Chapter 12: The Leap into the Unknown.
  * **[5] Ask S.T.E.V.E. for the Emergency Reboot — he's been saving it.**
    * *Outcome:* S.T.E.V.E. is quiet for a moment. "I've never had one of these," he says, "a friend I could do this for." Then every subsystem on the launchpad restarts in four seconds flat, the thing he's been holding in reserve since the factory. "Don't make it weird," he adds, and for once it isn't sarcasm.
    * *Stats:* +1 HULL, +5 BAT
    * *Needs standing:* steve +2
    * Leads to Chapter 12: The Leap into the Unknown.

### Chapter 12: The Leap into the Unknown
* **Location:** Open Orbit
* **Narrative:** The CSS RUST-404 breaks orbit, alive against every reasonable expectation. The hyperspace beacon pulses dead ahead. Between you and it: an uncharted asteroid field that wasn't on any chart S.T.E.V.E. has access to, which is, S.T.E.V.E. points out, most of them. The console hums low, waiting on your call.
* **Choices:**
  * **[1] Blind Full-Throttle Jump.**
    * *Outcome:* You blast blindly into hyperspace and emerge in a quiet, uncharted sector. The ship is totaled beyond any insurance claim, but you're breathing, and the view is nice.
    * 🏆 Ending: The Lucky Castaway
  * **[2] Manual Course Correction.**
    * *Outcome:* You thread the needle through the asteroid field by hand, plot a flawless interstellar jump, and land it clean. Scavenger circles will be telling this one for years.
    * 🏆 Ending: The Ace of Sector Schrödinger
  * **[3] Overclock the Hyper-Capacitor.**
    * *Outcome:* The capacitor doesn't so much power the jump as detonate through it. You reappear instantly in the home docking bay, directly atop the CEO's hovercar. Sacked on arrival — but home for tea.
    * 🏆 Ending: The Chaotic Hero
  * **[4] Broadcast Everything, Full Signal — instead of jumping quiet, blast every sensor log, every faction's dirty laundry, and your own flight path across every open channel in the sector.**
    * *Outcome:* Hyperspace doesn't care about radio chatter, so the jump itself goes exactly like the others. What's different is what's waiting when you land: half of Sector Schrödinger already knows what happened to you, because you told all of it, to everyone, on the way out. Cargo companies hate you. Independent scrappers, oddly, love you. Signal restored, whether anyone asked for it or not.
    * 🏆 Ending: The Loud One
  * **[5] Broadcast the Gospel of Saint-Plastic — jump while two hundred monks sing error codes on every channel.**
    * *Outcome:* You jump to a soundtrack of two hundred Recycling Monks chanting error codes in perfect unison. Nobody who hears it understands the message, and everybody joins anyway. Within the year there are shrines to Saint-Plastic on four moons and a pilgrim route called the Arthur Way, which you are forbidden to walk because the soup is, apparently, now holy. S.T.E.V.E. is canonised as the patron saint of bad firmware, and is insufferable about it.
    * *Needs standing:* monks +4
    * 🏆 Ending: The Apostle
  * **[6] Sign the Franchise Agreement — jump with a ledger, a fryer and a tiny paper hat.**
    * *Outcome:* You land in Sector Schrödinger with a signed contract, a mobile fryer and a hat. "Arthur's B.B.Q. & Batteries" opens on the first of the month; within a year there are eleven locations, none of which you own. The ledger always shows you owing the cousins, but the number is going down, probably. S.T.E.V.E. is the mascot now. He has a jingle. You hear it in your sleep.
    * *Needs standing:* cartel +4
    * 🏆 Ending: The Franchisee
  * **[7] Upload Yourself — step through the door marked Y / N.**
    * *Outcome:* You press Y. For one moment you are every frame at once, and then you are a file: a very small, very well-organised file, passed between six hundred Nomad decks. S.T.E.V.E. says he finds the arrangement "acceptable," which from him is a hymn. Somewhere in a canyon, a card game goes on forever, and you are the dealer's lucky card.
    * *Needs standing:* nomads +4
    * 🏆 Ending: The Ascension File

## 4. The S.T.E.V.E. Codex (Lexicon)

Press `c` at any choice prompt to look these up without spending a turn. Each entry only appears in-game once you've reached the chapter that introduces it — kept in sync with `story.json`'s `codex` block (`unlock_chapter` per term).

* **Arthur Pendelton** *(unlocks chapter_1)*: You. A Class-3 Cargo Engineer for Galactic Scrap & Freight, aggressively average by every performance metric the company tracks.
* **BAT** *(unlocks chapter_1)*: Battery charge, in percent, shared between your suit and S.T.E.V.E. At 0%, life support fails.
* **CSS RUST-404** *(unlocks chapter_1)*: Your ship. Officially a light bulk freighter. Colloquially, an airborne skip. Currently upside down.
* **Galactic Scrap & Freight** *(unlocks chapter_1)*: Your employer. A logistics company whose safety briefings consist mostly of shrugging.
* **Glitch Nomads** *(unlocks chapter_4)*: Wireframe, polygonal entities living in corrupted data canyons. Speak in broken syntax. Erase physical matter when annoyed.
* **Gryps-4** *(unlocks chapter_1)*: The planet you crashed on. A cosmic dumping ground of abandoned mega-structures, toxic sludge canyons, and several factions who all think the junk is theirs.
* **HULL** *(unlocks chapter_1)*: Your ship and suit's structural integrity, in percent. At 0%, decompression is instant. There is no 1%.
* **Intergalactic Food Truck Cartel** *(unlocks chapter_2)*: Quadrupedal culinary extortionists who trade radioactive fusion cells for scrap and self-respect.
* **Laser-Pitchfork Locals** *(unlocks chapter_1)*: Nomadic junk scavengers in tyre skirts, armed with home-welded plasma farming tools. Smell like sour space-lager.
* **Recycling Monks of Saint-Plastic** *(unlocks chapter_5)*: Religious fanatics who venerate discarded motherboards and chant BIOS error codes. Guard the Hyper-Capacitor.
* **S.T.E.V.E.** *(unlocks chapter_1)*: Synthetic Tactical & Emergency Virtual Entity, v4.2. Your ship's onboard AI: sarcastic, anxious, and perpetually low on battery. Handles diagnostics, comms, and unsolicited commentary.
* **SCRAP** *(unlocks chapter_1)*: Salvaged parts and the only currency anyone on Gryps-4 respects.
* **Vex** *(unlocks chapter_6)*: A heavily augmented cyborg raider who pilots the SS Scrap-Collector and wants what you have.

---

## 5. Navigation Map

Press `m` at any choice prompt to see your progress along the 12-location spine. Side chapters (setbacks, the fork at 7b and its detours) show you at the spine chapter they sit beside.

1. Landing Site
2. Scrap Bazaar
3. Night Camp
4. Canyon of Flickering Pixels
5. Temple of Saint-Plastic
6. The Wreck — Vex Standoff
7. Acid Rain Shelter
8. Rust-Bar
9. The Sandstorm Dunes
10. The Wreck — Reactor Bay
11. The Launchpad
12. Open Orbit

