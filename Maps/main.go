package main

import "fmt"

//Maps != Structs

func main() {
	colors := map[string]int {
		"red": 1,
		"blue": 2,
	}

	colors["white"] = 3

	printMap(colors)
}

func printMap(c map[string]int) {
	for _, num := range c {
		fmt.Println(num)
	}
}