package main

/*
import "fmt"

const SIZE = 1000000

type (
	board     [SIZE * 2]int // this is our rope
	positions [SIZE * 2]int // array that holds all the positions we have been in
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
		min_moves[i] = -1
	}

	min_moves[0] = 0      // it takes 0 moves to get where we are
	cur_positions[0] = 0  // first we are at 0 length of rope
	num_cur_positions = 1 // since we are at 0 we have 1 positions i.e. 0

	for num_cur_positions > 0 {
		fmt.Println("Number of available positions", num_cur_positions)
		num_new_positions = 0 // reset new positions to 0 for next loop

		for i := 0; i < num_cur_positions; i++ {
			from_height := cur_positions[i] // setting the starting height to be one of the current positions we are in

			// we have checked this position, now jump above
			add_position(from_height, from_height+jump_distance, target_height*2-1,
				&new_positions, &num_new_positions, itching, min_moves)

			// question also states that the guy can jump to any below rope level
			// so we need to do that
			for j := 0; j < from_height; j++ {
				add_position(from_height, j, target_height*2-1,
					&new_positions, &num_new_positions, itching, min_moves)
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

func add_position(
	from_height, to_height, max_height int,
	new_positions *positions,
	num_new_positions *int,
	itching [SIZE * 2]int,
	min_moves *board,
) {
	// if we are within rope limit
	// if we are not in the itching rope section
	// if the current height has not been visited yet
	if (to_height <= max_height) && (itching[to_height] == 0) && min_moves[to_height] == -1 {

		min_moves[to_height] = 1 + min_moves[from_height] // update the min_moves for the rope index, add 1+previous height value
		fmt.Println("Increasing moves value for ", to_height, "to", 1+min_moves[from_height])
		new_positions[*num_new_positions] = to_height // add the position we just visited as a new position to visit in the next loop in the array
		(*num_new_positions)++                        // since we increased the number of new positions we have, increase the counter too
	}
}

func solve(target_height int, min_moves board) {
	best := -1
	fmt.Println("Searching for height", target_height)
	for i := target_height; i < target_height*2; i++ {
		if min_moves[i] != -1 && (best == -1 || min_moves[i] < best) {
			fmt.Printf("better solution %d  at position %d", best, i)
			best = min_moves[i]
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
*/
