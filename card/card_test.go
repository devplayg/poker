package card

import (
	"github.com/devplayg/poker"
	"testing"
)

func TestNewDeck52(t *testing.T) {
	deck := poker.NewDeck52()
	if len(deck.CardsInDeck()) != 52 {
		t.Error("New52CardDeck size must be 52")
	}
}

func TestDeck52_Deal(t *testing.T) {
	deck := poker.NewDeck52()
	deck.Distribute()
	deck.Distribute()
	deck.Distribute()
	deck.Distribute()

	want := 52 - 4
	if len(deck.CardsInDeck()) != want {
		t.Errorf("len(deck.CardsInDeck()) = %d; want %d", len(deck.CardsInDeck()), want)
	}
}
