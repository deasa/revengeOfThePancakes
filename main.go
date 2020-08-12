package main

import "fmt"

func main() {
	numStacks, readErr := getNumStacksFromInput()
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	pancakeStacks, readErr := getPancakeStacksFromInput(numStacks)
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	i := 1
	for _, stackConfig := range pancakeStacks {
		numFlips := calculateMinimumFlipsNeeded(stackConfig)
		printCaseAndNumFlips(i, numFlips)
		i++
	}
}
