package main

import (
	"fmt"
)

/*
block comments are like so.
*/

const favColor string = "blue" //const means error if try and re-assign

func main() {
	var guess string //don't need semicolons

	for { // for {} just means infinite loop over the inner curly braces
		fmt.Println("Guess my favorite color.")
		if _, err := fmt.Scanln(&guess); err != nil { //Scanln actually returns the error_number, error message
			fmt.Printf("%s\n", err) //if the error is not nil, print the error and end
			return
		}
		if favColor == guess { //string equality with ==
			fmt.Printf("Correct! %q is my favorite color!\n", favColor) //%q prints the string surrounded by "quotation marks"
			return
		}
		fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guess)
	}
}
