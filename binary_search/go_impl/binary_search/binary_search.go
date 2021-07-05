package binary_search

import "fmt"

func binary_search_iter(a *[]int, x int) int {
	start := 0
	end := len(*a)

	for start <= end {
		mid := start + (end-start)/2
		if (*a)[mid] == x {
			return mid
		} else if x > (*a)[mid] {

			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func binary_search_recur(a *[]int, start int, end int, x int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if x == (*a)[mid] {
		return mid
	} else if x > (*a)[mid] {
		start = mid + 1
	} else {
		end = mid - 1
	}
	return binary_search_recur(a, start, end, x)
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("found in ", binary_search_recur(&s, 0, len(s), 6))
}
