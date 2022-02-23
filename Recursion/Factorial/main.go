package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return value
	}
	return value * factorialRecursive(value-1)
}

func factorialIterative(value int) int {
	result := 1
	for i := 1; i <= value; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorialRecursive(5))
	fmt.Println(factorialIterative(5))
}
