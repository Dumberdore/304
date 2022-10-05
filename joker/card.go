package joker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Invalid is the text shown when a valid face or suit is not found.
const Invalid = "?"

// Card defines a playing card with a face and suit.
type Card struct {
	Face CardFace
	Suit CardSuit
}

// Parse parses a card identifier to construct a card.
func Parse(identifier string) (Card, bool) {
	identifier = strings.TrimSpace(strings.ToUpper(identifier))
	l := len(identifier)
	if l != 2 && l != 3 {
		return Card{}, false
	}
	var cardFace CardFace
	switch identifier[:l-1] {
	case "A":
		cardFace = FaceAce
	case "2":
		cardFace = Face2
	case "3":
		cardFace = Face3
	case "4":
		cardFace = Face4
	case "5":
		cardFace = Face5
	case "6":
		cardFace = Face6
	case "7":
		cardFace = Face7
	case "8":
		cardFace = Face8
	case "9":
		cardFace = Face9
	case "10":
		cardFace = Face10
	case "J":
		cardFace = FaceJack
	case "Q":
		cardFace = FaceQueen
	case "K":
		cardFace = FaceKing
	case "!":
		cardFace = FaceJoker
	default:
		return Card{}, false
	}
	var cardSuit CardSuit
	switch identifier[l-1:] {
	case "H":
		cardSuit = SuitHearts
	case "D":
		cardSuit = SuitDiamonds
	case "C":
		cardSuit = SuitClubs
	case "S":
		cardSuit = SuitSpades
	case "!":
		cardSuit = SuitJoker
	default:
		return Card{}, false
	}
	return Card{cardFace, cardSuit}, true
}

// Value returns the numeric value of a card.
func (c Card) Value() int {
	return (int(c.Face) * 100) + int(c.Suit)
}

// Equal returns whether both cards have the same face and suit.
func (c Card) Equal(b Card) bool {
	return b.Face == c.Face && b.Suit == c.Suit
}

// Identifier returns a machine-readable representation of a card.
func (c Card) Identifier() string {
	if c.Suit.Name() == Invalid || c.Suit.Name() == Invalid {
		return "??"
	}
	var faceidentifier string
	if c.Face == FaceJoker {
		faceidentifier = "!"
	} else if c.Face == Face10 {
		faceidentifier = c.Face.Name()
	} else {
		faceidentifier = string(c.Face.Name()[0])
	}
	var suitidentifier string
	if c.Suit == SuitJoker {
		suitidentifier = "!"
	} else {
		suitidentifier = c.Suit.Name()[0:1]
	}
	return strings.ToUpper(faceidentifier + suitidentifier)
}

// String returns a human-readable representation of a Card.
func (c Card) String() string {
	var cardFace string
	if c.Face == 14 {
		cardFace = "! "
	} else {
		cardFace = c.Face.Name()[0:1]
		if c.Face == 10 {
			cardFace += "0"
		} else {
			cardFace += " "
		}
	}
	return fmt.Sprintf("[%s %s  %c]", cardFace, c.Suit.Symbol(), c.Suit.Name()[0])
}

// MarshalJSON marshals a Card.
func (c Card) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Identifier())
}

// UnmarshalJSON unmarshals a Card.
func (c *Card) UnmarshalJSON(b []byte) error {
	var identifier string
	if err := json.Unmarshal(b, &identifier); err != nil {
		return err
	}
	card, _ := Parse(identifier) // Ignore ok
	*c = card
	return nil
}
