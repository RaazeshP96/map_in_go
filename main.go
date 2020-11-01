package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#001200",
		"blue":  "#562255",
		"white": "#584652",
	}
	// colors["red"] = "#000000"
	// fmt.Println(colors["red"])
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
