package main

import (
	"fmt"
	"math"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	length := int(math.Round(float64(len(arr) / 2)))

	return merge(
		mergeSort(arr[:length]),
		mergeSort(arr[length:]),
	)
}

func merge(left, right []int) []int {
	fmt.Println(left, right)
	merged := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	merged = append(merged, left...)
	merged = append(merged, right...)

	return merged
}

func main() {
	fmt.Println(mergeSort([]int{1, 7, 3}))
}
