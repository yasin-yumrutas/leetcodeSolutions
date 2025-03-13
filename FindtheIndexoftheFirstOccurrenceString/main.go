package main

import (
	"fmt"
)

// strStr fonksiyonu, 'needle' dizgisinin 'haystack' dizgisi içinde ilk geçtiği indeksi döner.
// Eğer 'needle', 'haystack' içinde bulunmuyorsa -1 döner.
func strStr(haystack string, needle string) int {
	// Dizgilerin uzunluklarını alalım.
	hLen := len(haystack)
	nLen := len(needle)

	// haystack içinde needle uzunluğu kadar tarama yapıyoruz.
	for i := 0; i <= hLen-nLen; i++ {
		// Alt dizgi needle ile eşleşirse indeksi döndür.
		if haystack[i:i+nLen] == needle {
			return i
		}
	}

	// Eğer eşleşme bulunamadıysa -1 döndür.
	return -1
}

func main() {
	// Örnek kullanımlar:
	fmt.Println(strStr("sadbutsad", "sad"))  // Çıktı: 0
	fmt.Println(strStr("leetcode", "leeto")) // Çıktı: -1
}
