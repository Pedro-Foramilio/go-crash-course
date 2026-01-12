package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")

	loadedCards := newDeckFromFile("my_cards.txt")

	hand, remainingCards := deal(cards, 5)
	fmt.Println("hand:")
	hand.print()
	fmt.Println("remaining cards:")
	remainingCards.print()

	fmt.Println("loaded cards from file:")
	loadedCards.print()
	loadedCards.shuffle()
	println("Shuffled cards:")
	loadedCards.print()
}
