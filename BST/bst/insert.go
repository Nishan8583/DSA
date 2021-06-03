package bst

// Insert will add another node into the tree
func (b *BSTNode) Insert(value int) {

	// first checking if is an empty tree
	if *b == (BSTNode{}) {
		b.Data = value
		return
	}

	if value > b.Data {
		// this sepcific check is essential, cause we are using methods
		// if simple functions with just parameter passing was used, then
		// this check would no longer be needed
		if b.Right == nil {
			b.Right = &BSTNode{
				Data: value,
			}
			return
		} else {
			// this condition executes if there was alredy a right node
			b.Right.Insert(value)
		}
	}

	if value <= b.Data {
		if b.Left == nil {
			b.Left = &BSTNode{
				Data: value,
			}
		} else {
			b.Left.Insert(value)
		}
	}
}
