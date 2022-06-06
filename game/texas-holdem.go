package game

import "github.com/devplayg/poker"

type TexasHoldem struct {
	name           string
	config         *Config
	minPlayerCount int
	maxPlayerCount int
}

func (t *TexasHoldem) Play() error {
	return nil
}

func NewTexasHoldem(config *Config) poker.Game {
	return &TexasHoldem{
		name:           "Texas hold'em",
		config:         config,
		minPlayerCount: 2,
		maxPlayerCount: 8,
	}
}

//
//func NewTexasHoldem(config *Config) *poker.Game {
//	return &TexasHoldem{
//		config: config,
//	}
//}
/*
BTN
SB
BB
UTG
HJ
CO


BTN
SB
BB
UTG
UTG+1
UTG+2(MP)
UTG+3(MP2)
UTG+4(MP3, LJ)
HJ
CO

BTN
SB
BB
UTG
MP
CO

*/
