package deck

import (
	"cards/card"
	"cards/rank"
	"cards/suit"
	"math/rand"
)

// Repräsentiert einen Kartenstapel.
type Deck struct {
	cards []card.Card
}

// Empty erzeugt neues leeres Deck
func Empty() *Deck {
	return &Deck{}
}

// New32 erzeugt ein neues Skatblatt mit 32 Karten.
func New32() *Deck {

	result := &Deck{}
	for s := suit.Clubs; s <= suit.Diamonds; s++ {
		for r := rank.Seven; r <= rank.Ace; r++ {
			c := card.Card{Rank: r, Suit: s}
			result.cards = append(result.cards, c)
		}
	}

	return result
}

// New52 erzeugt ein neues französisches Blatt mit 52 Karten.
func New52() *Deck {

	result := &Deck{}
	for s := suit.Clubs; s <= suit.Diamonds; s++ {
		for r := rank.Two; r <= rank.Ace; r++ {
			c := card.Card{Rank: r, Suit: s}
			result.cards = append(result.cards, c)
		}
	}

	return result

}

// Add fügt eine Karte zum Deck hinzu.
func (d *Deck) Add(c card.Card) {
	d.cards = append(d.cards, c)
}

// Shuffle mischt das Deck.
func (d *Deck) Shuffle() {

	// Definiere eine Funktion, die zwei Karten im Deck vertauscht.
	swap := func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}

	// Verwende die eben definierte Swap-Funktion,
	// um die Karten im Deck zu mischen.
	rand.Shuffle(len(d.cards), swap)
}

// Draw zieht die oberste Karte.
// Entfernt die Karte aus dem Deck und liefert sie zurück.
func (d *Deck) Draw() card.Card {
	c := d.Top()
	d.cards = d.cards[:len(d.cards)-1]
	return c
}

// Top liefert die oberste Karte.
// Entfernt sie aber nicht.
func (d Deck) Top() card.Card {
	return d.cards[len(d.cards)-1]
}
