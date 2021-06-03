package bst

import "fmt"

// TraverseInOrder traverses using InOrder method of traversal
// Left Subtree-> Root Note -> Right Subtree
func (b *BSTNode) TraverseInOrder() {

	if *b == (BSTNode{}) {
		fmt.Println("The node is emtpy")
		return
	}

	if b.Left != nil {
		b.Left.TraverseInOrder()
	}

	fmt.Println("DATA ->", b.Data)

	if b.Right != nil {
		b.Right.TraverseInOrder()
	}

}
