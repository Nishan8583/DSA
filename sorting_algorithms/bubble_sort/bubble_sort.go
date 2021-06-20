package main

import (
	"fmt"
)

func bubbleSort(array []int) []int {
	for i := range array {
		fmt.Println(i)
		for k := 0; k < len(array)-1; k++ {
			if array[k] > array[k+1] {
				temp := array[k]
				array[k] = array[k+1]
				array[k+1] = temp
			}
		}
	}
	return array
}

func main() {
	s := []int{2, 7, 4, 1, 3}
	fmt.Println(bubbleSort(s))
}
