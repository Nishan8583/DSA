package link

import "testing"

func TestTraverse(t *testing.T) {

	l := LinkedList{}
	n1 := Node{data: 0}
	n2 := Node{data: 1}
	n3 := Node{data: 2}
	l.InsertEnd(n1)
	l.InsertEnd(n2)
	l.InsertEnd(n3)
	t.Log(l)

	for i := 0; i < 3; i++ {
		temp, err := l.Traverse(i)
		t.Log(temp, err)
	}

	n4 := Node{data: 4}
	l.InsertMiddle(n4, 1)

	for i := 0; i < 4; i++ {
		temp, err := l.Traverse(i)
		t.Log(temp, err)
	}

	t.Log("")
	l.DeleteFront()
	l.DeleteEnd()
	l.DeleteMid(1)
	for i := 0; i < 2; i++ {
		temp, err := l.Traverse(i)
		t.Log(temp, err)
	}

}
