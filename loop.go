package main

import (
	"fmt"
)

func main() {

	//sharks := x is equal to var sharks = x
	//:= will implicitly figure out the type required
	sharks := []string{"hammerhead", "great white", "dogfish"}

	for i, shark := range sharks { //defaults to Python's enumerate
		fmt.Println(i, shark) //can use Println for variables if we only print variables and no raw strings
	}
}
