package card

const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	RowCount = 4
	ColCount = 13
	Count52  = RowCount * ColCount
)

type Card struct {
	Suit  Suit
	Value int
}
