package main

import "testing"

func TestNewCard(t *testing.T) {
	c := newCard("♥", "K")
	expectedSuit := "♥"
	expectedValue := "K"
	if c.suit != expectedSuit {
		t.Errorf("Suit did not match, expected %v got %v", expectedSuit, c.suit)
	}
	if c.value != expectedValue {
		t.Errorf("Card value did not match, expected %v got %v", expectedValue, c.value)
	}
}
