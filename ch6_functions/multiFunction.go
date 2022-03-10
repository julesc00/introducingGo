package main

import "fmt"

func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	f(x)

	y, z := f3()
}

// Return types can have names.
func f2() (r int) {
	r = 1
	return
}

// Multiple values can be returned
func f3() (int, int) {
	return 5, 6
}
