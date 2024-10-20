/*
1. Find min cost
2. We do BFS
3. Optimal solution maintains 2 array, to reduce number of edges
Honestly too much to explain, look at book, CHapter 5, Rope CLimb problem
*/

package main

import "fmt"

const SIZE = 1000000

type position struct {
	height int
	state  int
}
type (
	board     [SIZE * 2][2]int   // this is our rope
	positions [SIZE * 2]position // array that holds all the positions we have been in
)

func find_distance(target_height, jump_distance int, itching [SIZE * 2]int, min_moves *board) {
	cur_positions := positions{} // available positions we have in the current iteration
	new_positions := positions{} // the new positions that we can use in the next iteration

	var num_cur_positions, num_new_positions int

	// we loop till we reach the 2x of the target_height, because till 2xtarget_height
	// we can fall back into the target_height accoring to the question
	// but beyond that we can not fall back into our target
	// set values to -1, indicatnig we have not reached this position yet
	for i := 0; i < target_height*2; i++ {
		min_moves[i][0] = -1
		min_moves[i][1] = -1
	}

	min_moves[0][0] = 0 // it takes 0 moves to get where we are

	cur_positions[0] = position{0, 0} // first we are at 0 length of rope
	num_cur_positions = 1             // since we are at 0 we have 1 positions i.e. 0

	for num_cur_positions > 0 {
		fmt.Println("Number of available positions", num_cur_positions)
		num_new_positions = 0 // reset new positions to 0 for next loop

		for i := 0; i < num_cur_positions; i++ {
			from_height := cur_positions[i].height // setting the starting height to be one of the current positions we are i
			fmt.Println("Checking height", from_height)
			from_state := cur_positions[i].state

			if from_state == 0 {
				add_position_up(from_height, from_height+jump_distance, target_height*2-1,
					&new_positions, &num_new_positions, itching, min_moves)

				add_position_01(from_height, &new_positions, &num_new_positions, min_moves)
			} else {
				add_position_down(from_height, from_height-1, &cur_positions, &num_cur_positions, min_moves)

				add_position_10(from_height, &cur_positions, &num_cur_positions, itching, min_moves)
			}
		}

		// update the current positions that we have
		num_cur_positions = num_new_positions
		for i := 0; i < num_cur_positions; i++ {
			cur_positions[i] = new_positions[i]
		}
	}
	fmt.Printf("Target height %d and move cost %d\n", target_height, min_moves[target_height])
}

// add_position_up signals that we are moving up the rope
// Increase the cost distance fir the value has not been updated i.e. still -1, in the rope array
func add_position_up(
	from_height, to_height, max_height int,
	pos *positions,
	num_pos *int,
	itching [SIZE * 2]int,
	min_moves *board,
) {
	distance := 1 + min_moves[from_height][0] // rope 0 is jump up rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array

	// if we are within rope limit
	// if we are not in the itching rope section
	// if the current height (rope array) has not been visited yet
	// If the cost already set was of larger value than the cost we just found out
	if (to_height <= max_height && itching[to_height] == 0) &&
		(min_moves[to_height][0] == -1 || min_moves[to_height][0] > distance) {

		min_moves[to_height][0] = distance // update the min_moves for the rope index, add 1+previous height value
		// fmt.Println("Increasing moves value for ", to_height, "to", 1+min_moves[from_height])
		pos[*num_pos] = position{
			to_height,
			0,
		} // add the position we just visited as a new position to visit in the next loop in the array
		(*num_pos)++ // since we increased the number of new positions we have, increase the counter too
	}
}

// add_position_down() signals we are moving down, hence update the state array
func add_position_down(
	from_height, to_height int,
	pos *positions,
	num_pos *int,
	min_moves *board,
) {
	distance := min_moves[from_height][1] // rope 1 is jump down rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array

	// if we are within rope limit
	// if we are not in the itching rope section
	// if the current height (state array) has not been visited yet
	// If the cost already set was of larger value than the cost we just found out
	if (to_height >= 0) &&
		(min_moves[to_height][1] == -1 || min_moves[to_height][1] > distance) {

		min_moves[to_height][1] = distance // update the min_moves for the rope index, add 1+previous height value
		// fmt.Println("Increasing moves value for ", to_height, "to", 1+min_moves[from_height])
		pos[*num_pos] = position{
			to_height,
			1,
		} // add the position we just visited as a new position to visit in the next loop in the array
		(*num_pos)++ // since we increased the number of new positions we have, increase the counter too
	}
}

// creating edge from 0 (jump rope) to 1 (state rope)
func add_position_01(
	from_height int,
	pos *positions,
	num_pos *int,
	min_moves *board,
) {
	distance := 1 + min_moves[from_height][0] // rope 1 is jump down rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array

	// if we are within rope limit
	// if we are not in the itching rope section
	// if the current height (state array) has not been visited yet
	// If the cost already set was of larger value than the cost we just found out
	if min_moves[from_height][1] == -1 || min_moves[from_height][1] > distance {

		min_moves[from_height][1] = distance // update the min_moves for the rope index, add 1+previous height value
		// fmt.Println("Increasing moves value for ", to_height, "to", 1+min_moves[from_height])
		pos[*num_pos] = position{
			from_height,
			1,
		} // add the position we just visited as a new position to visit in the next loop in the array
		(*num_pos)++ // since we increased the number of new positions we have, increase the counter too
	}
}

func add_position_10(
	from_height int,
	pos *positions,
	num_pos *int,
	itching [SIZE * 2]int,
	min_moves *board,
) {
	distance := min_moves[from_height][1] // rope 1 is jump down rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array  // rope 1 is jump up rope array

	// if we are within rope limit
	// if we are not in the itching rope section
	// if the current height (state array) has not been visited yet
	// If the cost already set was of larger value than the cost we just found out
	if (itching[from_height] == 0) && min_moves[from_height][0] == -1 ||
		min_moves[from_height][0] > distance {

		min_moves[from_height][0] = distance // update the min_moves for the rope index, add 1+previous height value
		// fmt.Println("Increasing moves value for ", to_height, "to", 1+min_moves[from_height])
		pos[*num_pos] = position{
			from_height,
			0,
		} // add the position we just visited as a new position to visit in the next loop in the array
		(*num_pos)++ // since we increased the number of new positions we have, increase the counter too
	}
}

func solve(target_height int, min_moves board) {
	best := -1
	fmt.Println("Searching for height", target_height)
	for i := target_height; i < target_height*2; i++ {
		if min_moves[i][0] != -1 && (best == -1 || min_moves[i][0] < best) {
			fmt.Printf("better solution %d  at position %d", best, i)
			best = min_moves[i][0]
		}
	}

	fmt.Printf("%d\n", best)
}

func main() {
	var target_height, jump_distance, num_itching_sections int
	itching := [SIZE * 2]int{}
	min_moves := board{}
	var itch_start, itch_end int
	fmt.Println("Give ur input")
	fmt.Scanf("%d%d%d", &target_height, &jump_distance, &num_itching_sections)
	for i := 0; i < num_itching_sections; i++ {
		fmt.Scanf("%d%d", &itch_start, &itch_end)
		for j := itch_start; j < itch_end; j++ {
			itching[i] = 1
		}
	}
	fmt.Println("finding distance")
	find_distance(target_height, jump_distance, itching, &min_moves)
	fmt.Println("Solving")
	solve(target_height, min_moves)
}
