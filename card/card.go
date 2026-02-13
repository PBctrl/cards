package card

import (
	"cards/rank"
	"cards/suit"
)

// Card repräsentiert eine Spielkarte mit einem Wert und einer Farbe.
type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}
