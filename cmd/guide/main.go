// guide prints the endings guide (Markdown) straight from the story files, so it can never drift
// from the game: go run ./cmd/guide -lang en > guide.md
// The winning lines are pinned in pkg/game's TestEveryEndingHasAWinningLine; the same prefixes are
// replayed on the English story to find each choice's position, then walked on the chosen language.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
)

var spine = []string{`"Right, afternoon`, "Honest Barter", "Barricade the Hatch", "Debug the Nomad", "Disguise as a Pilgrim", "Discharge Capacitor"}
var finale = []string{"Trade the Hyper-Capacitor", "Straight through the Crater", "Standard Reset Code", "Blow the Airlock Hatch"}

func line(mid string, route ...string) []string {
	l := append(append([]string{}, spine...), "Patch the Leaks", mid)
	return append(append(l, route...), finale...)
}

// order is the order the endings appear in the guide.
var order = []string{"lucky_castaway", "ace", "chaotic_hero", "loud_one", "apostle", "franchisee", "ascension_file", "loop_breaker"}
var lines = map[string][]string{
	"lucky_castaway": append(line("Take the Direct Road"), "Blind Full-Throttle"),
	"ace":            append(line("Take the Direct Road"), "Manual Course Correction"),
	"chaotic_hero":   append(line("Take the Direct Road"), "Overclock"),
	"loud_one":       append(line("Take the Direct Road"), "Broadcast Everything"),
	"apostle":        append(line("Join the Pilgrim Road", "Carry the Toaster Shrine", "Deliver a Sermon"), "Broadcast the Gospel"),
	"franchisee":     append(line("Hitch a Ride", "Work the Fryer", "Order the Combo Meal"), "Sign the Franchise"),
	"ascension_file": append(line("Slip Into the Glitch", "Join the Card Game", "Fix the Sky"), "Upload Yourself"),
	"loop_breaker":   {`"Right, afternoon`, "Covert Theft", "Run for It", "Open the Maintenance Hatch", "Walk Through the Exit"},
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// positions replays a line of prefixes on the English story and returns each choice's index.
func positions(en *story.StoryData, prefixes []string) []int {
	s := game.NewGameState(en)
	s.Seed, s.Flags["ng_plus"] = -1, true
	var idx []int
	for _, p := range prefixes {
		found := false
		for i, c := range game.AvailableChoices(s, en.Chapters[s.CurrentChapter]) {
			if strings.HasPrefix(c.Text, p) {
				idx, found = append(idx, i), true
				game.ResolveChoice(s, c, en)
				break
			}
		}
		if !found {
			fail(fmt.Errorf("no choice %q in %s", p, s.CurrentChapter))
		}
	}
	return idx
}

func main() {
	lang := flag.String("lang", "en", "en or de")
	flag.Parse()
	en, err := story.LoadStory("story.json")
	if err != nil {
		fail(err)
	}
	st := en
	if *lang == "de" {
		if st, err = story.LoadStory("story_de.json"); err != nil {
			fail(err)
		}
	}
	t := map[string]map[string]string{
		"en": {"title": "Signal Loss: The Endings Guide", "warn": "**Spoilers.** This guide gives away every ending. Play the game first, come back when you are stuck or curious.", "how": "How to use it", "howText": "Each ending has one winning line: the numbered choices, chapter by chapter, with luck switched off. Your own run can differ a little where a choice is a gamble.", "ends": "The eight endings", "items": "Items and where to find them", "item": "Item", "where": "Found in", "use": "Used for", "saver": "Saves you once from a fatal choice", "gate": "Unlocks a shortcut", "tips": "Good to know", "step": "Step", "chapter": "Chapter", "pick": "Pick", "hidden": "Hidden ending (New Game+)", "road": "The shared road (chapters 1-7)", "roadText": "Seven endings start the same way. Follow these steps first, then continue with the ending you want. Numbers in [brackets] are the choice numbers on screen.", "tipsT": "Good to know", "tip1": "- **Satellite data:** choosing *Play Dead* in chapter 1 (or *Barricade the Hatch* in chapter 3) gives S.T.E.V.E. a data module. From then on you see the route map, and ⚠ marks the choices he flags as deadly.\n- **Luck:** a few choices are a gamble, fixed per run. The lines above were tested with luck switched off; if a gamble goes against you, start a new run.\n- **Hidden ending:** three different endings unlock New Game+ and a secret chapter 13. The short line above is the fast way once it is open.\n- **Items:** the three savers each rescue you once from a fatal choice; keep them for the last chapters."},
		"de": {"title": "Signal Loss: Der Enden-Guide", "warn": "**Spoiler.** Dieser Guide verrät jedes Ende. Spiel erst, komm zurück, wenn du feststeckst oder neugierig bist.", "how": "So benutzt du ihn", "howText": "Jedes Ende hat eine Gewinnlinie: die nummerierten Entscheidungen, Kapitel für Kapitel, ohne Glück. Dein eigener Lauf kann etwas abweichen, wo eine Entscheidung ein Glücksspiel ist.", "ends": "Die acht Enden", "items": "Gegenstände und wo man sie findet", "item": "Gegenstand", "where": "Gefunden in", "use": "Wofür", "saver": "Rettet einmal vor einer tödlichen Wahl", "gate": "Schaltet eine Abkürzung frei", "tips": "Gut zu wissen", "step": "Schritt", "chapter": "Kapitel", "pick": "Wahl", "hidden": "Verstecktes Ende (New Game+)", "road": "Der gemeinsame Weg (Kapitel 1-7)", "roadText": "Sieben Enden beginnen gleich. Folge zuerst diesen Schritten und dann dem Ende, das du willst. Die Zahlen in [Klammern] sind die Nummern der Wahl auf dem Bildschirm.", "tipsT": "Gut zu wissen", "tip1": "- **Satellitendaten:** Wer in Kapitel 1 *Sich tot stellen* wählt (oder in Kapitel 3 *Die Luke verbarrikadieren*), gibt S.T.E.V.E. ein Datenmodul. Ab dann siehst du die Routenkarte, und ⚠ markiert die Wahlen, die er als tödlich meldet.\n- **Glück:** Ein paar Entscheidungen sind ein Glücksspiel, pro Lauf festgelegt. Die Linien oben sind ohne Glück getestet; geht ein Glücksspiel gegen dich aus, starte einen neuen Lauf.\n- **Verstecktes Ende:** Drei verschiedene Enden öffnen New Game+ und ein geheimes Kapitel 13. Die kurze Linie oben ist der schnelle Weg, sobald es offen ist.\n- **Gegenstände:** Die drei Lebensretter bewahren dich je einmal vor einer tödlichen Wahl; spar sie für die letzten Kapitel auf."},
	}[*lang]

	fmt.Printf("---\ntitle: %q\n---\n\n%s\n\n## %s\n\n%s\n\n", t["title"], t["warn"], t["how"], t["howText"])

	// each line as text steps; the first stretch is shared by all endings but the hidden one
	walk := func(id string) ([]string, string) {
		s := game.NewGameState(st)
		s.Seed, s.Flags["ng_plus"] = -1, true
		var steps []string
		var ending string
		for _, k := range positions(en, lines[id]) {
			ch := st.Chapters[s.CurrentChapter]
			c := game.AvailableChoices(s, ch)[k]
			steps = append(steps, fmt.Sprintf("**%s** — [%d] %s", ch.Title, k+1, c.Text))
			game.ResolveChoice(s, c, st)
			if c.Ending != "" {
				ending = c.Ending
			}
		}
		return steps, ending
	}
	all := map[string][]string{}
	names := map[string]string{}
	for _, id := range order {
		all[id], names[id] = walk(id)
	}
	common := all[order[0]]
	for _, id := range order[:len(order)-1] {
		n := 0
		for n < len(common) && n < len(all[id]) && common[n] == all[id][n] {
			n++
		}
		common = common[:n]
	}
	num := func(steps []string, from int) string {
		var out []string
		for i, x := range steps {
			out = append(out, fmt.Sprintf("%d. %s", from+i+1, x))
		}
		return strings.Join(out, "\n")
	}
	fmt.Printf("## %s\n\n%s\n\n%s\n\n", t["road"], t["roadText"], num(common, 0))
	fmt.Printf("## %s\n\n", t["ends"])
	for n, id := range order {
		name := names[id]
		if i := strings.Index(name, ": "); i >= 0 && strings.HasPrefix(name, "🏆") {
			name = name[i+2:]
		}
		h := fmt.Sprintf("%d. %s", n+1, name)
		if id == "loop_breaker" {
			fmt.Printf("### %s (%s)\n\n%s\n\n", h, t["hidden"], num(all[id], 0))
			continue
		}
		fmt.Printf("### %s\n\n%s\n\n", h, num(all[id][len(common):], len(common)))
	}

	// items: first place each is handed out, and what they do
	type src struct{ chapter, choice string }
	found := map[string]src{}
	var ids []string
	for id := range st.Chapters {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		for _, c := range st.Chapters[id].Choices {
			for _, it := range c.ItemsAdded {
				if _, ok := found[it]; !ok {
					found[it] = src{st.Chapters[id].Title, c.Text}
				}
			}
		}
	}
	use := map[string]string{}
	for _, id := range ids {
		for _, c := range st.Chapters[id].Choices {
			if c.SavedByItem != "" {
				use[c.SavedByItem] = t["saver"]
			}
			if c.RequiresItem != "" && use[c.RequiresItem] == "" {
				use[c.RequiresItem] = t["gate"]
			}
		}
	}
	var itemNames []string
	for it := range use {
		itemNames = append(itemNames, it)
	}
	sort.Strings(itemNames)
	fmt.Printf("## %s\n\n| %s | %s | %s |\n|---|---|---|\n", t["items"], t["item"], t["where"], t["use"])
	for _, it := range itemNames {
		f := found[it]
		fmt.Printf("| %s | %s — %s | %s |\n", it, f.chapter, strings.ReplaceAll(f.choice, "|", "/"), use[it])
	}
	fmt.Printf("\n## %s\n\n%s\n", t["tipsT"], t["tip1"])
}
