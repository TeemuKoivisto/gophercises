package blackjack

import (
	"github.com/TeemuKoivisto/gophercises/deck-of-cards/deck"
)

type BlackJackStatus uint8

const (
	NOT_STARTED BlackJackStatus = iota
	PLAYER_TURN
	DEALER_TURN
	PLAYER_WON
	DEALER_WON
	DRAW
)

type PlayerType int

const (
	HUMAN PlayerType = iota
	AI
)

type Player struct {
	Type PlayerType
	Name string
}

type GameState struct {
	Status      BlackJackStatus
	PlayerCards Hand
	DealerCards Hand
}

type GameOpts struct {
	decks int
}

type Game struct {
	decks  int
	state  GameState
	cards  []deck.Card
	player Player
	dealer Player
}

func New(opts GameOpts) *Game {
	decks := opts.decks
	if decks == 0 {
		decks = 3
	}
	return &Game{
		decks: decks,
		state: GameState{
			Status: NOT_STARTED,
		},
		cards: deck.New(deck.Deck(decks), deck.Shuffle),
	}
}

func (g *Game) setPlayer(t PlayerType, name string) {
	g.player = Player{
		Type: t,
		Name: name,
	}
}

func (g *Game) setDealer(t PlayerType, name string) {
	g.dealer = Player{
		Type: t,
		Name: name,
	}
}

func (g *Game) startGame() {
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&g.state.PlayerCards, &g.state.DealerCards} {
			card, cards := draw(g.cards)
			*hand = append(*hand, card)
			g.cards = cards
		}
	}
	g.state.Status = PLAYER_TURN
}

func (g *Game) handlePlayerHit() {
	card, cards := draw(g.cards)
	g.state = GameState{
		Status:      PLAYER_TURN,
		PlayerCards: append(g.state.PlayerCards, card),
		DealerCards: g.state.DealerCards,
	}
	g.cards = cards
}

func (g *Game) handlePlayerStand() {
	g.state = GameState{
		Status:      DEALER_TURN,
		PlayerCards: g.state.PlayerCards,
		DealerCards: g.state.DealerCards,
	}
}

func (g *Game) handleDealerHit() {
	card, cards := draw(g.cards)
	g.state = GameState{
		Status:      DEALER_TURN,
		PlayerCards: g.state.PlayerCards,
		DealerCards: append(g.state.DealerCards, card),
	}
	g.cards = cards
}

func (g *Game) endGame() {
	pScore, dScore := g.state.PlayerCards.Score(), g.state.DealerCards.Score()
	var result BlackJackStatus
	switch {
	case pScore > 21:
		result = PLAYER_WON
	case dScore > 21:
		result = DEALER_WON
	case pScore > dScore:
		result = PLAYER_WON
	case dScore > pScore:
		result = DEALER_WON
	case pScore == dScore:
		result = DRAW
	}
	g.state = GameState{
		Status:      result,
		PlayerCards: g.state.PlayerCards,
		DealerCards: g.state.DealerCards,
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
