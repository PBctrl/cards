package hand

import (
	"cards/card"
	"cards/rank"
	"cards/suit"
	"fmt"
)

func ExampleHand_String() {
	h := New()
	h.AddCard(card.Card{Rank: rank.Ace, Suit: suit.Spades})
	h.AddCard(card.Card{Rank: rank.Ten, Suit: suit.Hearts})

	fmt.Println(h.String())

	// Output:
	// ┌─────┐┌─────┐
	// │A    ││10   │
	// │  ♠  ││  ♥  │
	// │    A││   10│
	// └─────┘└─────┘
}
