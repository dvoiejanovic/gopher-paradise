package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can't read content of %v", err)
	}

	instructions := strings.Split(string(content), " ")
	fmt.Println(len(instructions))
	for i := 1; i < len(instructions); i += 2 {
		direction := instructions[i-1]
		steps := instructions[i]
		fmt.Println(direction, steps)
	}
}
