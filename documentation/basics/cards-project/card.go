package main

import (
	"strconv"
)

var cardFaces = [13]string{
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

var cardSuits = [4]string{
	"Spades",
	"Diamonds",
	"Hearts",
	"Clubs",
}

// card is a custom type that represents a playing card in a deck
type card struct {
	Face      string
	FaceValue int
	Suit      string
}

func (c card) print() string {
	return c.getCardFace() + " of " + c.getCardSuit() + " (" + strconv.Itoa(c.getCardValue()) + ")"
}

func (c card) getCardFace() string {
	return c.Face
}

func (c card) getCardSuit() string {
	return c.Suit
}

func (c card) getCardValue() int {
	return c.FaceValue
}
