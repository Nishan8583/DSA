/*
We do BFS
We check each possible moves for knight, and increase the moves count until we reach the row and col the pawn is
if we can not reach, check for stalemeate case
*/

package main

import "fmt"

const (
	MAX_ROWS = 99
	MAX_COLS = 99
)

type (
	board    [MAX_ROWS + 1][MAX_COLS + 1]int // The chess board representatiob
	position struct {                        // A particular position
		row int
		col int
	}
	positions [MAX_ROWS * MAX_COLS]position // array of all position
)

func find_distance(knight_row, knight_col, destination_row, destination_col,
	num_rows, num_cols int) int {
	cur_pos := positions{}
	new_pos := positions{}

	num_cur_pos := 0
	num_new_pos := 0

	min_moves := board{}

	// setting to default of -1
	for i := 1; i <= num_rows; i++ {
		for j := 1; j <= num_cols; j++ {
			min_moves[i][j] = -1
		}
	}
	min_moves[knight_row][knight_col] = 0
	cur_pos[0] = position{knight_row, knight_col} // In the beginning the current position is provided by the user
	num_cur_pos = 1                               // since we are in a position, number of current position will always be 1 here

	for num_cur_pos > 0 { // loop till we keep on finding new positions to check
		num_new_pos = 0                    // reset it to 0, we will update it if we find new positions to check
		for i := 0; i < num_cur_pos; i++ { // loop through all possible positions for knight

			// we will be starting from the current position, so get the current row and column
			from_row := cur_pos[i].row
			from_col := cur_pos[i].col

			// knight has reached the pawns position
			if from_row == destination_row && from_col == destination_col {
				fmt.Println("got into the same palce as pawn in moves", from_row, from_col, min_moves[destination_row][destination_col])
				return min_moves[destination_row][destination_col]
			}

			// BFS.
			// Do all the possible 1 level moves that the knight can do
			// move knight to up right L
			add_position(from_row, from_col, from_row+1, from_col+2, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)

			// move knight to up left
			add_position(from_row, from_col, from_row+1, from_col-2, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)

			// down right
			add_position(from_row, from_col, from_row-1, from_col+2, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)

			// down left
			add_position(from_row, from_col, from_row-1, from_col-2, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)

			// Do for all Ls that are flat
			add_position(from_row, from_col, from_row+2, from_col+1, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)
			add_position(from_row, from_col, from_row+2, from_col-1, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)
			add_position(from_row, from_col, from_row-2, from_col+1, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)
			add_position(from_row, from_col, from_row+2, from_col-1, num_rows, num_cols, &new_pos, &num_new_pos, &min_moves)

			num_cur_pos = num_new_pos
			for i := 0; i < num_cur_pos; i++ {
				cur_pos[i] = new_pos[i]
			}
		}
	}

	return -1
}

func add_position(from_row, from_col, to_row, to_col, num_rows, num_cols int, new_positions *positions, num_new_pos *int, min_moves *board) {

	new_pos := position{}

	if to_row >= 1 && to_col >= 1 && to_row <= num_rows && to_col <= num_cols && // check if valid row and column
		min_moves[to_row][to_col] == -1 { // check if the current row/col combo has not been checked yet

		min_moves[to_row][to_col] = 1 + min_moves[from_row][from_col] // add 1 move more, because we moved from from_row/from_col to to_row and to_col
		new_pos = position{row: to_row, col: to_col}
		(*new_positions)[*num_new_pos] = new_pos
		*num_new_pos++
	}
}

func solve(pawn_row, pawn_col, knight_row, knight_col, num_rows, num_col int) {
	fmt.Printf("Pawn row = %d pawn col = %d, knight_row = %d knight_col = %d num_rows = %d num_col = %d\n", pawn_row, pawn_col, knight_row, knight_col, num_rows, num_col)
	cur_pawn_row, num_moves := 0, 0

	cur_pawn_row = pawn_row
	num_moves = 0

	// check if knight can win
	for cur_pawn_row < num_rows { // loop while pawn has not reached the top of the row
		knight_takes := find_distance(knight_row, knight_col, cur_pawn_row, pawn_col, num_rows, num_col)
		if knight_takes >= 0 && num_moves >= knight_takes && (num_moves-knight_takes)%2 == 0 {
			fmt.Printf("Win in %d knight move(s).\n", num_moves)
			return
		}
		cur_pawn_row++
		num_moves++
	}

	cur_pawn_row = pawn_row
	num_moves = 0
	// checking if stalemate is possible
	for cur_pawn_row < num_rows {
		knight_takes := find_distance(knight_row, knight_col, cur_pawn_row+1, pawn_col, num_rows, num_col)
		if knight_takes >= 0 && num_moves >= knight_takes && (num_moves-knight_takes)%2 == 0 {
			fmt.Printf("Stalemate in %d knight move(s).\n", num_moves)
			return
		}
		cur_pawn_row++
		num_moves++
	}

	// If neither stalemate or win, then its a loss

	fmt.Printf("Loss in %d knight move(s).\n", num_rows-pawn_row-1)
}

func main() {
	num_cases := 0
	var num_rows, num_cols, pawn_row, pawn_col, knight_row, knight_col int
	fmt.Println("Started")
	fmt.Scanf("%d", &num_cases)
	for i := 0; i < num_cases; i++ {
		fmt.Scanf("%d", &num_rows)
		fmt.Scanf("%d", &num_cols)
		fmt.Scanf("%d", &pawn_row)
		fmt.Scanf("%d", &pawn_col)
		fmt.Scanf("%d", &knight_row)
		fmt.Scanf("%d", &knight_col)
		fmt.Println("Solving")
		solve(pawn_row, pawn_col, knight_row, knight_col, num_rows, num_cols)
		fmt.Println("Finished sovling")
	}
}
