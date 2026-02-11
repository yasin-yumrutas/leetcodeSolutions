package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle checks if a linked list has a cycle.
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	// Example with cycle:
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // cycle here

	fmt.Println("Has Cycle (expected true):", hasCycle(node1))

	// Example without cycle:
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	a.Next = b

	fmt.Println("Has Cycle (expected false):", hasCycle(a))
}
