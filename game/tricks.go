package game

import (
	"github.com/prakashdumbre-toast/304/Deck"
	"github.com/prakashdumbre-toast/304/player"
)

const (
	J = 21
	N = 20
	A = 11
	T = 10
	Q = 3
	K = 2
	S = 0
	E = 0
)

type TrickCard struct {
	Card   *Deck.Card
	Player *player.Player
}

type Trick struct {
	Leader      *player.Player
	Winner      *player.Player
	TrickCards  []TrickCard
	WinningCard *Deck.Card
	LeadingCard *Deck.Card
	Trump       string
}

func TrickDecision(T *Trick) {
	cardRanks := map[string]int{
		"JACK":  21,
		"9":     20,
		"ACE":   11,
		"TEN":   10,
		"QUEEN": 3,
		"KING":  2,
		"7":     0,
		"8":     0,
	}

	wCard := T.LeadingCard
	wPlayer := T.Leader

	if T.Trump == T.LeadingCard.Face {
		//	Get the card with the highest rank , discard other faces
		for _, c := range T.TrickCards {
			if cardRanks[c.Card.Face] > cardRanks[wCard.Face] && c.Card.Face == T.LeadingCard.Face {
				wCard = c.Card
				wPlayer = c.Player
			}
		}
	} else {
		//	Add +10 if trump card and discard other Suits
		for _, c := range T.TrickCards {
			currCardRank := cardRanks[c.Card.Face]
			if c.Card.Suit == T.Trump {
				currCardRank += 10
			}
			if currCardRank > cardRanks[wCard.Face] && c.Card.Face == T.LeadingCard.Face {
				wCard = c.Card
				wPlayer = c.Player
			}
		}
	}

	T.WinningCard = wCard
	T.Winner = wPlayer
}
