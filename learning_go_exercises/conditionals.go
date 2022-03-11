package main

import "fmt"

func main() {
	answer := 42
	var result string

	if answer < 0 {
		result = "less than zero"
	} else if answer == 0 {
		result = "equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	// Same as previous, but inline initialized variable x
	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}
