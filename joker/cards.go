package joker

import (
	"sort"
	"strings"
)

// Cards is a slice of Cards.
type Cards []Card

// StandardCards is a slice of standard cards.
var StandardCards = Cards{
	Card{FaceAce, SuitHearts},
	Card{Face2, SuitHearts},
	Card{Face3, SuitHearts},
	Card{Face4, SuitHearts},
	Card{Face5, SuitHearts},
	Card{Face6, SuitHearts},
	Card{Face7, SuitHearts},
	Card{Face8, SuitHearts},
	Card{Face9, SuitHearts},
	Card{Face10, SuitHearts},
	Card{FaceJack, SuitHearts},
	Card{FaceQueen, SuitHearts},
	Card{FaceKing, SuitHearts},
	Card{FaceAce, SuitDiamonds},
	Card{Face2, SuitDiamonds},
	Card{Face3, SuitDiamonds},
	Card{Face4, SuitDiamonds},
	Card{Face5, SuitDiamonds},
	Card{Face6, SuitDiamonds},
	Card{Face7, SuitDiamonds},
	Card{Face8, SuitDiamonds},
	Card{Face9, SuitDiamonds},
	Card{Face10, SuitDiamonds},
	Card{FaceJack, SuitDiamonds},
	Card{FaceQueen, SuitDiamonds},
	Card{FaceKing, SuitDiamonds},
	Card{FaceAce, SuitClubs},
	Card{Face2, SuitClubs},
	Card{Face3, SuitClubs},
	Card{Face4, SuitClubs},
	Card{Face5, SuitClubs},
	Card{Face6, SuitClubs},
	Card{Face7, SuitClubs},
	Card{Face8, SuitClubs},
	Card{Face9, SuitClubs},
	Card{Face10, SuitClubs},
	Card{FaceJack, SuitClubs},
	Card{FaceQueen, SuitClubs},
	Card{FaceKing, SuitClubs},
	Card{FaceAce, SuitSpades},
	Card{Face2, SuitSpades},
	Card{Face3, SuitSpades},
	Card{Face4, SuitSpades},
	Card{Face5, SuitSpades},
	Card{Face6, SuitSpades},
	Card{Face7, SuitSpades},
	Card{Face8, SuitSpades},
	Card{Face9, SuitSpades},
	Card{Face10, SuitSpades},
	Card{FaceJack, SuitSpades},
	Card{FaceQueen, SuitSpades},
	Card{FaceKing, SuitSpades},
}

func (c Cards) Len() int {
	return len(c)
}
func (c Cards) Less(i, j int) bool {
	return c[i].Value() < c[j].Value()
}
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Cards) String() string {
	var s strings.Builder
	for i := range c {
		if i > 0 {
			s.WriteRune(',')
		}
		s.WriteString(c[i].String())
	}
	return s.String()
}

// Copy returns a copy of the supplied cards.
func (c Cards) Copy() Cards {
	cc := make(Cards, len(c))
	copy(cc, c)
	return cc
}

// Remove returns the supplied cards with the specified card removed one time.
func (c Cards) Remove(card Card) Cards {
	cc := c.Copy()
	for i, searchCard := range c {
		if searchCard == card {
			return append(cc[:i], cc[i+1:]...)
		}
	}
	return cc
}

// RemoveIndex returns the supplied cards excluding the card at the specified index.
func (c Cards) RemoveIndex(i int) Cards {
	cc := c.Copy()
	return append(cc[:i], cc[i+1:]...)
}

// Contains returns whether the supplied cards contain the specified card.
func (c Cards) Contains(card Card) bool {
	for _, compcard := range c {
		if compcard.Equal(card) {
			return true
		}
	}
	return false
}

// Count returns the number of occurrences of the specified card.
func (c Cards) Count(card Card) int {
	var n int
	for _, compcard := range c {
		if compcard.Equal(card) {
			n++
		}
	}
	return n
}

// Equal returns whether the supplied cards are equal to another set of cards.
func (c Cards) Equal(cards Cards) bool {
	if len(c) != len(cards) {
		return false
	}
	for _, compcard := range c {
		if c.Count(compcard) != cards.Count(compcard) {
			return false
		}
	}
	return true
}

// Sorted returns the supplied cards in order.
func (c Cards) Sorted() Cards {
	cc := c.Copy()
	sort.Sort(cc)
	return cc
}

// Reversed returns the supplied cards in reverse order.
func (c Cards) Reversed() Cards {
	l := len(c)
	cc := make(Cards, l)
	for i := 0; i < l; i++ {
		cc[i] = c[l-i-1]
	}
	return cc
}

// Permutations returns all permutations of the supplied cards.
func (c Cards) Permutations() []Cards {
	var permute func(Cards, int)
	var res []Cards
	permute = func(c Cards, n int) {
		if n == 1 {
			res = append(res, c.Copy())
		} else {
			for i := 0; i < n; i++ {
				permute(c, n-1)
				if n%2 == 1 {
					tmp := c[i]
					c[i] = c[n-1]
					c[n-1] = tmp
				} else {
					tmp := c[0]
					c[0] = c[n-1]
					c[n-1] = tmp
				}
			}
		}
	}
	permute(c, len(c))
	return res
}

// Low returns the lowest valued card.
func (c Cards) Low() Card {
	if len(c) == 0 {
		return Card{}
	}
	l := c[0]
	for _, comp := range c[1:] {
		if comp.Value() < l.Value() {
			l = comp
		}
	}
	return l
}

// High returns the highest valued card.
func (c Cards) High() Card {
	if len(c) == 0 {
		return Card{}
	}
	h := c[0]
	for _, comp := range c[1:] {
		if comp.Value() > h.Value() {
			h = comp
		}
	}
	return h
}
