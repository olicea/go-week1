package main

import "fmt"

func main() {
	fmt.Println("add (1, 4) returned ", add(1, 4))
	fmt.Println("multiply (2, 4) returned ", multiply(2, 4))
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	var result = a * b
	return result
}
