package main

import "fmt"

func can_make_distance(distance int, rocks []int, num_rocks int, num_remove int, length int) bool {
	prev_rock_location := 0
	cur_rock_location := 0
	removed := 0
	if length < distance { // means we have
		return false
	}

	for i := 0; i < num_rocks; i++ {
		cur_rock_location = rocks[i]

		// if the distance between current rock and previous rock is less than the distance, then we must remove the current rock
		if (cur_rock_location - prev_rock_location) < distance {
			removed++
		} else {
			prev_rock_location = cur_rock_location
		}
	}

	// if the distance between previous rock and last rock is less than the distacne then remove it
	if (length - prev_rock_location) < distance {
		removed++
	}
	return removed <= num_remove
}

func sovle(
	rocks []int,
	num_rocks int,
	num_remove int,
	length int,
) {
	var low, mid, high int
	low = 0
	high = length + 1

	for (high - low) > 1 {
		mid = (low + high) / 2
		fmt.Println("Trying ", mid)
		if can_make_distance(mid, rocks, num_rocks, num_remove, length) {
			fmt.Println("Success", mid)
			low = mid
		} else {
			fmt.Println("failure", mid)
			high = mid
		}
	}

	fmt.Printf("%d\n", low)
}

func main() {
	rocks := []int{2, 4, 5, 8}

	// Min distance betweenr rocks must be 6, pass in rocks, length of total rocks besides 0 and last, number of rocks u can remove, and last rock is always at 12
	fmt.Println("Can make ", can_make_distance(6, rocks, len(rocks), 2, 12))

	// Min distance betweenr rocks must be 3, pass in rocks, length of total rocks besides 0 and last, number of rocks u can remove, and last rock is always at 12
	fmt.Println("Can make ", can_make_distance(3, rocks, len(rocks), 2, 12))

	sovle(rocks, len(rocks), 2, 12)
}
