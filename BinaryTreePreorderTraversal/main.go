package main

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ---------- Recursive Solution ----------
func preorderTraversal(root *TreeNode) []int {
	result := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val) // Root
		dfs(node.Left)                    // Left
		dfs(node.Right)                   // Right
	}

	dfs(root)
	return result
}

// ---------- Iterative Solution ----------
func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	result := []int{}

	for len(stack) > 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		// IMPORTANT: push Right first, then Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

func main() {
	// Example tree:
	//      1
	//       \
	//        2
	//       /
	//      3
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println("Recursive:", preorderTraversal(root))          // [1 2 3]
	fmt.Println("Iterative:", preorderTraversalIterative(root)) // [1 2 3]
}
