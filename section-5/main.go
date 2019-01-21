package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"lime":  "#00ff00",
		"white": "#ffffff",
	}

	// colors[10] = "#fffff"
	// delete(colors, 10)
	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
