package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		b = 5
		c = 20
		d = 32
	)
	// Constants are variables whose values cannot be changed later.
	const a = "Hello world!"
	fmt.Println(a)
	fmt.Println("5 + 20 + 32 = ", b+c+d)
	doubleNum()
	temperatureConverter()
	feetToMeters()
}

func doubleNum() {
	fmt.Print("Enter a number: ")
	var input float64
	// Get value from user input
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return
	}

	output := input * 2
	fmt.Println("Result is: ", output)
	x := 5
	x += 1
	fmt.Println(x)
}

func temperatureConverter() {
	// Convert from Fahrenheit to Celsius.
	var farhInput float64
	var celsius float64
	fmt.Print("Enter Fahrenheit temperature: ")
	_, err := fmt.Scanf("%f", &farhInput)
	if err != nil {
		return
	}
	celsius = (farhInput - 32) * (5.0 / 9.0)
	fmt.Println(fmt.Sprintf("%.2f", celsius))
	// or
	fmt.Println(math.Round(celsius*100) / 100)
}

func feetToMeters() {
	// Turn user feet input to meters.
	var feetInput float64
	var meters float64
	fmt.Print("Enter amount in feet: ")
	_, err := fmt.Scanf("%f", &feetInput)
	if err != nil {
		return
	}
	meters = feetInput * 0.3048
	fmt.Println(math.Round(meters*100) / 100)
}
