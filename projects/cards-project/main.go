package main

import "fmt"

func main() {
	deck_path := "deck.json"
	// slice

	existingDeck := readDeckFromFile(deck_path)
	cards := deck{}

	if existingDeck.Size >= 10 {
		cards = existingDeck
	} else {
		cards = newDeck()
	}

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

	// aceCount := cards.faceCount(cardFaces[0])
	// spadeCount := cards.suitCount(cardSuits[0])

	// fmt.Println("Aces in deck:", aceCount)
	// fmt.Println("Spades in deck:", spadeCount)
	fmt.Println()

	// write remaining deck to file
	cards.saveDeckToFile(deck_path)

	// read deck from file
	newDeck := readDeckFromFile(deck_path)

	fmt.Println("Remaining Cards:")
	newDeck.print()
}
