package main

import (
	"kingdomino/src/gamestate"
)

func main() {
	width := 1280
	height := 720

	var players = []string{"Bob", "Alice"}

	game := gamestate.NewGame(width, height, players)
	RegisterEventListeners()

	// Never return
	done := make(chan struct{})

	game.TileDeck.ShuffleDeck()

	<-done
}
