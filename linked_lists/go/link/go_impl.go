package link

import (
	"errors"
	"log"
)

// Node represents each Node in the linked list
type Node struct {
	data int
	next *Node
}

// the begninning of linked list
type LinkedList struct {
	head *Node
}

var ErrLinkedListIsEmpty = errors.New("Linked List is EMTPY")
var errEndOfLinkedList = errors.New("END OF LINKED LIST")

func (l *LinkedList) PrintValues() {
	b := l.head
	for {
		log.Printf("Data: %v, Pointer to next:%v\n", b.data, b.next)
		b = b.next
		if b.next == nil {
			break
		}

	}
}

// EmtpyList will completely empty the Linked List
func (l *LinkedList) EmtpyList() {
	l.head = nil
}
