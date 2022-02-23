package main

import (
	"errors"
	"fmt"
)

func firstRecurringChar(nums []int) (int, error) {
	counter := map[int]int{}

	for _, num := range nums {
		if _, ok := counter[num]; ok {
			return num, nil
		} else {
			counter[num] = 1
		}
	}

	return -1, errors.New("no recurring char")
}

func main() {
	fmt.Println(firstRecurringChar([]int{}))
}
