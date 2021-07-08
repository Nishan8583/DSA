package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func print_all(head *node) {
	n := head

	for n != nil {
		fmt.Println(n)
		n = n.next
	}
}

/*
1 - 2 - 3

n = 1
next = &2
1 - nil
prev = &1

n = 2
next = 3
2-1-nil
prev = &2

n = 3
next = nil
3-2-1-nil
since next == nil
head - 3 - 2- 1- nil
*/
func reverse_lined_list(head *node) *node {

	// prev will hold the previous node address
	// and next will hold the next node address
	var prev, next *node

	// will hold the current note
	n := head

	for {
		next = n.next

		n.next = prev

		prev = n

		// if this was the final loop
		if next == nil {
			head = n
			break
		} else {
			n = next
		}
	}
	return head
}
func main() {
	n1 := node{
		value: 1,
		next:  nil,
	}

	head := &n1

	n2 := node{
		value: 2,
		next:  nil,
	}

	n3 := node{
		value: 3,
		next:  nil,
	}
	head.next = &n2
	n2.next = &n3
	fmt.Println("Before reversing")
	print_all(head)
	head = reverse_lined_list(head)
	fmt.Println("After reversing")
	print_all(head)
}
