package main

import "fmt"

/*
An array is a numbered sequence of elements of a single type with a fixed length.
*/
func main() {
	arrOne()
	arrTwo()
}

func arrOne() {
	var x [5]int // Starts counting from 1 to 5
	x[4] = 100
	x[2] = 20
	fmt.Println(x)
}

func arrTwo() {
	var arr [5]float64
	arr[0] = 98
	arr[1] = 93
	arr[2] = 77
	arr[3] = 82
	arr[4] = 83

	// var total float64 = 0
	// or
	total := float64(0)
	for i := 0; i < 5; i++ {
		total += arr[i]
	}
	fmt.Println(arr)
	// I have to change the len() int type since the array is a float64 type array.
	fmt.Println("Length:", float64(len(arr)))
	fmt.Println(total)
	totalDivided := total / float64(len(arr))
	fmt.Println(totalDivided)
	fmt.Println("86.6 * 5 =", totalDivided*5)

	// Short syntax for creating arrays
	arr2 := [5]float64{98, 88, 78, 99, 97}
	// or, the trailing comma here is required
	arr3 := [5]float64{
		50,
		20,
		64,
		99,
		86,
	}
	fmt.Println(arr2, arr3)
}
