package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#61970",
		"white":  "gh982y2",
		"orange": "n14789",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
