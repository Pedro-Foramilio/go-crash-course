package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	greeting := b.getGreeting()
	fmt.Println(greeting)
}

func (eb englishBot) getGreeting() string {
	// type specific logic
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	//type specific logic
	return "Hola!"

}
