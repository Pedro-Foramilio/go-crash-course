package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	cards := []string{"Ace of Spades", newCard(), card}

	//does not modify original slice, creates a new one
	cards = append(cards, "Six of Spades")

	// println(cards) prints memory address

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
