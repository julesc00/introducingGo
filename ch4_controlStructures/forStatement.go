package main

import "fmt"

func main() {
	exampleOne()
	exampleTwo()
	switchStatement()
}

func exampleOne() {
	// A range generator
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
}

func exampleTwo() {
	// Another range generator
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func switchStatement() {
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println(i, "Zero")
		case 1:
			fmt.Println(i, "One")
		case 2:
			fmt.Println(i, "Two")
		case 3:
			fmt.Println(i, "Three")
		case 4:
			fmt.Println(i, "Four")
		case 5:
			fmt.Println(i, "Five")
		case 6:
			fmt.Println(i, "Six")
		case 7:
			fmt.Println(i, "Seven")
		case 8:
			fmt.Println(i, "Eight")
		case 9:
			fmt.Println(i, "Nine")
		case 10:
			fmt.Println(i, "Ten")
		default:
			fmt.Println("Unknown number")
		}
	}
}
