package main

import "fmt"

// plusOne, digits dilimiyle temsil edilen sayıya bir ekler
// ve elde edilen rakamları döndürür.
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// Eğer tüm rakamlar 9 olsaydı (örneğin, 999 -> 1000)
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func main() {
	examples := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9, 9, 9},
	}

	for _, example := range examples {
		fmt.Printf("Input: %v -> Output: %v\n", example, plusOne(example))
	}
}
