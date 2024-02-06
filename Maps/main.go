package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "1",
		"green": "2",
	}
	colors["white"] = "3"

	for color, hex := range colors {
		fmt.Println(color, hex)
	}
}
