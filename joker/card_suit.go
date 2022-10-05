package joker

// CardSuit defines a card suit.
type CardSuit int

// Card suits
const (
	SuitHearts   CardSuit = 1
	SuitDiamonds CardSuit = 2
	SuitClubs    CardSuit = 3
	SuitSpades   CardSuit = 4
	SuitJoker    CardSuit = 5
)

// StandardSuits is a slice of all card suits except Jokers.
var StandardSuits = []CardSuit{
	SuitHearts,
	SuitDiamonds,
	SuitClubs,
	SuitSpades,
}

// AllSuits is a slice of all suits.
var AllSuits = []CardSuit{
	SuitHearts,
	SuitDiamonds,
	SuitClubs,
	SuitSpades,
	SuitJoker,
}

// Card symbols
const (
	SymbolHearts   = "♥"
	SymbolDiamonds = "♦"
	SymbolClubs    = "♣"
	SymbolSpades   = "♠"
	SymbolJoker    = "!"
)

// Symbol returns the card suit symbol.
func (s CardSuit) Symbol() string {
	switch s {
	case SuitHearts:
		return SymbolHearts
	case SuitDiamonds:
		return SymbolDiamonds
	case SuitClubs:
		return SymbolClubs
	case SuitSpades:
		return SymbolSpades
	case SuitJoker:
		return SymbolJoker
	default:
		return "Invalid"
	}
}

// Name returns the card suit name.
func (s CardSuit) Name() string {
	switch s {
	case SuitHearts:
		return "Hearts"
	case SuitDiamonds:
		return "Diamonds"
	case SuitClubs:
		return "Clubs"
	case SuitSpades:
		return "Spades"
	case SuitJoker:
		return "Joker"
	default:
		return "Invalid"
	}
}

// String returns the card suit name.
func (s CardSuit) String() string {
	return s.Name()
}
