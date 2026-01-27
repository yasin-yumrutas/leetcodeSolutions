package main

import "fmt"

// climbStairs fonksiyonu, en tepeye çıkmanın farklı yollarının sayısını döndürür.
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev1 := 1 // f(1)
	prev2 := 2 // f(2)

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev1 = prev2
		prev2 = current
	}

	return prev2
}

func main() {
	tests := []int{1, 2, 3, 4, 5, 10}

	for _, t := range tests {
		fmt.Printf("Input: %d -> Output: %d\n", t, climbStairs(t))
	}
}
