package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a new type of 'deck'
// which is a slice of cards
type deck []card

// return a standard deck of 52 cards
func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardFaces := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	newDeck := deck{}
	for _, suit := range cardSuits {
		for j, face := range cardFaces {
			newDeck = append(newDeck, card{face: face, face_value: j + 1, suit: suit})
		}
	}
	return newDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card.print())
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) removeCard(index int) deck {
	return append(d[:index], d[index+1:]...)
}

func (d deck) addCard(c card) deck {
	return append(d, c)
}

func (d deck) shuffle() deck {

	if len(d) < 2 {
		return d
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
