package main

import "fmt"

func insertionSort(arr []int) []int {
	pointerIndex := 2

	for i := pointerIndex; i <= len(arr)-1; i++ {
		subArray := arr[:pointerIndex]
		for j := 0; j < len(subArray); j++ {
			if arr[pointerIndex] < subArray[j] {
				temp := subArray[j]
				subArray[j] = arr[pointerIndex]
				arr[pointerIndex] = temp
			}
		}
		pointerIndex++
	}

	return arr
}

func main() {
	fmt.Println(insertionSort([]int{1, 6, 5, 3, 8, 7, 2, 9, 7}))
}
