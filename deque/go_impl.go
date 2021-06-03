package main

import (
	"fmt"
)

type deque struct {
	values []int
}

func (d *deque) isEmtpty() bool {
	if len(d.values) == 0 {
		return true
	}

	return false
}

func (d *deque) addFront(i int) {
	temp := []int{i}
	d.values = append(temp,d.values...)
}

func (d *deque) removeFront() {
	d.values = d.values[1:]
}

func (d *deque) addRear(i int) {
	d.values = append(d.values,i)
}

func (d *deque) removeRear() {
	d.values = d.values[:len(d.values)-1]
}

func (d *deque) size() int{
	return len(d.values)
}

func main() {
	s := deque{}
	s.addFront(0);
	s.addFront(-1)
	s.addRear(1);
	s.addRear(2);
	fmt.Println(s);

	s.removeRear();
	fmt.Println(s)
	s.removeFront();
	fmt.Println(s)
}
