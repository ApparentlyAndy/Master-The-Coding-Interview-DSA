package main

import (
	"strconv"
)

func findLargestPosition(array []int) int {
	largest := 0

	for _, num := range array {
		if num > largest {
			largest = num
		}
	}

	return len(strconv.Itoa(largest))
}

func getDigitPosition(num int, position int) int {
	strNum := strconv.Itoa(num)
	if position > len(strNum) {
		return 0
	}
	number, _ := strconv.Atoi(string(strNum[len(strNum)-position]))
	return number
}

func RadixSort(array []int) []int {
	count := findLargestPosition(array)

	for i := 1; i <= count; i++ {
		buckets := make([][]int, 10)

		for _, num := range array {
			key := getDigitPosition(num, i)
			buckets[key] = append(buckets[key], num)
		}

		list := []int{}

		for _, numbers := range buckets {
			list = append(list, numbers...)
		}

		array = list
	}
	return array
}
