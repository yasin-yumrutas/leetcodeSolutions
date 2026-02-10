package main

import "fmt"

// singleNumber returns the element that appears only once.
func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func main() {
	tests := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}

	for _, t := range tests {
		fmt.Printf("Input: %v -> Output: %d\n", t, singleNumber(t))
	}
}
