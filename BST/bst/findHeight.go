package bst

func (b *BSTNode) FindHeight() int {

	// if this was the leaf node
	if b.Left == nil && b.Right == nil {
		return -1
	}
	if b.Left == nil {
		return max(-1, b.Right.FindHeight()) + 1
	} else if b.Right == nil {
		return max(b.Left.FindHeight(), -1) + 1
	} else {
		return max(b.Left.FindHeight(), b.Right.FindHeight()) + 1
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
