package joker

import (
	"math/rand"
	"time"
)

// Deck defines a playing card deck containing any number of cards.
type Deck struct {
	Cards Cards
	rand  *rand.Rand
}

// NewDeck initializes a deck of cards. A seed value of 0 is replaced with the
// current unix time in nanoseconds.
func NewDeck(c Cards, seed int64) *Deck {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &Deck{Cards: c.Copy(), rand: rand.New(rand.NewSource(seed))}
}

// Shuffle randomizes the order of the deck using the Fisher-Yates shuffle
// algorithm.
func (d *Deck) Shuffle() {
	for i := len(d.Cards) - 1; i > 0; i-- {
		j := d.rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Draw removes cards from the deck and returns them as a slice.
func (d *Deck) Draw(count int) (cards Cards, ok bool) {
	if count > len(d.Cards) {
		return nil, false
	}
	cards = d.Cards[0:count].Copy()
	d.Cards = d.Cards[count:]
	return cards, true
}
