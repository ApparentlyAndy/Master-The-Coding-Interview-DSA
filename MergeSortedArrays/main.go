package main

import "fmt"

// Input: [0,3,4,31] [4,6,30]
// Output: [0,3,4,4,6,30,31]

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	result := []int{}

	if len(arr1) == 0 || len(arr2) == 0 {
		return append(arr1, arr2...)
	}

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] <= arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}

	result = append(result, arr1...)
	result = append(result, arr2...)

	return result
}

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 4, 31, 69, 70, 71}, []int{4, 6, 30, 66}))
}
