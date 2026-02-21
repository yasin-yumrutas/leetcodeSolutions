package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists (reversed order).
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}

	return dummy.Next
}

// helper: create list from slice (e.g., [2,4,3] -> 2->4->3)
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

// helper: print list
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
	// Example:
	// l1 = [2,4,3] (342)
	// l2 = [5,6,4] (465)
	// output = [7,0,8] (807)
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})

	fmt.Print("l1: ")
	printList(l1)
	fmt.Print("l2: ")
	printList(l2)

	res := addTwoNumbers(l1, l2)
	fmt.Print("sum: ")
	printList(res) // 7 -> 0 -> 8
}
