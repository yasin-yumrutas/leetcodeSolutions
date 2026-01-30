package main

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal (Iterative) - Sol -> Kök -> Sağ
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		// en sola kadar in
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// stack'ten bir düğüm al
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, cur.Val)

		// sağa geç
		cur = cur.Right
	}

	return result
}

// (Opsiyonel) Recursive çözüm
func inorderTraversalRec(root *TreeNode) []int {
	result := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

func main() {
	// Example tree:
	//     1
	//      \
	//       2
	//      /
	//     3
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println("Iterative:", inorderTraversal(root))    // [1 3 2]
	fmt.Println("Recursive:", inorderTraversalRec(root)) // [1 3 2]
}
