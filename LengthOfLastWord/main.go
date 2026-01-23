package main

import "fmt"

// LengthOfLastWord, bir dizenin son kelimesinin uzunluğunu döndürür.

// Bir kelime boşluk olmayan karakterlerden oluşan maksimum alt dize olarak tanımlanır.
func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1

	// Sondaki boşlukları atla
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// Son kelimenin karakter sayısını say
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	examples := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}

	for _, example := range examples {
		fmt.Printf("Input: \"%s\" -> Output: %d\n", example, lengthOfLastWord(example))
	}
}
