package main

import (
	"github.com/devplayg/poker"
	"github.com/devplayg/poker/game"
)

func main() {
	//fmt.Printf("hello pocker")
	deck := poker.NewDeck52()
	deck.Shuffle()
	deck.Distribute()
	deck.Distribute()
	deck.Distribute()
	deck.Distribute()
	poker.Look(deck)
	//deck.Deal()
	//deck.Deal()
	//deck.Deal()
	//deck.Look()
	//deck.Look(false)
	//spew.Dump(deck)

	p1 := poker.NewPlayer("won", 10000)
	p2 := poker.NewPlayer("san", 10000)
	p3 := poker.NewPlayer("soo", 10000)

	config := game.Config{
		Players: []*poker.Player{
			p1, p2, p3,
		},
	}
	game := game.NewTexasHoldem(&config)

	game.Play()

	//game2 := game.NewSevenCardStud(&config)
	//game2.Play()

}
