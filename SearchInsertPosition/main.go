package main

import (
	"fmt"
)

// searchInsert fonksiyonu sıralı bir dizide hedef (target) değerinin indeksini döndürür.
// Eğer hedef dizi içinde yoksa, hedefin dizide olması gereken doğru indeksi döndürür.
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // hedef bulunduğunda indeksi döndür
		} else if nums[mid] < target {
			left = mid + 1 // hedef büyükse sağa doğru ilerle
		} else {
			right = mid - 1 // hedef küçükse sola doğru ilerle
		}
	}

	// hedef bulunamazsa olması gereken konumu döndür
	return left
}

func main() {
	// Örnek test durumları:
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // Çıktı: 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // Çıktı: 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // Çıktı: 4
}
