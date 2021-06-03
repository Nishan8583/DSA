package bst

import "fmt"

// FindSuccessor solves a common problem in binary search tree which is to find the inorder
// successor of a give node i.e. the next value just greater than the current value
// the following code could be simplified if a extra pointer for parent was added
// to the implementaion structure of binary tree
func (b *BSTNode) FindSuccessor(value int) *BSTNode {

	currentNode := b.FindNode(value)
	fmt.Println("FOUND NODE", currentNode)
	if currentNode == nil {
		return nil
	}

	// CASE 1 if right node is not nil, then the next successor
	// will be the lowest value of right subnode
	if currentNode.Right != nil {
		return currentNode.Right.findMinNode()
	} else {

		// CASE 2 if right Subtree is present then
		// the successor will be the parent a bit just higher
		// than this nodes value
		ansestor := b                // currently root node
		var successor *BSTNode = nil // set to nil in beginning
		for {
			if ansestor == currentNode {
				break
			}
			if ansestor.Data > currentNode.Data {
				successor = ansestor

				// after that we try to search if any other values are less than
				// the current successor, so we go left
				ansestor = ansestor.Left
			} else {
				// we did not find a node with data greater, so
				// to find greter value we must move to right
				ansestor = ansestor.Right
			}
		}
		return successor
	}
}
