package main

import "fmt"

// TreeNode, ikili ağaç düğümünü tanımlar.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric fonksiyonu, ikili ağacın merkez etrafında simetrik olup olmadığını kontrol eder.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror fonksiyonu, iki ağacın birbirinin ayna görüntüsü olup olmadığını kontrol eder.
func isMirror(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}

func main() {
	// Symmetric example:
	//        1
	//      /   \
	//     2     2
	//    / \   / \
	//   3  4  4  3
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}

	fmt.Println("Is Symmetric?:", isSymmetric(root)) // true
}
