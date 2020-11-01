package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#001200",
		"blue": "#562255",
	}
	colors["red"] = "#000000"
	fmt.Println(colors["red"])
}
