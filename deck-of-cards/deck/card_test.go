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

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Diamond}
	if cards[0] != exp {
		t.Error("Expected Ace of Diamonds to be first card in a default sort. Received:", cards[0])
	}
}

func TestRegularSort(t *testing.T) {
	cards := New()
	exp := Card{Rank: Ace, Suit: Diamond}
	if cards[0] != exp {
		t.Error("Expected Ace of Diamonds to be first card of a regular deck. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected 3 Jokers, received: %d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Ten
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Ten {
			t.Error("Expected to have filtered all twos and tens")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 4*13*3 {
		t.Errorf("Expected to have created 3 decks of size %d, received %d", 4*13*3, len(cards))
	}
}
