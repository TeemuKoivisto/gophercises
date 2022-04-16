package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Ten, Suit: Diamond})
	fmt.Println(Card{Rank: Queen, Suit: Heart})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Spades
	// Ten of Diamonds
	// Queen of Hearts
	// King of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 4*13 {
		t.Error("Incorrect deck size for a new deck")
	}
}

func TestRegularSort(t *testing.T) {
	cards := New()
	exp := Card{Rank: Ace, Suit: Diamond}
	if cards[0] != exp {
		t.Error("Expected Ace of Diamonds to be first card of a regular deck. Received:", cards[0])
	}
}
