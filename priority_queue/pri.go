package main

import "fmt"

func heapify(arr *[]int, n int, i int) {
	// find the largest among the root, l is the left child and r is the right child
	fmt.Printf("Array %v, n is %d, and i is %i\n", *arr, n, i)
	largest := i

	l := 2*i + 1
	r := 2*i + 2

	if l < n && *arr[i] < *arr[l] {
		largest = l
	}

	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	// swap and continue heapyfing if the current node is not the largest
	if largest != i {
		t1 := arr[i]
		t2 := arr[largest]

		arr[i] = t2
		arr[largest] = t1

		heapify(arr, n, largest)

	}
}

// for insertint into the queue
func insert(arr *[]int, num int) {
	size := len(*arr)

	if size == 0 {
		*arr = append(*arr, num)
		return
	} else {
		sizeMinOne := size - 1
		*arr = append(*arr, num)
		for i := sizeMinOne; i >= 0; i-- {
			heapify(arr, size, i)
		}
	}
}

func remove(arr *[]int, num int) {
	size := len(*arr)
	i := 0
	for i; i < size; i++ {
		if num == arr[i] {
			break
		}
	}

	t1 := arr[size-1]
	t2 := arr[i]

	arr[i] = t1
	arr[size-1] = t2

	arr = append(arr[:size-1], *arr[size:]...)
	sizeminone := size - 1
	for i := sizeminone; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

}

func main() {
	arr := []int{}

	insert(&arr, 1)
	insert(&arr, 2)
	fmt.Println(arr)
}
