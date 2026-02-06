package main

import "fmt"

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum checks if the tree has a root-to-leaf path such that
// adding up all the values along the path equals targetSum.
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// If it's a leaf node, check if remaining sum equals node value.
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	remaining := targetSum - root.Val
	return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
}

func main() {
	// Example tree:
	//        5
	//       / \
	//      4   8
	//     /   / \
	//    11  13  4
	//   /  \      \
	//  7    2      1
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	fmt.Println("Has Path Sum 22?:", hasPathSum(root, 22)) // true (5->4->11->2)
	fmt.Println("Has Path Sum 26?:", hasPathSum(root, 26)) // true (5->8->13)
	fmt.Println("Has Path Sum 18?:", hasPathSum(root, 18)) // true (5->8->4->1)
	fmt.Println("Has Path Sum 27?:", hasPathSum(root, 27)) // false
}
