/*
Package joker provides playing cards.
Cards
To initialize a card:
    card := joker.Card{FaceAce, SuitSpades}
To sort a set of cards:
    hand := joker.Cards{
      Card{FaceAce, SuitSpades},
      Card{Face3, SuitSpades},
      Card{Face2, SuitSpades}
    }
    sort.Sort(hand)
Deck
To initialize a joker, call NewDeck with a seed for the random number generator
used when shuffling. A seed value of zero will be replaced with the
current unix time in nanoseconds.
    joker := joker.NewDeck(StandardCards, 0)
*/

package joker
