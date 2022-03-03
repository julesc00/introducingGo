package main

import "fmt"

var msg = "Hi there my friend"

func main() {
	fmt.Println("Your message has", len(msg), "words.")
	stringIndexing()
}

func stringIndexing() {
	fmt.Println(msg[3:5])
}

func stringCondition() {
	a, b := "Hello", "World"
	fmt.Println(a == b)
}
