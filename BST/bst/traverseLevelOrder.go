package bst

import "fmt"

// TraverseLevelOrder
func (b *BSTNode) TraverseLevelOrder() {

	// if an empty node, just exist
	if *b == (BSTNode{}) {
		return
	}

	// if only root node, no need for extra allocation and steps
	if b.Left == nil && b.Right == nil {
		fmt.Println("DATA->", b.Data)
		return
	}

	// implementing sort of queue, lol
	v := []*BSTNode{}
	v = append(v, b)

	for {
		if len(v) == 0 {
			break
		}

		temp := v[0]

		// poppping of first value
		v = v[1:]

		fmt.Println("DATA->", temp.Data)
		if temp.Left != nil {
			v = append(v, temp.Left)
		}
		if temp.Right != nil {
			v = append(v, temp.Right)
		}
	}
}
