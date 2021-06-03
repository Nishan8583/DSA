package bst

import (
	"errors"
	"fmt"
)

// FindMaxIter is the iteration version of finding maximum value in the Binary Search Tree
func (b *BSTNode) FindMaxIter() (int, error) {
	if *b == (BSTNode{}) {
		return 0, errors.New("The Binary Search Tree is empty")
	}

	if b.Right == nil {
		return b.Data, nil
	}

	temp := b.Right
	for {
		if temp.Right == nil {
			return temp.Data, nil
		}
		temp = temp.Right
	}
}

// FindMaxRecur is recursive implementation of finding the maximum value in the Binary Search tree
func (b *BSTNode) FindMaxRecur() (int, error) {
	if *b == (BSTNode{}) {
		return 0, errors.New("The Binary Search Tree is empty")
	}

	if b.Right == nil {
		fmt.Println("FOUND")
		return b.Data, nil
	}

	return b.Right.FindMaxRecur()
}
