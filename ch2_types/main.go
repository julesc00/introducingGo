package main

import "fmt"

var stars = "*****************"

func main() {
	addition()
	subtraction()
	multiplication()
	division()
}

func addition() {
	fmt.Println(stars, "Addition", stars)
	fmt.Println("2 + 2 = ", 2+2)
	fmt.Println("5 + 5 = ", 5.0+5.0)
}

func subtraction() {
	fmt.Println(stars, "Subtraction", stars)
	fmt.Println("10 - 5 = ", 10-5)
	fmt.Println("23 - 20 = ", 23-20)
}

func multiplication() {
	fmt.Println(stars, "Multiplication", stars)
	a, b, c, d := 20, 3, 32132, 42452
	fmt.Println("20 * 3 = ", a*b)
	fmt.Println("32,132 * 42,452 =", c*d)
}

func division() {
	fmt.Println(stars, "Division", stars)
	a, b := 450.0, 20.0
	fmt.Println("450 / 20 = ", a/b)
}
