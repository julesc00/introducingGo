package main

import (
	"fmt"
	"sort"
)

// slices are of the same data type, but editable.
// It's just an array without a default number indexes.
func main() {
	// As an empty slice
	var slice1 []string
	// appending a new value
	slice1 = append(slice1, "hello")

	// With set values
	slice2 := []string{
		"Hello",
		"there!",
	}
	fmt.Println(slice2)
	colors := []string{"Red", "green", "white", "pink", "orange"}
	fmt.Println(colors)
	// Remove an item, in this case it's the first one.
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// Declare with initial size and the type; using
	// make(type of slice's items, initial length, optional capacity that limits the number of items )
	// If I exceed the number of appended items, it'll throw an error.
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 723
	numbers[2] = 233
	numbers[3] = 324
	numbers[4] = 421

	// with no item limitation, and appending several values at once
	numbers2 := make([]int, 2)
	numbers2[0] = 23
	numbers2[1] = 55
	numbers2 = append(numbers2, 100, 200, 105)
	fmt.Println(numbers2)

	fmt.Println(numbers)

	// Remove the last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// Sorting items
	sort.Ints(numbers2)
	sort.Strings(colors)
	fmt.Println(numbers2, colors)
}
