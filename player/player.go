package player

import (
	"github.com/prakashdumbre-toast/304/Deck"
)

type Player struct {
	Name         string
	Score        int
	HoldingCards []*Deck.Card
}
