package main

import "testing"

func TestNewDeckSize(t *testing.T) {
	d := newDeck()
	expected := 52
	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, len(d))
	}
}
func TestNewDeckCards(t *testing.T) {
	d := newDeck()
	expectedStart := "A♠"
	if d[0] != expectedStart {
		t.Errorf("Expected first card of deck %v, but got %v", expectedStart, d[0])
	}
	expectedEnd := "K♥"
	if d[len(d)-1] != expectedEnd {
		t.Errorf("Expected last card of deck %v, but got %v", expectedEnd, d[len(d)-1])
	}
}
func TestShuffledDeckCards(t *testing.T) {
	d := newDeck()
	s := newDeck()
	sP := &s
	sP.shuffle()
	if len(d) != len(s) {
		t.Errorf("Expected decks to be equal length, but got %v and %v", len(s), len(d))
	}
	if d[0] == s[0] && d[1] == s[1] && d[2] == s[2] {
		//it's a 1/140608 chance to fail because all 3 cards ended up in the same spot
		t.Errorf("Expected shuffled deck to be in a different order than a non-shuffled deck")
	}
}
