package main

import (
	"fmt"
)

// ListNode tanımlaması
type ListNode struct {
	Val  int
	Next *ListNode
}

// İki sıralı bağlı listeyi birleştiren fonksiyon
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Geçici bir baş düğüm (dummy node)
	tail := dummy        // Kuyruğu takip etmek için tail kullanıyoruz

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Kalan düğümleri ekle
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next // Birleştirilmiş listenin başını döndür
}

// Bağlı listeyi yazdıran yardımcı fonksiyon
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

// Test fonksiyonu
func main() {
	// Örnek bağlı listeler
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList) // Çıktı: 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil
}
