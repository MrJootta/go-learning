package main

import "fmt"


type bot interface {
	getGreeting()  string
}
type englishBot struct {}
type portugueseBot struct {}

func main() {
	eb := englishBot{}
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//With receiver that we don't use
func (eb englishBot) getGreeting() string {
	return "Hello!"
}

//Just specification of receiver works the same
func (portugueseBot) getGreeting() string {
	return "Olá!"
}
