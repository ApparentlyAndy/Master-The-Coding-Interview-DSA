package main

import "fmt"

func bubbleSort(arr []int) []int {
	done := false
	count := 0
	for !done {
		switched := false
		for i, num := range arr {
			if i == len(arr)-1 {
				continue
			}
			if num > arr[i+1] {
				arr[i] = arr[i+1]
				arr[i+1] = num
				switched = true
			}
		}
		if !switched {
			done = true
		}
		count++
	}
	fmt.Printf("I ran %d times.\n", count)
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{1, 7, 4, 8, 2, 3, 7, 8, 3, 3, 4, 5, 456, 35, 45, 7, 35, 4, 57, 568, 45}))
}
