package blackjack

type DealerAI struct {
	name string
}

func newDealerAI(name string) *DealerAI {
	return &DealerAI{
		name,
	}
}

func (ai *DealerAI) getAIMove(state *GameState) string {
	dScore := Score(state.DealerCards...)
	if dScore <= 16 || (dScore == 17 && Soft(state.DealerCards...)) {
		return "h"
	} else {
		return "s"
	}
}
