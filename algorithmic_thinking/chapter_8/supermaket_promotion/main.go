package main

import "fmt"

const (
	MAX_RECEIPTS = 1000000
	MAX_COST     = 1000000
)

type heap_element struct {
	receipt_index int
	cost          int
}

// num_heap is the counter that holds number of elelemtns in heap
// rest 2 are problem specific values
func max_heap_insert(heap *[MAX_RECEIPTS + 1]heap_element, num_heap *int, receipt_index, cost int) {
	temp := heap_element{}

	(*num_heap)++ // since we are inserting, increase the number of elelemnts by 1

	// add value to the heap
	heap[*num_heap] = heap_element{receipt_index: receipt_index, cost: cost}
	i := *num_heap

	// since we are indexing array by 1, make sure index is greater than 1,
	// continue looping till the current element is > than its parent, which means we need to swap
	// since we are indexing starting from 1, parent is i/2
	for (i > 1) && (heap[i].cost > heap[i/2].cost) {
		temp = heap[i]      // hold current value in temp
		heap[i] = heap[i/2] // replace current index pos with parents
		heap[i/2] = temp    // replace parents index with current
		i = i / 2           // replace index with parents, next we will check if parent is at the right place
	}
}

func max_heap_extract(heap *[MAX_RECEIPTS + 1]heap_element, num_heap *int) heap_element {
	var remove, temp heap_element

	remove = heap[1]          // return the root
	heap[1] = heap[*num_heap] // replace the root with bottom most right child
	(*num_heap)--
	i := 1

	// loop till we have at least left child
	for (i * 2) <= *num_heap {
		child := i * 2

		// check if child index is less than num_heap, and right child is > left child, increase
		// child index, since +1 is right child
		// basically we are chekcking if right child is greater than left child
		// if it is, increase the child index
		if (child < *num_heap) && (heap[child+1].cost > heap[child].cost) {
			child++
		}

		// if the child is greater than the cost then the heap is not sorted
		if heap[child].cost > heap[i].cost {
			temp = heap[i]
			heap[i] = heap[child]
			heap[child] = temp
			i = child
		} else {
			break
		}

	}

	return remove
}

func min_heap_insert(heap *[MAX_RECEIPTS + 1]heap_element, num_heap *int, receipt_index, cost int) {
	temp := heap_element{}

	(*num_heap)++ // since we are inserting, increase the number of elelemnts by 1

	// add value to the heap
	heap[*num_heap] = heap_element{receipt_index: receipt_index, cost: cost}
	i := *num_heap

	// since we are indexing array by 1, make sure index is greater than 1,
	// continue looping till the current element is > than its parent, which means we need to swap
	// since we are indexing starting from 1, parent is i/2
	for (i > 1) && (heap[i].cost < heap[i/2].cost) {
		temp = heap[i]      // hold current value in temp
		heap[i] = heap[i/2] // replace current index pos with parents
		heap[i/2] = temp    // replace parents index with current
		i = i / 2           // replace index with parents, next we will check if parent is at the right place
	}
}

func min_heap_extract(heap *[MAX_RECEIPTS + 1]heap_element, num_heap *int) heap_element {
	var remove, temp heap_element

	remove = heap[1]          // return the root
	heap[1] = heap[*num_heap] // replace the root with bottom most right child
	(*num_heap)--
	i := 1

	// loop till we have at least left child
	for (i * 2) <= *num_heap {
		child := i * 2

		// check if child index is less than num_heap, and right child is > left child, increase
		// child index, since +1 is right child
		// basically we are chekcking if right child is greater than left child
		// if it is, increase the child index
		if (child < *num_heap) && (heap[child+1].cost < heap[child].cost) {
			child++
		}

		// if the child is greater than the cost then the heap is not sorted
		if heap[child].cost < heap[i].cost {
			temp = heap[i]
			heap[i] = heap[child]
			heap[child] = temp
			i = child
		} else {
			break
		}

	}

	return remove
}

func main() {
	used := [MAX_RECEIPTS]bool{}
	max_heap := [MAX_RECEIPTS + 1]heap_element{}
	min_heap := [MAX_RECEIPTS + 1]heap_element{}

	var num_days, receipt_index_today, cost int
	receipt_index := 0
	var total_prizes int64

	max_num_heap, min_num_heap := 0, 0
	max_element := heap_element{}
	min_element := heap_element{}

	fmt.Printf("Please enter the number of days of compition: ")
	fmt.Scanf("%d", &num_days)

	for i := 0; i < num_days; i++ {
		fmt.Printf("Receipt index for today:")
		fmt.Scanf("%d", &receipt_index_today)
		for j := 0; j < receipt_index_today; j++ {
			fmt.Scanf("%d", &cost)
			fmt.Println("read", cost)
			fmt.Printf(
				"Inserting in max heap with num heap %d, receipt index %d cost %d\n",
				max_num_heap,
				receipt_index,
				cost,
			)
			max_heap_insert(&max_heap, &max_num_heap, receipt_index, cost)

			fmt.Printf(
				"Inserting in min heap with num heap %d, receipt index %d cost %d\n",
				min_num_heap,
				receipt_index,
				cost,
			)
			min_heap_insert(&min_heap, &min_num_heap, receipt_index, cost)
			receipt_index++
		}

		fmt.Println("Crossed that part")

		// now for each day

		fmt.Printf("Max heap number %d\n", max_num_heap)
		max_element = max_heap_extract(&max_heap, &max_num_heap)
		for used[max_element.receipt_index] {

			fmt.Printf("Max heap number %d\n", max_num_heap)
			max_element = max_heap_extract(&max_heap, &max_num_heap)
		}
		used[max_element.receipt_index] = true

		fmt.Printf("Min heap number %d\n", min_num_heap)
		min_element = min_heap_extract(&min_heap, &min_num_heap)
		for used[min_element.receipt_index] {

			fmt.Printf("Min heap number %d\n", min_num_heap)
			min_element = min_heap_extract(&min_heap, &min_num_heap)
		}
		used[min_element.receipt_index] = true

		total_prizes += int64(max_element.cost - min_element.cost)
	}
	fmt.Printf("%d\n", total_prizes)
}
