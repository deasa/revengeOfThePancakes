package main

import (
	"fmt"
)

type customError struct {
	CustomMessage string
	OriginalError error
}

func (err customError) Error() string {
	return fmt.Sprintf("Custom Error:   %v\nOriginal Error: %v\n", err.CustomMessage, err.OriginalError)
}

func getNumStacksFromInput() (numStacks int, err error) {
	_, err = fmt.Scan(&numStacks)

	if err != nil || numStacks > 100 || numStacks < 0 {
		return 0, customError{fmt.Sprintf("the number of pancake stacks must be an integer from 1 to 100"), err}
	}
	return numStacks, nil
}

func getStackConfigFromInput() (config string, err error) {
	_, err = fmt.Scan(&config)

	if err != nil || !isValidStackConfiguration(config) {
		return "", customError{"each pancake stack configuration must only have the characters '+' and '-'", err}
	}
	return config, nil
}

func getPancakeStacksFromInput(numStacks int) (stacks []string, err error) {
	var pancakeStacks []string

	for i := 0; i < numStacks; i++ {
		stackConfig, readErr := getStackConfigFromInput()
		if readErr != nil {
			return make([]string, 0), readErr
		}
		pancakeStacks = append(pancakeStacks, stackConfig)
	}
	return pancakeStacks, nil
}
