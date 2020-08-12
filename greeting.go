package main //declares that "code belongs to the main package"?

import ( //import other packages
	"fmt"     //fmt package comes with standard library, handles printing
	"strings" //so does strings, handles string manipulation
)

func main() {
	fmt.Println("Please enter your name.") //print with newline at end
	var name string                        //variable variable_name variable_type
	fmt.Scanln(&name)                      //wait for newline character
	name = strings.TrimSpace(name)         //trim whitespace before and after variable, like Python's .strip()
	fmt.Printf("Hi, %s! I'm Go!\n", name)  //if printing variables AND raw strings together need to use Printf, %s for string
}
