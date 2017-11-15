package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards, _ = deal(cards, 5)
	cards.saveToFile("tmp/my_cards")

	fmt.Println(newDeckFromFile("tmp/my_cards"))
}
