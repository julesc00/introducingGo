package main

import (
	"fmt"
)

// It's better to use slices instead of arrays, slices are editable.
// An array in Go is an ordered collection of values with a fixed number of indexes.
// An array is = tuple in Python
func main() {
	// array example
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "white"
	fmt.Println("colors", colors)
	fmt.Println(colors[0])

	var numArr = [5]int{5, 1, 2, 3, 4}
	fmt.Println(numArr)
	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:")
}
