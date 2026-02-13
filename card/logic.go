package card

// MatchesRank prüft, ob die Karte den gleichen Wert wie eine andere Karte hat.
func (c Card) MatchesRank(other Card) bool {
	return c.Rank == other.Rank
}

// MatchesSuit prüft, ob die Karte die gleiche Farbe wie eine andere Karte hat.
func (c Card) MatchesSuit(other Card) bool {
	return c.Suit == other.Suit
}
