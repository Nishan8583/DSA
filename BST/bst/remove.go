package bst

// Delete will take a value, and delete that value from the node
// it will return a new tree, so make sure to replace ur old tree
func (b *BSTNode) Delete(value int) *BSTNode {

	if value < b.Data {
		b.Left = b.Left.Delete(value)
	} else if value > b.Data {
		b.Right = b.Right.Delete(value)
	} else {

		// this means the value is equals now

		// CASE 1 i.e. no child
		if b.Left == nil && b.Right == nil {
			return nil
		}

		// CASE 2 i.e. one child
		if b.Left == nil {
			return b.Right // ignoring the current node, so when we return the right, previous node will be deleted
		} else if b.Right == nil {
			return b.Left
		} else {

			//CASE 3 i.e. Two childs
			minNode := b.Right.findMinNode()
			b.Data = minNode.Data
			b.Right = b.Right.Delete(minNode.Data)
		}

	}
	return b
}

func (b *BSTNode) findMinNode() *BSTNode {
	temp := b
	for {
		if b.Left == nil {
			return temp
		}
		temp = b.Left
	}
}
