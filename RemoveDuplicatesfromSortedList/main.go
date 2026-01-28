package main

import "fmt"

// ListNode, tek yönlü bağlantılı bir liste düğümünü tanımlar
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates fonksiyonu, sıralı bağlantılı listeden yinelenen öğeleri kaldırır.
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

// Dilimden bağlantılı liste oluşturmak için yardımcı fonksiyon
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// Bağlı listeyi yazdırmak için yardımcı fonksiyon
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	nums := []int{1, 1, 2, 3, 3}
	head := createList(nums)

	fmt.Print("Input:  ")
	printList(head)

	result := deleteDuplicates(head)

	fmt.Print("Output: ")
	printList(result)
}
