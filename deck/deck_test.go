package deck

import (
	"cards/card"
	"cards/rank"
	"cards/suit"
	"fmt"
)

func ExampleEmpty() {
	d := Empty()

	fmt.Println(len(d.cards))

	// Output:
	// 0
}

func ExampleNew32() {
	d := New32()

	fmt.Println(len(d.cards))

	fmt.Println(d.cards[0])
	fmt.Println(d.cards[1])
	fmt.Println(d.cards[2])
	fmt.Println(d.cards[3])
	fmt.Println(d.cards[4])
	fmt.Println(d.cards[5])
	fmt.Println(d.cards[6])
	fmt.Println(d.cards[7])
	fmt.Println(d.cards[8])
	fmt.Println(d.cards[9])
	fmt.Println(d.cards[10])
	fmt.Println(d.cards[11])
	fmt.Println(d.cards[12])
	fmt.Println(d.cards[13])
	fmt.Println(d.cards[14])
	fmt.Println(d.cards[15])
	fmt.Println(d.cards[16])
	fmt.Println(d.cards[17])
	fmt.Println(d.cards[18])
	fmt.Println(d.cards[19])
	fmt.Println(d.cards[20])
	fmt.Println(d.cards[21])
	fmt.Println(d.cards[22])
	fmt.Println(d.cards[23])
	fmt.Println(d.cards[24])
	fmt.Println(d.cards[25])
	fmt.Println(d.cards[26])
	fmt.Println(d.cards[27])
	fmt.Println(d.cards[28])
	fmt.Println(d.cards[29])
	fmt.Println(d.cards[30])
	fmt.Println(d.cards[31])

	// Output:
	// 32
	// {7 ♣}
	// {8 ♣}
	// {9 ♣}
	// {10 ♣}
	// {J ♣}
	// {Q ♣}
	// {K ♣}
	// {A ♣}
	// {7 ♠}
	// {8 ♠}
	// {9 ♠}
	// {10 ♠}
	// {J ♠}
	// {Q ♠}
	// {K ♠}
	// {A ♠}
	// {7 ♥}
	// {8 ♥}
	// {9 ♥}
	// {10 ♥}
	// {J ♥}
	// {Q ♥}
	// {K ♥}
	// {A ♥}
	// {7 ♦}
	// {8 ♦}
	// {9 ♦}
	// {10 ♦}
	// {J ♦}
	// {Q ♦}
	// {K ♦}
	// {A ♦}
}

func ExampleNew52() {
	d := New52()

	fmt.Println(len(d.cards))

	fmt.Println(d.cards[0])
	fmt.Println(d.cards[1])
	fmt.Println(d.cards[2])
	fmt.Println(d.cards[3])
	fmt.Println(d.cards[4])
	fmt.Println(d.cards[5])
	fmt.Println(d.cards[6])
	fmt.Println(d.cards[7])
	fmt.Println(d.cards[8])
	fmt.Println(d.cards[9])
	fmt.Println(d.cards[10])
	fmt.Println(d.cards[11])
	fmt.Println(d.cards[12])
	fmt.Println(d.cards[13])
	fmt.Println(d.cards[14])
	fmt.Println(d.cards[15])
	fmt.Println(d.cards[16])
	fmt.Println(d.cards[17])
	fmt.Println(d.cards[18])
	fmt.Println(d.cards[19])
	fmt.Println(d.cards[20])
	fmt.Println(d.cards[21])
	fmt.Println(d.cards[22])
	fmt.Println(d.cards[23])
	fmt.Println(d.cards[24])
	fmt.Println(d.cards[25])
	fmt.Println(d.cards[26])
	fmt.Println(d.cards[27])
	fmt.Println(d.cards[28])
	fmt.Println(d.cards[29])
	fmt.Println(d.cards[30])
	fmt.Println(d.cards[31])
	fmt.Println(d.cards[32])
	fmt.Println(d.cards[33])
	fmt.Println(d.cards[34])
	fmt.Println(d.cards[35])
	fmt.Println(d.cards[36])
	fmt.Println(d.cards[37])
	fmt.Println(d.cards[38])
	fmt.Println(d.cards[39])
	fmt.Println(d.cards[40])
	fmt.Println(d.cards[41])
	fmt.Println(d.cards[42])
	fmt.Println(d.cards[43])
	fmt.Println(d.cards[44])
	fmt.Println(d.cards[45])
	fmt.Println(d.cards[46])
	fmt.Println(d.cards[47])
	fmt.Println(d.cards[48])
	fmt.Println(d.cards[49])
	fmt.Println(d.cards[50])
	fmt.Println(d.cards[51])

	// Output:
	// 52
	// {2 ♣}
	// {3 ♣}
	// {4 ♣}
	// {5 ♣}
	// {6 ♣}
	// {7 ♣}
	// {8 ♣}
	// {9 ♣}
	// {10 ♣}
	// {J ♣}
	// {Q ♣}
	// {K ♣}
	// {A ♣}
	// {2 ♠}
	// {3 ♠}
	// {4 ♠}
	// {5 ♠}
	// {6 ♠}
	// {7 ♠}
	// {8 ♠}
	// {9 ♠}
	// {10 ♠}
	// {J ♠}
	// {Q ♠}
	// {K ♠}
	// {A ♠}
	// {2 ♥}
	// {3 ♥}
	// {4 ♥}
	// {5 ♥}
	// {6 ♥}
	// {7 ♥}
	// {8 ♥}
	// {9 ♥}
	// {10 ♥}
	// {J ♥}
	// {Q ♥}
	// {K ♥}
	// {A ♥}
	// {2 ♦}
	// {3 ♦}
	// {4 ♦}
	// {5 ♦}
	// {6 ♦}
	// {7 ♦}
	// {8 ♦}
	// {9 ♦}
	// {10 ♦}
	// {J ♦}
	// {Q ♦}
	// {K ♦}
	// {A ♦}
}

func ExampleDeck_Add() {
	d := Empty()

	d.Add(card.Card{Rank: rank.Queen, Suit: suit.Spades})
	fmt.Println(len(d.cards))

	// Output:
	// 1
}
