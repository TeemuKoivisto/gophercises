package deck

type CardSuit int

const (
	DIAMOND CardSuit = iota
	CLUBS
	HEARTS
	SPADES
)

type Card struct {
	number int
	suit   CardSuit
}
