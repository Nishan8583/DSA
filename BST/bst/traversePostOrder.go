package bst

import "fmt"

// TraversePostOrder traverse using PostOrder algorithm
// Right Subtree -> Root Node -> left Subtree
func (b *BSTNode) TraversePostOrder() {

	if *b == (BSTNode{}) {
		fmt.Println("Node is empty")
		return
	}

	if b.Right != nil {
		b.Right.TraversePostOrder()
	}

	fmt.Println("DATA->", b.Data)

	if b.Left != nil {
		b.Left.TraversePostOrder()
	}
}
