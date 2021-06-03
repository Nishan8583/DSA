package main

import (
	"fmt"
	"errors"
)
type queue struct {
	front int   // points to the first element
	rear  int   // points to the final element
	size  int   // size of the queue
	value []int // holds the actual values
}

// new makes a new array
func new(size int) queue {
	return queue{
		front: -1,
		rear:  -1,
		size:  size,
		value: make([]int, size),
	}
}

func (q *queue) enqueue(v int) error {

	// example, if hread = 0, size = 5, rear = 4, then
	// the result of the following operation will also be zero
	// thus queue was full, but if a value was dequeued, then
	// head would point to 1, and 0 and 1 would not be empty
	// thus queue would not be full
	if (q.rear+1)%q.size == q.front {
		return errors.New("ERROR the circular queue is empty")
	}

	// if this is the first insert
	if q.front == -1 {
		q.front = 0
		q.rear = 0
		q.value[q.rear] = v
		return nil
	}

	// since we already checked if queue is full in beginning
	// and we need to add the value to the last space
	q.rear = (q.rear + 1) % q.size
	q.value[q.rear] = v
	return nil
}

// dequeue does dequeue operation on a circular queue
func (q *queue) dequeue() (int, error) {
	if q.front == -1 {
		return 0, errors.New("The queue is empty")
	}

	if q.front == q.rear{
		v := q.value[q.front]
		q.value[q.front] = 0
		q.front =-1
		q.rear = -1
		return v,nil
	} else {
		v := q.value[q.front]
		q.value[q.front] = 0
		q.front = (q.front + 1) % q.size
		return v,nil
	}
}

func main() {
	q := new(5)

	for i := 0; i < 5; i++ {
		err := q.enqueue(i)
		if err != nil {
			fmt.Println("ERROR could not enqueue", err)
			return
		}
	}
	fmt.Println(q.value)
	v,err := q.dequeue()
	fmt.Println(v,err)
	err = q.enqueue(90)
	fmt.Println(q.value)
	for i:= 0; i <5; i++ {
		v,err := q.dequeue()
		fmt.Println(v)
		if err != nil {
			fmt.Println("FATAL could not dequeue due to error",err,q)
			return
		}
	}
	fmt.Println(q)
}
