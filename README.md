# SIGNAL LOSS: SECTOR SCHRÖDINGER

[![Created by abteilung83](https://img.shields.io/badge/Crafted%20by-abteilung83-black.svg)](https://github.com/aeon022)

A retro-futuristic sci-fi text adventure for the terminal. You crash-land on Gryps-4 — a cosmic
dumping ground — with a sarcastic, anxious, perpetually-low-on-battery ship AI and twelve
chapters' worth of ways to make things worse. Real branching: every chapter but the finale has a
choice that's a genuine dead end (instant death, or a setback that sends you back to the start
with something lost). Decisions echo forward, too — several choices only appear once an earlier
one unlocked them. Four possible endings.

Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea)/[Lipgloss](https://github.com/charmbracelet/lipgloss).
Also runs in a browser — the exact same engine, compiled to WebAssembly, no reimplementation.
See `roadmap.md` for how that works and everything else that went into this.

## Play

**In your browser: https://aeon022.github.io/signal_loss/** — no install, same Go engine
compiled to WebAssembly.

Or on the command line:

```sh
go install github.com/aeon022/signal_loss/cmd/signal_loss@v0.1.0
signal_loss
```

Or clone and build:

```sh
git clone https://github.com/aeon022/signal_loss.git
cd signal_loss
./install.sh
./signal_loss
```

### Flags

- `-fast` — skip the typewriter effect and compress time-locks to a couple of seconds
- `-realtime` — make time-locks real minutes/hours instead of compressed seconds; the wait is
  saved to disk and enforced on next launch too, `-fast` can't skip it
- `-lang en|de` — game language, UI chrome and story content (default `en`); the browser build
  has the same choice as a launch-screen toggle
- `-save PATH` — save file location (default `save.json`)
- `-story PATH` — story data file (default `story.json` for `en`, `story_de.json` for `de`)

### Controls

`1`–`9` choose · `↑↓` + `Enter` navigate · `c` codex · `m` map · `q` quit

## Development

```sh
go build ./...
go vet ./...
```

`cmd/wasm/` is a separate Go module (its own `go.mod`) — it pins a small, audited fork of Bubble
Tea that adds the couple of platform hooks (`tty_js.go`, `signals_js.go`) needed to compile for
`GOOS=js`, which isn't supported upstream. The native CLI here is unaffected; it uses real,
unmodified Bubble Tea.

## Support

If you enjoyed getting stranded on Gryps-4: [support the project on Polar](https://buy.polar.sh/polar_cl_CbYo27mWKgPdiEv3IJS680uCrzqDus7LWzd131V74y0).

## License

MIT — see `LICENSE`.

## 🏢 About abteilung83

Engineered and maintained by **abteilung83** — developer tools, terminal UIs, and the occasional
sci-fi text adventure. [github.com/aeon022](https://github.com/aeon022)
