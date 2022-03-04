package main

import "fmt"

/*
	Write a program that prints the numbers from 1 to 100, but for multiples of three, print “Fizz” instead of
	the number, and for the multiples of five, print “Buzz.” For numbers that are multiples of both three and
	five, print “FizzBuzz.”

	I'll come back to it some other time.
*/
func main() {
	for i := 1; i <= 100; i++ {
		switch i {
		case i + 3:
			fmt.Println("fizz")
		case i + 5:
			fmt.Println("buzz")

		}

	}
}
