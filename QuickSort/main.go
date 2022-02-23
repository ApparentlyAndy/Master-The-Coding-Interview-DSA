package main

import "fmt"

// The Pivot helper function will accept the array, start index and end index
// The start index can default to 0 and end index can be array length minus 1
func pivot(arr []int, start, end int) ([]int, int) {
	// Grab the pivot element from start of array
	pivot := arr[start]

	// Store the current pivot index in a variable
	// This will keep track where the final pivot will end up
	pivotIndex := start

	// Loop through the array, EXCEPT the first element, which is our pivot
	for i := start + 1; i < len(arr); i++ {
		// If the pivot is greater than current element
		if pivot > arr[i] {
			// Increment pivotIndex
			pivotIndex++

			// Swap pivot index and the current element
			temp := arr[i]
			arr[i] = arr[pivotIndex]
			arr[pivotIndex] = temp
		}
	}

	// Finally, swap where the pivot is now to where the pivot should be
	// Remember, the pivotIndex is how we kept track of this location
	arr[start], arr[pivotIndex] = arr[pivotIndex], arr[start]

	// Return the pivotIndex
	// In Golang, we have to return the array and pivotIndex
	return arr, pivotIndex
}

func quickSort(arr []int, left int, right int) []int {
	// If left and right are not the same
	// (There is more than 1 element in the array)
	if left < right {
		// Use our pivot function and sort the array
		array, pivotIndex := pivot(arr, left, right)
		// Now you know where the pivotIndex is...
		// Call quickSort again with the left side of the array
		// Starting from the pivotIndex - 1.
		quickSort(array, left, pivotIndex-1)
		// Call quickSort again with the right side of the array
		// Starting at the pivotIndex + 1
		quickSort(array, pivotIndex+1, right)
	}
	return arr
}

func main() {

	// [99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0]
	array := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(quickSort(array, 0, len(array)-1))
}
