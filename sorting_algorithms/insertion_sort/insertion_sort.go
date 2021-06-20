package main

import "fmt"

func insertion_sort(A []int) []int {
	for i := 1; i < len(A); i++ {
		hole := i
		value := A[i]

		for {
			if hole <= 0 || A[hole-1] < value {
				break
			}
			A[hole] = A[hole-1] // shifting greater vlaue to right
			hole = hole - 1
		}
		A[hole] = value // shifting value to sorted place
	}

	return A
}

func main() {
	t := insertion_sort([]int{7, 2, 5, 3, 1})
	fmt.Println(t)
}
