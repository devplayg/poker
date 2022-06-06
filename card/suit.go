package card

const (
	Club = iota
	Diamond
	Heart
	Spade
)

var suits = map[Suit]string{
	Club:    "Club",
	Diamond: "Diamond",
	Heart:   "Heart",
	Spade:   "Spade",
}

type Suit int

func (s Suit) Name() string {
	suit, _ := suits[s]
	return suit
}

func (s Suit) Value() int {
	return int(s)
}
