package main

import "fmt"

func stackHasBlankSideOnBottom(stackConfig string) bool {
	return stackConfig[len(stackConfig)] == '-'
}

func getCountOfSignChanges(stackConfig string) int {

	prevSign := string(stackConfig[0])
	numChanges := 0
	for sign := range stackConfig {
		if string(sign) != prevSign {
			numChanges += 1
			prevSign = string(sign)
		}
	}

	return numChanges
}

func printCaseAndNumFlips(caseNum int, numFlips int) {
	fmt.Printf("Case # %v: %v\n", caseNum, numFlips)
}

func calculateMinimumFlipsNeeded(stackConfig string) int {
	numSignChanges := getCountOfSignChanges(stackConfig)

	if stackHasBlankSideOnBottom(stackConfig) {
		numSignChanges += 1
	}

	return numSignChanges
}
