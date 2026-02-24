package main

import "fmt"

// lengthOfLongestSubstring returns the length of the longest substring
// without repeating characters.
func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[byte]int) // char -> last seen index
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// If ch is inside the current window, move left
		if idx, ok := lastIndex[ch]; ok && idx >= left {
			left = idx + 1
		}

		lastIndex[ch] = right

		// window length = right-left+1
		curLen := right - left + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

func main() {
	tests := []string{
		"abcabcbb", // 3 ("abc")
		"bbbbb",    // 1 ("b")
		"pwwkew",   // 3 ("wke")
		"",         // 0
		"dvdf",     // 3 ("vdf")
	}

	for _, t := range tests {
		fmt.Printf("Input: %q -> Output: %d\n", t, lengthOfLongestSubstring(t))
	}
}
