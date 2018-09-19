package main

import (
	"fmt"
)

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

func (eb englishBot) getGreeting() string {
	//custom logic for english
	return "hi there!"
}

func (sb spanishBot) getGreeting() string {
	//custom logic for Spanish
	return "hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
