package poker

import (
	"github.com/devplayg/poker/card"
	"sync/atomic"
)

type Player struct {
	ID    int64
	Name  string
	Chip  int64
	cards []*card.Card
}

var playerId int64

func NewPlayer(name string, chip int64) *Player {
	atomic.AddInt64(&playerId, 1)
	return &Player{
		ID:    playerId,
		Name:  name,
		cards: make([]*card.Card, 0),
		Chip:  chip,
	}
}
