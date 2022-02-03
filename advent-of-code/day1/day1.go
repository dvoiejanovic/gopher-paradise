package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	depths, _ := ioutil.ReadFile("input.txt")
	// depthIncrements := 0
	// for _, depth := range depths {
	// 	fmt.Println(depth)
	// }
	fmt.Println(strconv.Atoi(int(depths)))
}
