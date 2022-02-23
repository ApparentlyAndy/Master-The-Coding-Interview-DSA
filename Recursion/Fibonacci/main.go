package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144

func fibonacciIterative(n int) int {
	results := []int{0, 1}
	for i := 2; i <= n; i++ {
		results = append(results, results[i-1]+results[i-2])
	}
	return results[n]
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func main() {
	fmt.Println(fibonacciRecursive(6))
	fmt.Println(fibonacciIterative(12))
}
