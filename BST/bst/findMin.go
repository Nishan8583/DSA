package bst

import "errors"

// FindMinIter is the iteration version of finding minimum value in the Binary Search Tree
func (b *BSTNode) FindMinIter() (int, error) {
	if *b == (BSTNode{}) {
		return 0, errors.New("The Binary Search Tree is empty")
	}

	if b.Left == nil {
		return b.Data, nil
	}

	temp := b.Left
	for {
		if temp.Left == nil {
			return temp.Data, nil
		}
		temp = temp.Left
	}
}

// FindMinRecur is recursive implementation of finding the minimum value in the Binary Search tree
func (b *BSTNode) FindMinRecur() (int, error) {
	if *b == (BSTNode{}) {
		return 0, errors.New("The Binary Search Tree is empty")
	}

	if b.Left == nil {
		return b.Data, nil
	}

	return b.Left.FindMinRecur()
}
