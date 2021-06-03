package bst

import "fmt"

type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

// New returns a BSTNode
func New() BSTNode {
	return BSTNode{}
}

// Recursion example for traversing left
func (b *BSTNode) TraverseLeft() {
	if *b == (BSTNode{}) {
		return
	}

	fmt.Println("Value = ", b.Data)
	if b.Left == nil {
		return
	}
	b.Left.TraverseLeft()
}

func (b *BSTNode) TraverseRight() {
	if *b == (BSTNode{}) {
		return
	}

	fmt.Println("Value = ", b.Data)
	if b.Right == nil {
		return
	}
	b.Right.TraverseRight()
}

func (b *BSTNode) Traverse() {
	fmt.Println("Traversing LEFT")
	b.TraverseLeft()
	fmt.Println("Traversing RIGHT")
	b.TraverseRight()
}

func (b *BSTNode) TraverseNonRecur() {
	if *b == (BSTNode{}) {
		return
	}
	t := b
	fmt.Println(b.Data)
	for {
		if t.Left == nil {
			break
		}
		t = t.Left
	}
	for {
		if t.Right == nil {
			break
		}
		t = t.Right
	}
}
