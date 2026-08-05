package main

type Player struct {
	health int
	position Point
	resources map[string]int
	flashlight bool
}

func (p *Player) Move(direction string) {
	// move player based on direction...
}

func (p *Player) UseFlashlight() {
	if p.flashlight && !game.dayTime {
		// use flashlight...
	}
}
