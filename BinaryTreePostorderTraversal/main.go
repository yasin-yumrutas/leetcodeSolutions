package main

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ---------- Recursive Solution ----------
func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)                    // Left
		dfs(node.Right)                   // Right
		result = append(result, node.Val) // Root
	}

	dfs(root)
	return result
}

// ---------- Iterative Solution (Single Stack) ----------
func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{}
	result := []int{}
	var lastVisited *TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			peek := stack[len(stack)-1]

			if peek.Right != nil && lastVisited != peek.Right {
				current = peek.Right
			} else {
				result = append(result, peek.Val)
				lastVisited = peek
				stack = stack[:len(stack)-1]
			}
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

	fmt.Println("Recursive:", postorderTraversal(root))          // [3 2 1]
	fmt.Println("Iterative:", postorderTraversalIterative(root)) // [3 2 1]
}
