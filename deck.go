package poker

import (
	"fmt"
	"github.com/devplayg/poker/card"
)

type Deck interface {
	CardsInDeck() []*card.Card
	Shuffle()
	DealtCards() []*card.Card
	Distribute() *card.Card
}

func Look(d Deck) {
	cards := [4][13]*card.Card{}

	for i := card.Club; i <= card.Spade; i++ {
		for j := card.Two; j <= card.Ace; j++ {
			cards[i][j] = &card.Card{card.Suit(i), j}
		}
	}

	for _, c := range d.DealtCards() {
		cards[c.Suit][c.Value] = nil
	}

	for i := card.Club; i <= card.Spade; i++ {
		fmt.Printf("[%-7s]", card.Suit(i).Name())
		for j := card.Two; j <= card.Ace; j++ {
			if cards[i][j] == nil {
				fmt.Printf("%2s ", "_")
			} else {
				fmt.Printf("%2d ", cards[i][j].Value)
			}
		}
		fmt.Println()
	}
}
