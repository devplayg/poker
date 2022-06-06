package game

import (
	"fmt"
	"github.com/devplayg/poker"
)

type SevenCardStud struct {
	name           string
	config         *Config
	minPlayerCount int
	maxPlayerCount int
}

func (g *SevenCardStud) Play() error {
	fmt.Printf("%s is about to play", g.name)
	return nil
}

func NewSevenCardStud(config *Config) poker.Game {
	return &TexasHoldem{
		name:           "7 Card Stud",
		config:         config,
		minPlayerCount: 2,
		maxPlayerCount: 8,
	}
}
