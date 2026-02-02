package main

import "fmt"

// TreeNode, ikili ağaç düğümünü tanımlar
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth fonksiyonu, ikili ağacın maksimum derinliğini döndürür.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	// Example tree:
	//       3
	//      / \
	//     9  20
	//        / \
	//       15  7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Maximum Depth:", maxDepth(root)) // Output: 3
}
