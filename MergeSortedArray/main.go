package main

import "fmt"

// `merge` komutu, `nums2` dizisini `nums1` dizisine tek bir sıralı dizi olarak (yerinde) birleştirir.
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1     // nums1'in dolu kısmının sonu
	j := n - 1     // nums2'nin sonu
	k := m + n - 1 // nums1'in en son indexi

	for j >= 0 { // nums2 bitene kadar devam
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println("Before:", nums1)
	merge(nums1, m, nums2, n)
	fmt.Println("After: ", nums1)
}
