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

