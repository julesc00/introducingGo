package main

import (
	"fmt"
)

func main() {
	lengthSlice()
	arr1()
	returnSmallestNum()
}

func lengthSlice() {
	slice1 := make([]int, 3, 9)
	fmt.Println(len(slice1))
	fmt.Println(slice1)
}

func arr1() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func returnSmallestNum() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallestNum := x[0]
	for _, el := range x {
		if el < smallestNum {
			smallestNum = el
		}
	}
	fmt.Println("the smallest number is:", smallestNum)

}
