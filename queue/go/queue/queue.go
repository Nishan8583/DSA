package queue

import "errors"

type queue struct {
	first  int   // points to the first value in queue
	rear   int   // points to the last value in queue
	size   int   // size of the queue
	values []int // actual values from queue
}

// New returns a new queue instance created
func New(size int) queue {
	return queue{
		first:  -1,
		rear:   -1,
		size:   size,
		values: make([]int, size),
	}
}

// Enqueue insert value to the queue
func (q *queue) Enqueue(v int) error {
	if q.rear+1 == q.size {
		return errors.New("Queue already full")
	}

	if q.first == -1 {
		q.first = 0
	}
	q.rear += 1
	q.values[q.rear] = v
	return nil
}

// Dequeue retrieves the available first value in the queue
func (q *queue) Dequeue() (int, error) {

	// we use q.first instead of using len() cause we are setting the
	// the values to 0, which means len will think the value is still there
	// instead to allocating and deallocating the slice, we will use q.first
	// and q.rear for such purposes
	if q.first == -1 {
		return 0, errors.New("ERROR no more values to return")
	}

	// if the last element, then reset the queue
	if q.first == q.rear {
		t := q.values[q.first]

		q.first = -1
		q.rear = -1

		return t, nil
	}

	t := q.values[q.first]
	q.first = q.first + 1
	return t, nil
}
