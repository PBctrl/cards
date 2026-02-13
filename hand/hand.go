package hand

import (
	"cards/card"
	"strings"
)

type Hand struct {
	cards []card.Card
}

func New() *Hand {
	return &Hand{}
}

func (h *Hand) AddCard(c card.Card) {
	h.cards = append(h.cards, c)
}

func (h Hand) String() string {
	lines := make([]string, 5)
	for _, c := range h.cards {
		c_lines := c.FrontLines()
		for i, line := range c_lines {
			lines[i] += line
		}
	}
	return strings.Join(lines, "\n")
}
