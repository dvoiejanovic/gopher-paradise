package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}

}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colors)

}
