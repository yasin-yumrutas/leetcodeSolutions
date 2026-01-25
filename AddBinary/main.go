package main

import "fmt"

// addBinary adds two binary strings and returns their sum as a binary string.
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = string(byte(sum%2+'0')) + result
		carry = sum / 2
	}

	return result
}

func main() {
	tests := [][]string{
		{"11", "1"},
		{"1010", "1011"},
		{"0", "0"},
	}

	for _, t := range tests {
		fmt.Printf("Input: %s + %s -> Output: %s\n", t[0], t[1], addBinary(t[0], t[1]))
	}
}
