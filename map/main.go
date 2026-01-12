package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	delete(colors, "green")

	// var map2 map[string]string

	map3 := make(map[string]string)
	map3["white"] = "#FFFFFF"
	map3["yellow"] = "#FFFF00"

	fmt.Println(colors)
	fmt.Println(map3["white"])
	fmt.Println(map3["yellow"])

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color:", color, "Hex:", hex)
	}
}
