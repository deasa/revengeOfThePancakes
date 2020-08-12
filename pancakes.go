package main

import (
	"strings"
)

func isValidStackConfiguration(stackConfig string) bool {
	return len(stackConfig) > 0 && len(strings.Trim(stackConfig, "+-")) == 0
}

func stackHasBlankSideOnBottom(stackConfig string) bool {
	return stackConfig[len(stackConfig)-1] == '-'
}

func getCountOfSignChanges(stackConfig string) int {

	if !isValidStackConfiguration(stackConfig) {
		return 0
	}

	var (
		numChanges int    = 0
		runes      []rune = []rune(stackConfig)
		prevSign   string = string(runes[0])
	)

	for i := 0; i < len(runes); i++ {
		if string(runes[i]) != prevSign {
			numChanges++
			prevSign = string(runes[i])
		}
	}

	return numChanges
}

func calculateMinimumFlipsNeeded(stackConfig string) int {

	if !isValidStackConfiguration(stackConfig) {
		return 0
	}

	numSignChanges := getCountOfSignChanges(stackConfig)

	if stackHasBlankSideOnBottom(stackConfig) {
		numSignChanges++
	}

	return numSignChanges
}
