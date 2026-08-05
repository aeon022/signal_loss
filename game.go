package main

import "time"

type Game struct {
	player *Player
	enemies []*Enemy
	items []*Item
	dayTime bool
	startTime time.Time
}

func (g *Game) Start() {
	// game loop goes here...
}

func (g *Game) UpdateDayNightCycle() {
	elapsed := time.Since(g.startTime)
	if elapsed > 12*time.Hour {
		g.dayTime = !g.dayTime
		g.startTime = time.Now()
	}
}
