package bst

// IsTreeBST checks if the given tree is a Binary Search Tree or not
// not much useful in this case, cause our package will only insert
// in the tree fullfilling the Binary Search Tree conditions i.e. the
// value is left subtree will always be less than or equals than root
// and value of right subtree will always be greater than that of root
func IsTreeBST(b *BSTNode) bool {
	if b == nil || *b == (BSTNode{}) {
		return true
	}

	if isSubtreeLesser(b.Left, b.Data) &&
		isSubtreeGreater(b.Right, b.Data) &&
		IsTreeBST(b.Left) &&
		IsTreeBST(b.Right) {
		return true
	}
	return false
}

func isSubtreeLesser(b *BSTNode, v int) bool {
	if b == nil {
		return true
	}
	if *b == (BSTNode{}) {
		return true
	}

	// since this will be the inner search, i mean already inside the left subtree from
	// the caller, we need to check the right subtree as well
	if b.Data <= v &&
		isSubtreeLesser(b.Left, v) &&
		isSubtreeLesser(b.Right, v) {
		return true
	}

	return false
}

func isSubtreeGreater(b *BSTNode, v int) bool {
	if b == nil {
		return true
	}
	if *b == (BSTNode{}) {
		return true
	}

	// since this will be the inner search, i mean already inside the left subtree from
	// the caller, we need to check the right subtree as well
	if b.Data > v &&
		isSubtreeGreater(b.Left, v) &&
		isSubtreeGreater(b.Right, v) {
		return true
	}

	return false
}
