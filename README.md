# revengeOfThePancakes

Hello! 

This is my solution written in Go to a problem that came from a Google Code Jam project that was made available online in 2016.

## To Use:
* Install Go <i>needs Go 1.15 or later</i>
* Clone this repo
* While inside the new repo's directory, run either `go build` or `go install`
* To run, you can directly run the executable and manually enter the required inputs, or you can redirect the stdin to a file or similar that contains the required inputs.
>* As an example, you can redirect the stdin to the `input.txt` file found in the root directory. In bash, run the command `./revengeofthepancakes.exe < input.txt` to do this.
* To run the tests in `main_test.go` simply run the command `go test`

There are five main files - `main.go`, `readFromInput.go`, `pancakes.go`, `main_test.go`, and `input.txt`

### `main.go`:
This the entrypoint of the application. It accepts input from stdin in the following format:
* `{{number of test cases}}` i.e. 3
* `{{first stack configuration}}` i.e. +--++--
* `{{second stack configuration}}` i.e. ---++-
* ...
    
>NOTES & GOTCHAS:
>* The number of test cases must be first and must be an integer between [1, 100] inclusive. 
>* The pancake stack configuration must be comprised of only the characters '+' and '-' and must have length between [1, 100] inclusive.
>* An error will be thrown and the program will exit if all conditions are not met.

### `readFromInput.go`:
  This file contains the functions used to read in data from stdin, as well as some error checking.
  
### `pancakes.go`:
  This file has the business logic for figuring out how many flips are needed.
  One key realization I had was that the minimum number of flips was equal to the number of times the orientation of the pancakes switched throughout the stack.
  However, if the stack had a sunny-side down pancake ('-') on the very bottom, exactly one more flip was needed.

### `main_test.go`:
  This was a little extra credit for me, as I wanted to play around with Go's testing module.
  The tests throw valid and invalid stack configurations into pancakes.go's critical functions and the output is verified against known expected values.
