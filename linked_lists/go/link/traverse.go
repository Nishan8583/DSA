package link

// Traverse loops through the linked list, pass the item number u want the value of
// the counting starts from 0,
// NOTE: pass -1 if you want the last item
func (l *LinkedList) Traverse(k int) (*Node, error) {

	if l.head == nil {
		return nil, ErrLinkedListIsEmpty
	}

	// let b point to first
	b := l.head
	if k == 0 {
		return b, nil
	}

	if k == -1 {
		for {
			if b.next == nil {
				return b, nil
			}
			b = b.next
		}
	}
	// cause first value already in b
	// we start the loop with i = 1
	for i := 1; i <= k; i++ {

		if b == nil || b.next == nil {
			return b, errEndOfLinkedList
		}

		b = b.next
	}
	return b, nil
}
