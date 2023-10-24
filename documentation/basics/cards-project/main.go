package main

import "fmt"

func main() {
	// slice
	cards := newDeck()
	cards.shuffle()

	handSize := 5
	players := 2

	for i := 0; i < players; i++ {
		hand, remainingCards := cards.deal(handSize)
		fmt.Println("Player", i+1, "Hand:")
		hand.print()
		fmt.Println()
		cards = remainingCards
	}

	fmt.Println("Remaining Cards:")
	cards.print()
	fmt.Println()
}
