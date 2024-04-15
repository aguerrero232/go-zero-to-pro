package main

import (
	"testing"
)

// Using go test runner.
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d.Cards) != d.Size {
		t.Errorf("Expected deck size of %v, but got %v", d.Size, len(d.Cards))
	}
}

func TestFaceCount(t *testing.T) {
	d := newDeck()
	for _, face := range cardFaces {
		c := d.faceCount(face)
		if c != 4 {
			t.Errorf("Expected 4 cards with face %v, but got %v", face, c)
		}
	}		
}

func TestSuitCount(t *testing.T) {
	d := newDeck()
	for _, suit := range cardSuits {
		c := d.suitCount(suit)
		if c != 13 {
			t.Errorf("Expected 13 cards with suit %v, but got %v", suit, c)
		}
	}
}