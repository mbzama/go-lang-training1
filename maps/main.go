package main

import (
	"fmt"
)

func main() {
	//Creating syntax 3
	//colors := make(map[string]string)

	//Creating syntax 2
	//var colors map[string]string

	//Creating syntax 1

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#DD0001",
		"white": "#FFFFFF",
	}

	printMap(colors)
	fmt.Println("---------")

	//delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
	c["orange"] = "#AA0002"
}
