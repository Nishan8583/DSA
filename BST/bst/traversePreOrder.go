package bst

import "fmt"

// TraversePreOrder follows the PreOrder traversal algorithm
// Root node-> Left Node -> RightNode
func (b *BSTNode) TraversePreOrder() {

	// check if empty
	if *b == (BSTNode{}) {
		fmt.Println("Empty Node")
		return
	}

	// Since Preorder traversal, print root node first then go to left and then right
	fmt.Println("DATA -> ", b.Data)

	// First Traverse all in the left subtree
	if b.Left != nil {
		b.Left.TraversePreOrder()
	}

	// Then Traverse Right
	if b.Right != nil {
		b.Right.TraversePreOrder()
	}
}
