package card

import (
	"fmt"
	"strings"
)

// Front gibt die Karte als String zurück.
// Dabei soll die Karte in AsciiArt mit
// einem Rahmen dargestellt werden, z.B.:
// ┌─────┐
// │A    │
// │  ♠  │
// │    A│
// └─────┘
func (c Card) Front() string {
	rank := c.Rank
	suit := c.Suit

	lines := []string{
		"┌─────┐",
		fmt.Sprintf("│%-2s   │", rank),
		fmt.Sprintf("│  %s  │", suit),
		fmt.Sprintf("│   %2s│", rank),
		"└─────┘",
	}
	return strings.Join(lines, "\n")
}

// Back gibt die Rückseite der Karte als String zurück.
func (c Card) Back() string {
	lines := []string{
		"┌─────┐",
		"│\\   /│",
		"│ | | │",
		"│/   \\│",
		"└─────┘",
	}
	return strings.Join(lines, "\n")
}
