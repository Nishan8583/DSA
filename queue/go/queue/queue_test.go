package queue

import (
	"testing"
)

func TestQueueFunc(t *testing.T) {
	q := New(5)

	// first time
	someInternalTest(q, t)

	// second time same test should work

	someInternalTest(q, t)
}

func someInternalTest(q queue, t *testing.T) {
	for i := 0; i < 5; i++ {
		err := q.Enqueue(i)
		if err != nil {
			t.Fatalf(
				"There should not be any error here in index %d,error=%v\n",
				i, err,
			)
		}
	}

	err := q.Enqueue(90)
	if err == nil {
		t.Fatal("There should be error cause now queue already full")
	}

	for i := 0; i < 5; i++ {
		v, err := q.Dequeue()
		if err != nil {
			t.Fatalf(
				"Error value must not be recieved during this deque step at %d, error=%v\n",
				i, err,
			)
		}
		if v != i {
			t.Logf("ERROR value mismatched, expected=%d,got=%d\n", i, v)
		}
	}

	v, err := q.Dequeue()
	if err == nil {
		t.Fatalf(
			"Error value must have recieved during this deque step at %d, error=%v\n",
			v, err,
		)
	}
	if v != 0 {
		t.Logf("ERROR value mismatched, expected=%d,got=%d\n", 0, v)
	}

}
