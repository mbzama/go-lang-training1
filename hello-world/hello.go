package main

import (
	"fmt"
)

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func main() {
	//c := color("Red")
	// fmt.Println(c.describe("is an awesome color"))

	fmt.Println([]byte("a b c"))
}
