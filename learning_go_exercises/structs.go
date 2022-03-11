package main

import "fmt"

// Data structures
func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println("Structs", poodle)

	// This prints the object
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	// Change values
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// If struct starts with uppercase letter it could be used in other parts of the application, exported.
// If struct starts with lowercase letter, then it's private=non-exported.

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
