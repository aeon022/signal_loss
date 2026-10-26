// storydoc prints the chapter section of story.md from story.json, so the
// story bible never drifts from the game: go run ./cmd/storydoc > chapters.md
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aeon022/signal_loss/pkg/game"
	"github.com/aeon022/signal_loss/pkg/story"
)

func main() {
	st, err := story.LoadStory("story.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Spine in order, each followed by the side chapters that sit beside it.
	var ids []string
	for _, id := range game.ChapterOrder {
		ids = append(ids, id)
		var side []string
		for sid, c := range st.Chapters {
			if c.Spine == id {
				side = append(side, sid)
			}
		}
		sort.Strings(side)
		ids = append(ids, side...)
	}
	for _, id := range ids {
		ch := st.Chapters[id]
		fmt.Printf("### %s\n* **Location:** %s\n", ch.Title, ch.Location)
		if ch.Spine != "" {
			fmt.Printf("* **Side chapter** beside %s (`%s`)\n", st.Chapters[ch.Spine].Title, id)
		}
		fmt.Printf("* **Narrative:** %s\n* **Choices:**\n", ch.Content)
		for i, c := range ch.Choices {
			fmt.Printf("  * **[%d] %s**\n    * *Outcome:* %s\n", i+1, c.Text, c.Outcome)
			if m := c.StatMutations; m.Hull != 0 || m.Bat != 0 || m.Scrap != 0 {
				fmt.Printf("    * *Stats:* %s\n", stats(m))
			}
			if c.RequiresFlag != "" {
				fmt.Printf("    * *Unlocks after:* flag `%s`\n", c.RequiresFlag)
			}
			if len(c.RequiresRep) > 0 {
				fmt.Printf("    * *Needs standing:* %s\n", fmtMap(c.RequiresRep))
			}
			if len(c.Reputation) > 0 {
				fmt.Printf("    * *Standing:* %s\n", fmtMap(c.Reputation))
			}
			switch {
			case c.Fatal:
				fmt.Println("    * 💀 **Fatal.**")
			case c.Ending != "":
				fmt.Printf("    * %s\n", c.Ending)
			case c.NextChapter != "":
				fmt.Printf("    * Leads to %s.\n", st.Chapters[c.NextChapter].Title)
			}
		}
		fmt.Println()
	}
}

func stats(m story.StatMutations) string {
	var p []string
	for _, s := range []struct {
		n string
		v int
	}{{"HULL", m.Hull}, {"BAT", m.Bat}, {"SCRAP", m.Scrap}} {
		if s.v != 0 {
			p = append(p, fmt.Sprintf("%+d %s", s.v, s.n))
		}
	}
	return strings.Join(p, ", ")
}

func fmtMap(m map[string]int) string {
	var p []string
	for k, v := range m {
		p = append(p, fmt.Sprintf("%s %+d", k, v))
	}
	sort.Strings(p)
	return strings.Join(p, ", ")
}
