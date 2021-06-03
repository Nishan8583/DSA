package bst

import "fmt"

// Search takes an integer and checks if the value exists in the binary tree
func (b *BSTNode) Search(value int) bool {
	if *b == (BSTNode{}) {
		fmt.Println("The BST is emtpty")
		return false
	}

	// we have to add check to see if the pointers to right and left are not null
	// since this is a method call, and not a function call
	if b.Data == value {
		return true
	} else if value > b.Data && b.Right != nil {
		return b.Right.Search(value)
	} else if value < b.Data && b.Left != nil {
		return b.Left.Search(value)
	} else {
		return false
	}
}

// Find Node similar stuff like search, but it returs the node found
func (b *BSTNode) FindNode(value int) *BSTNode {
	if *b == (BSTNode{}) {
		fmt.Println("The BST is emtpty")
		return nil
	}

	// we have to add check to see if the pointers to right and left are not null
	// since this is a method call, and not a function call
	if b.Data == value {
		return b
	} else if value > b.Data && b.Right != nil {
		return b.Right.FindNode(value)
	} else if value < b.Data && b.Left != nil {
		return b.Left.FindNode(value)
	} else {
		return nil
	}
}
