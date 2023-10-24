package main

import "fmt"

func main() {
	// slice
	cards := newDeck()
	cards.shuffle()

	handSize := 5
	players := 2

	fmt.Println()

	for i := 0; i < players; i++ {
		hand, remainingCards := cards.deal(handSize)
		fmt.Println("Player", i+1, "Hand:")
		hand.print()
		fmt.Println()
		cards = remainingCards
	}

	aceCount := cards.faceCount(cardFaces[0])
	spadeCount := cards.suitCount(cardSuits[0])

	fmt.Println("Aces in deck:", aceCount)
	fmt.Println("Spades in deck:", spadeCount)
	fmt.Println()

	// write remaining deck to file
	cards.saveToFile("documentation/basics/cards-project/my_cards.json")

	// read deck from file
	newDeck := readDeckFromFile("documentation/basics/cards-project/my_cards.json")

	fmt.Println("Deck from file:")
	newDeck.print()
}
