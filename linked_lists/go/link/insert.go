package link

import (
	"fmt"
	"log"
)

// InsertBeginning inserts at the beninning of linked list, i.e. l.head is replaced
func (l *LinkedList) InsertBeginning(t Node) {
	t.next = l.head
	l.head = &t
}

// insets at the end of LinkedList
func (l *LinkedList) InsertEnd(t Node) {

	finalNode, err := l.Traverse(-1)

	if err == ErrLinkedListIsEmpty {
		l.head = &t
		return

	}

	// if the list is not empty, and another error occurs then it is considered unhandeled for now
	if err != nil {
		log.Println("FATAL unhandled error occured = ", err)
		return
	}

	finalNode.next = &t

}

// InsertMiddle takes a certain position from the linked list, and inserts there
func (l *LinkedList) InsertMiddle(t Node, i int) {
	s, err := l.Traverse(i)
	if err != nil {
		fmt.Println("ERROR could not get the right Node ", err)
		return
	}

	t.next = s.next
	s.next = &t
}
