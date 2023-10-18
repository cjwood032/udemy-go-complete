package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, remaining := cards.deal(5)
	hand.print()
	remaining.print()

}
