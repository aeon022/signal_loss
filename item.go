package main

type Item struct {
	name string
	description string
	position Point
	resourceType string
	resourceAmount int
}

func (i *Item) Use() {
	if i.resourceType != "" {
		game.player.resources[i.resourceType] += i.resourceAmount
	}
}
