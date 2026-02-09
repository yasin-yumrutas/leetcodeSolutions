package main

import (
	"fmt"
	"unicode"
)

// isPalindrome checks if s is a palindrome considering only alphanumeric characters
// and ignoring case.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// skip non-alphanumeric from left
		for left < right && !isAlphaNum(rune(s[left])) {
			left++
		}
		// skip non-alphanumeric from right
		for left < right && !isAlphaNum(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNum(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func main() {
	tests := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
		"0P",
	}

	for _, t := range tests {
		fmt.Printf("Input: %q -> Output: %v\n", t, isPalindrome(t))
	}
}
