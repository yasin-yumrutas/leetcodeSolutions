package main

import "fmt"

// TreeNode, ikili ağaç düğümünü tanımlar
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minDepth fonksiyonu, ikili ağacın minimum derinliğini döndürür.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Çocuklardan birinin puanı sıfırsa, diğer çocuğun puanını dikkate almalıyız.
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	// Her iki çocuk da mevcuttur: iki derinliğin minimumunu alın.
	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left < right {
		return left + 1
	}
	return right + 1
}

func main() {
	// Example:
	//     3
	//    / \
	//   9  20
	//      / \
	//     15  7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Minimum Depth:", minDepth(root)) // 2 (3 -> 9)

	// Another example (unbalanced):
	//   1
	//    \
	//     2
	//      \
	//       3
	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}
	root2.Right.Right = &TreeNode{Val: 3}

	fmt.Println("Minimum Depth:", minDepth(root2)) // 3
}
