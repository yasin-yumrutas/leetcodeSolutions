package main

import "fmt"

// reverse reverses digits of a 32-bit signed integer.
// If overflow occurs, it returns 0.
func reverse(x int) int {
	result := 0

	for x != 0 {
		pop := x % 10
		x /= 10

		// overflow check
		if result > 214748364 || (result == 214748364 && pop > 7) {
			return 0
		}
		if result < -214748364 || (result == -214748364 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}

	return result
}

func main() {
	tests := []int{
		123,
		-123,
		120,
		1534236469, // overflow case
	}

	for _, t := range tests {
		fmt.Printf("Input: %d -> Output: %d\n", t, reverse(t))
	}
}
