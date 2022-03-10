package main

import "fmt"

func main() {
	// a pointer, which points to another variable
	// In Go, if a pointer is only initialized and not provided a value, it'll default to nil
	anInt := 42
	var p = &anInt // & means I'm pointing at the address of the variable it follows.
	fmt.Println("Value of p address:", p)
	fmt.Println("Value of p value: ", *p)

	// Another way of pointing
	val1 := 50.23
	pointer1 := &val1
	fmt.Println("Value of pointer1 address:", pointer1)
	fmt.Println("Value of pointer1:", *pointer1)

	// If I change the value of the pointer, it changes the value it's pointing at.
	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value of val1:", val1)
}
