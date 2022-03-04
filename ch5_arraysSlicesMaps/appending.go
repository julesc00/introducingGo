package main

import "fmt"

func main() {
	// Append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	var slice3 []int // for an empty slice
	fmt.Println(slice3)

	addStringToSlice()
}

// Append full name to slice
func addStringToSlice() {
	var strSlice1 []string
	name := "Julio"
	lastName := "Briones"
	strSlice1 = append(strSlice1, name, lastName)
	fmt.Println(strSlice1[0], strSlice1[1])
}
