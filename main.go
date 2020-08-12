package main

import "fmt"

func main() {
	numStacks, readErr := getNumStacksFromInput()
	if readErr != nil {
		fmt.Println(readErr)
		return
	} else if numStacks < 1 || numStacks > 100 {
		fmt.Println("The number of stacks must be in the range [1, 100] inclusive.")
	}

	pancakeStacks, readErr := getPancakeStacksFromInput(numStacks)
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	i := 1
	for _, stackConfig := range pancakeStacks {
		numFlips := calculateMinimumFlipsNeeded(stackConfig)
		fmt.Printf("Case # %v: %v\n", i, numFlips)
		i++
	}
}
