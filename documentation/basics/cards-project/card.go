package main

import "strconv"

// card is a custom type that represents a playing card in a deck
type card struct {
	face       string
	face_value int
	suit       string
}

func (c card) print() string {
	return c.face + " of " + c.suit + " (" + strconv.Itoa(c.face_value) + ")"
}

func (c card) getCardFace() string {
	return c.face
}

func (c card) getCardSuit() string {
	return c.suit
}

func (c card) getCardValue() int {
	return c.face_value
}
