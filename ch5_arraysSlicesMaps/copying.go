package main

import "fmt"

/*
copy takes two arguments: dst and src. All of the entries in src are copied into dst overwriting whatever is there.
If the lengths of the two slices are not the same, the smaller of the two will be used.
*/
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)

	// because slice2 has room for only two elements,
	//only the first two elements of slice1 are copied.
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
