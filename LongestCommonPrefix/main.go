package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i] // İlk kelimenin i. karakteri
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i] // Eşleşme bitti, sonucu döndür
			}
		}
	}
	return strs[0]
}

func main() {
	examples := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"interstellar", "internet", "internal"},
		{"apple", "apricot", "ap"},
		{"single"},
	}

	for _, strs := range examples {
		fmt.Printf("Input: %v\nOutput: \"%s\"\n\n", strs, longestCommonPrefix(strs))
	}
}
