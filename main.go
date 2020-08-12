package main

import (
	"fmt"
)

type CustomError struct {
	CustomMessage string
	OriginalError error
}

func (err CustomError) Error() string {
	return fmt.Sprintf("Custom Error:   %v\nOriginal Error: %v\n", err.CustomMessage, err.OriginalError)
}

func getNumCases() (numCases int, err error) {
	_, err = fmt.Scan(&numCases)

	if err != nil {
		return 0, CustomError{"couldn't read the number of cases", err}
	} else if numCases > 100 || numCases < 0 {
		return 0, CustomError{fmt.Sprintf("the number of cases '%v' is not within [0, 100] inclusive", numCases), err}
	}
	return numCases, nil
}

func printErrorIfExists(err error) (wasError bool) {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func main() {
	numCases, readErr := getNumCases()
	if printErrorIfExists(readErr) {
		return
	}

	fmt.Println(numCases)
}
