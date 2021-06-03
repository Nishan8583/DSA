package stack

import (
	"errors"
	"log"
)

// Stack implements stack data structure implementation
type Stack struct {
	value []int
	top   int
}

// New returns a new stack structure
func New() Stack {
	return Stack{
		value: []int{},
		top:   -1,
	}
}

// Push pushes data into the stack strucutre and increments the top count
func (s *Stack) Push(a int) {
	s.value = append(s.value, a)
	s.top += 1
}

func (s *Stack) GetValues() []int {
	log.Println("This", s)
	return s.value
}

// IsEmpty checks if the stack is already empty
func (s Stack) IsEmpty() bool {
	if len(s.value) == 0 {
		return true
	}
	return false
}

// Pop pops the value from the stack and decreases the TOP
func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("Stack is empty")
	}
	v := s.value[s.top]
	s.value = s.value[:len(s.value)-1]
	s.top -= 1
	return v, nil
}
