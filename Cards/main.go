package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	fmt.Println("hand:")
	hand.print()
	fmt.Println("remaining cards:")
	remainingCards.print()

}
