package main

type card struct {
	suit  string
	value string
}

func newCard(suit string, value string) card {
	c := card{suit: suit, value: value}
	return c
}
