package main

import "fmt"

// mySqrt, x'in tam sayı karekökünü (taban değerini) döndürür.
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	result := 0

	for left <= right {
		mid := left + (right-left)/2

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	tests := []int{0, 1, 4, 8, 16, 26}

	for _, t := range tests {
		fmt.Printf("Input: %d -> Output: %d\n", t, mySqrt(t))
	}
}
