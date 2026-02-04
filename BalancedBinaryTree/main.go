package main

import "fmt"

// TreeNode, ikili ağaç düğümünü tanımlar.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced fonksiyonu, ikili ağacın yükseklik açısından dengeli olup olmadığını kontrol eder.
func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

// `height` fonksiyonu ağacın yüksekliğini döndürür.
// Alt ağaç dengesiz ise -1 (erken durdurma) değeri döndürür.
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftH := height(node.Left)
	if leftH == -1 {
		return -1
	}

	rightH := height(node.Right)
	if rightH == -1 {
		return -1
	}

	if leftH-rightH > 1 || rightH-leftH > 1 {
		return -1
	}

	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func main() {
	// Balanced example:
	//      3
	//     / \
	//    9  20
	//       / \
	//      15  7
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Tree1 Balanced?:", isBalanced(root1)) // true

	// Unbalanced example:
	//      1
	//     /
	//    2
	//   /
	//  3
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 3}

	fmt.Println("Tree2 Balanced?:", isBalanced(root2)) // false
}
