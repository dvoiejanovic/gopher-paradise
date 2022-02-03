package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	depths := strings.Fields(string(content))
	depthIncrements := 0

	for i, _ := range depths {
		if i == 0 {
			continue
		}

		previousDepthInt, _ := strconv.Atoi(depths[i-1])
		depthInt, err := strconv.Atoi(depths[i])

		if err != nil {
			log.Fatalf("Can't convert to int")
			break
		}

		if depthInt > previousDepthInt {
			depthIncrements++
		}
	}
	fmt.Println(depthIncrements)
}
