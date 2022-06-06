package poker

import (
	"github.com/devplayg/poker/card"
	"math/rand"
	"time"
)

type Deck52 struct {
	cardsInDeck []*card.Card
	dealtCards  []*card.Card
}

func (d *Deck52) CardsInDeck() []*card.Card {
	return d.cardsInDeck
}

func (d *Deck52) DealtCards() []*card.Card {
	return d.dealtCards
}

func (d *Deck52) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cardsInDeck), func(i, j int) {
		d.cardsInDeck[i], d.cardsInDeck[j] = d.cardsInDeck[j], d.cardsInDeck[i]
	})
}

func (d *Deck52) Distribute() *card.Card {
	c := d.cardsInDeck[0]
	d.cardsInDeck = d.cardsInDeck[1:]
	d.dealtCards = append(d.dealtCards, c)
	return c
}

func NewDeck52() Deck {
	cards := make([]*card.Card, card.Count52, card.Count52)
	for i := card.Club; i <= card.Spade; i++ {
		for j := card.Two; j <= card.Ace; j++ {
			idx := (i * card.ColCount) + j
			cards[idx] = &card.Card{card.Suit(i), j}
		}
	}
	return &Deck52{
		cardsInDeck: cards,
	}
}
