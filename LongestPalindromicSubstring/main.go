package main

import "fmt"

// longestPalindrome returns the longest palindromic substring in s.
// Approach: Expand Around Center (O(n^2) time, O(1) extra space).
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		l1, r1 := expand(s, i, i)   // odd length center
		l2, r2 := expand(s, i, i+1) // even length center

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// expand expands around the center (left, right) and returns the bounds of the palindrome.
func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// when loop ends, left/right are one step beyond palindrome
	return left + 1, right - 1
}

func main() {
	tests := []string{
		"babad",            // "bab" or "aba"
		"cbbd",             // "bb"
		"a",                // "a"
		"ac",               // "a" or "c"
		"forgeeksskeegfor", // "geeksskeeg"
	}

	for _, t := range tests {
		fmt.Printf("Input: %q -> Output: %q\n", t, longestPalindrome(t))
	}
}
