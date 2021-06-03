package stack

import "testing"

func TestStackImplementation(t *testing.T) {
	s := New()

	for i := 0; i <= 10; i++ {
		s.Push(i)
	}

	for i := 10; i >= 0; i-- {
		v, _ := s.Pop()
		t.Log(v)
		if v != i {
			t.Fatal("FATAL value is not correct")
		}
	}

	//t.Log(s.GetValues())

	_, err := s.Pop()
	if err != nil {
		t.Log("ERROR given, success", err)
	}
}
