package main

import "fmt"

func selectionSort(arr []int) []int {
	done := false
	count := 0

	for !done {
		lowest := arr[count]
		index := 0

		for i := count; i <= len(arr)-1; i++ {
			if arr[i] < lowest {
				lowest = arr[i]
				index = i
			}
		}

		if index != 0 {
			temp := arr[count]
			arr[count] = lowest
			arr[index] = temp
		}

		if count == len(arr)-1 {
			done = true
		}

		count++
	}

	return arr
}

func main() {
	fmt.Println(selectionSort([]int{1, 6, 2, 3, 4, 8, 7, 4, 3}))
}
