package binary_search

import "testing"

func BenchmarkF(b *testing.B) {
	s := []int{}
	for i := 0; i < 100000000; i++ {
		s = append(s, i)
	}
	for i := 0; i < b.N; i++ {
		binary_search_recur(&s, 0, len(s)-1, 900)
	}

}

func BenchmarkP(b *testing.B) {
	s := []int{}
	for i := 0; i < 100000000; i++ {
		s = append(s, i)
	}
	for i := 0; i < b.N; i++ {
		binary_search_iter(&s, 900)
	}
}
