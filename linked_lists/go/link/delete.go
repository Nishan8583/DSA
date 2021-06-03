package link

import "fmt"

// delete from front
func (l *LinkedList) DeleteFront() {
	if l.head == nil {
		fmt.Println("Linked list is already empty")
		return
	}

	t := l.head.next
	l.head = t
}

// delete from end
func (l *LinkedList) DeleteEnd() {
	t := l.head
	for {
		if t == nil {
			fmt.Println("ERROR linked list is already empty, could not delete from end")
			break
		}
		s := t
		// hold next one temp
		t = t.next

		if t == nil {
			s.next = nil
			break
		}
	}
}

func (l *LinkedList) DeleteMid(i int) {
	s2, err := l.Traverse(i)
	s1, err := l.Traverse(i - 1)
	if err != nil {
		fmt.Println("Could not traverse to the required position", i, err)
		return
	}

	temp := s2.next
	s1.next = temp
}
