package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	NUM_ATTEMPTS = 60
	MAX_PIECES   = 200000
	MAX_FLAVORS  = 200000
)

func random_piece(left, right int) int {
	n := time.Now().Unix()
	r := rand.New(rand.NewSource(n))
	if right > left {
		return r.Intn(right-left) + left
	} else if right > left {
		return r.Intn(left-right) + right
	}

	// if neither gte or smaller, than its equals, and if equals rand.Intn(0) will panic, so just return 1
	return 1
}

// this function is used to get the index for which left most value within the range
// SO lets say we want to get yokan from 3-11 slab, we pass in at least 3
// lets say favor array has values [1,5,6,9,13,14], so the return value will be index 1 because 5 is atleast 3
// for 11 we pass 12, cause we want atleast 11, and we are doing < comparison, we will return index 4, since it is ateast 12-1=11
func lowest_index(pieces []int, num_pieces int, at_least int) int {
	var low, mid, high int

	low = 0
	high = num_pieces

	for (high - low) >= 1 {
		mid = (low + high) / 2
		if (pieces[mid]) < at_least {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func num_in_range(pieces []int, num_pieces, left, right int) int {
	left_index := lowest_index(pieces, num_pieces, left) // get the left most index
	right_index := lowest_index(pieces, num_pieces, right+1)

	return right_index - left_index // return the number of flavors within the range
}

func solve(
	yokan [MAX_PIECES + 1]int,
	pieces_for_flavor *[MAX_FLAVORS + 1][]int,
	num_of_flavor [MAX_PIECES + 1]int,
	left, right int,
) {
	var attempt, rand_piece, flavor, result int
	width := right - left + 1 // width of the slab piece
	threshold := width / 3.0  // question wants us to make sure each friend can get 1/3 of the candy
	first_flavor := 0         // need to record the flvor with which the first friend was happy, so second firend does not choose the same flavor

	for attempt = 0; attempt < NUM_ATTEMPTS; attempt++ {
		rand_piece = random_piece(left, width) // get a random piece index
		flavor = yokan[rand_piece]             // random flavor
		result = num_in_range((*pieces_for_flavor)[flavor], num_of_flavor[flavor], left, right)

		// if the returned width croses 2x the threshold, then both friends are happy
		if result >= 2*threshold {
			fmt.Printf("YES\n")
			return
		}

		// we solved for first friend
		if result >= threshold {
			first_flavor = flavor
		}

	}

	// if we could not find for the first friend, no point in moving forward
	if first_flavor == 0 {
		fmt.Printf("No\n")
		return
	}

	for attempt = 0; attempt < NUM_ATTEMPTS; attempt++ {
		rand_piece = random_piece(left, width)
		flavor = yokan[rand_piece]
		result = num_in_range((*pieces_for_flavor)[flavor], num_of_flavor[flavor], left, right)

		// we solved for the second friend, if we got here, it means we solved for first firned already, that means,
		// we have solved for both first and second friend
		if result >= threshold {
			fmt.Printf("YES\n")
			return
		}

	}

	// reaching here means we could not solve the problem
	fmt.Printf("NO\n")
}

func main() {
	// test, i did not want to create a _testt file
	if len(os.Args) > 1 {
		if os.Args[1] == "--test" {
			testMyMainSuccess()
			testMyMainFailure()
			return
		}
	}
	yokan := [MAX_PIECES + 1]int{}
	num_of_flavor := [MAX_FLAVORS + 1]int{}

	pieces_for_flavor := &[MAX_FLAVORS + 1][]int{}

	var num_pieces, num_flavor, num_queries, l, r int

	fmt.Scanf("%d%d", &num_pieces, &num_flavor)

	for i := 1; i <= num_pieces; i++ {
		fmt.Scanf("%d", &yokan[i])
		num_of_flavor[yokan[i]] += 1
	}

	for i := 1; i <= num_flavor; i++ {
		pieces_for_flavor[i] = make([]int, num_of_flavor[i])
	}

	init_flavor_arrays(yokan, num_pieces, pieces_for_flavor)

	fmt.Scanf("%d", &num_queries)

	for i := 0; i <= num_queries; i++ {
		fmt.Scanf("%d%d", &l, &r)
		solve(yokan, pieces_for_flavor, num_of_flavor, l, r)
	}
}

func init_flavor_arrays(
	yokan [MAX_PIECES + 1]int,
	num_pieces int,
	pieces_for_flavor *[MAX_FLAVORS + 1][]int,
) {
	cur_of_flavor := [MAX_FLAVORS + 1]int{}
	var flavor, j int

	for i := 1; i <= num_pieces; i++ {
		flavor = yokan[i]
		j = cur_of_flavor[flavor]
		(*pieces_for_flavor)[flavor][j] = i
		cur_of_flavor[flavor]++
	}
}

func testMyMainFailure() {
	yokan := [MAX_PIECES + 1]int{}

	yokan[1] = 1
	yokan[2] = 3
	yokan[3] = 4
	yokan[4] = 2
	yokan[5] = 1
	yokan[6] = 1
	yokan[7] = 2
	yokan[8] = 4
	yokan[9] = 1
	yokan[10] = 2
	yokan[11] = 2
	yokan[12] = 4
	yokan[13] = 1
	yokan[14] = 1

	yokan[1] = 1
	yokan[2] = 2
	yokan[3] = 3
	yokan[4] = 4
	yokan[5] = 5
	yokan[6] = 6
	yokan[7] = 7
	yokan[8] = 8
	yokan[9] = 9
	yokan[10] = 20
	yokan[11] = 22
	yokan[12] = 42
	yokan[13] = 12
	yokan[14] = 13

	num_of_flavor := [MAX_FLAVORS + 1]int{}

	for i := 1; i <= 14; i++ {
		num_of_flavor[yokan[i]] += 1
	}
	num_flavor := 100 // 4
	pieces_for_flavor := [MAX_FLAVORS + 1][]int{}
	for i := 1; i <= num_flavor; i++ {
		pieces_for_flavor[i] = make([]int, num_of_flavor[i])
	}
	num_pieces := 14
	init_flavor_arrays(yokan, num_pieces, &pieces_for_flavor)

	queries := [][2]int{
		{3, 11},
		{8, 11},
		{5, 6},
		{1, 1},
		{7, 9},
		{1, 4},
	}
	for _, query := range queries {
		l := query[0]
		r := query[1]
		fmt.Printf("Checking for l=%d and r=%d\n", l, r)
		solve(yokan, &pieces_for_flavor, num_of_flavor, l, r)
	}
}

func testMyMainSuccess() {
	yokan := [MAX_PIECES + 1]int{}

	yokan[1] = 1
	yokan[2] = 3
	yokan[3] = 4
	yokan[4] = 2
	yokan[5] = 1
	yokan[6] = 1
	yokan[7] = 2
	yokan[8] = 4
	yokan[9] = 1
	yokan[10] = 2
	yokan[11] = 2
	yokan[12] = 4
	yokan[13] = 1
	yokan[14] = 1

	num_of_flavor := [MAX_FLAVORS + 1]int{}

	for i := 1; i <= 14; i++ {
		num_of_flavor[yokan[i]] += 1
	}
	num_flavor := 4
	pieces_for_flavor := [MAX_FLAVORS + 1][]int{}
	for i := 1; i <= num_flavor; i++ {
		pieces_for_flavor[i] = make([]int, num_of_flavor[i])
	}
	num_pieces := 14
	init_flavor_arrays(yokan, num_pieces, &pieces_for_flavor)

	queries := [][2]int{
		{3, 11},
		{8, 11},
		{5, 6},
		{1, 1},
		{7, 9},
		{1, 4},
	}
	for _, query := range queries {
		l := query[0]
		r := query[1]
		fmt.Printf("Checking for l=%d and r=%d\n", l, r)
		solve(yokan, &pieces_for_flavor, num_of_flavor, l, r)
	}
}
