package main

import (
	"fmt"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)

	// 1️⃣ Skip spaces
	for i < n && s[i] == ' ' {
		i++
	}

	// 2️⃣ Sign
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	result := 0
	INT_MAX := 2147483647
	INT_MIN := -2147483648

	// 3️⃣ Read digits
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// overflow check
		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > 7) {
			if sign == 1 {
				return INT_MAX
			}
			return INT_MIN
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func main() {

	tests := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
	}

	for _, t := range tests {
		fmt.Printf("Input: %q -> Output: %d\n", t, myAtoi(t))
	}
}
