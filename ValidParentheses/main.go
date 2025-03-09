package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		// Eğer kapanan parantezse
		if open, found := brackets[char]; found {
			// Stack boşsa veya eşleşme yoksa false döndürür
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Eşleşme varsa son elemanı çıkarıyoruz
			stack = stack[:len(stack)-1]
		} else {
			// Açılan parantezse stack'e ekliyoruz
			stack = append(stack, char)
		}
	}
	// Stack boşsa tüm parantezler eşleşmiştir
	return len(stack) == 0
}

func main() {
	testCases := []string{
		"()", "()[]{}", "(]", "([])", "{[()]}", "{[}]",
	}

	for _, test := range testCases {
		fmt.Printf("Input: %s -> Output: %v\n", test, isValid(test))
	}
}
