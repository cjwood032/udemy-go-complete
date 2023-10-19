package main

import "testing"

func TestNewDeckSize(t *testing.T) {
	d := newDeck()
	expected := 52
	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but got %v", expected, len(d))
	}
}
