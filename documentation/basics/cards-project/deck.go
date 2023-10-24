package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Create a new type of 'deck'
// which is a slice of cards
type deck struct {
	Cards []card
	Size  int
}

// return a standard deck of 52 cards
func newDeck() deck {
	newDeck := deck{}
	for _, suit := range cardSuits {
		for j, face := range cardFaces {
			newDeck.Cards = append(
				newDeck.Cards,
				card{
					Face:      face,
					FaceValue: j + 1,
					Suit:      suit,
				},
			)
		}
	}
	return newDeck
}

// write the json formatted deck to a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toJson(), 0666)
}

func readDeckFromFile(filename string) deck {
	jsonData, _ := os.ReadFile(filename)
	var d deck
	json.Unmarshal(jsonData, &d)
	return d
}

func (d deck) print() {
	for i, card := range d.Cards {
		fmt.Println(i+1, card.print())
	}
	fmt.Println("Total Cards:", d.Size)
}

func (d deck) toJson() []byte {
	jsonData, _ := json.Marshal(d)
	return jsonData
}

// deal a hand of cards from the deck and return the hand and the remaining cards
func (d deck) deal(handSize int) (deck, deck) {

	dealtDeck := deck{
		Cards: d.Cards[:handSize],
		Size:  handSize,
	}

	remainingDeck := deck{
		Cards: d.Cards[handSize:],
		Size:  len(d.Cards) - handSize,
	}

	return dealtDeck, remainingDeck
}

func (d deck) shuffle() deck {

	if len(d.Cards) < 2 {
		return d
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards))
		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
	return d
}

func (d deck) faceCount(face string) int {
	count := 0
	for _, card := range d.Cards {
		if card.Face == face {
			count++
		}
	}
	return count
}

func (d deck) suitCount(suit string) int {
	count := 0
	for _, card := range d.Cards {
		if card.Suit == suit {
			count++
		}
	}
	return count
}
