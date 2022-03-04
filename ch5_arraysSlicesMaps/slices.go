package main

import "fmt"

func main() {
	// Create a slice
	slice1 := make([]float64, 5)

	// The make function allows a third parameter
	slice2 := make([]float64, 5, 10)

	// Another way to create slices is to use the [how : high] expression.
	slice3 := [5]float64{1, 2, 3, 4, 5}
	x := slice3[2:5]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(x)
}
