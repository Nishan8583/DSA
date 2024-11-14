package main

/*import (
	"fmt"
	"sort"
)

const (
	MAX_ROWS = 3001
	MAX_COLS = 3001
)

type board [MAX_ROWS][MAX_COLS]int

func median(top_row, left_col, bottom_row, right_col int, q board) int {
	cur_rectangle := []int{}
	num_cur_rectangle := 0

	for i := top_row; i <= bottom_row; i++ {
		for j := left_col; j <= right_col; j++ {
			fmt.Println("Appending ", q[i][j])
			cur_rectangle = append(cur_rectangle, q[i][j])
			num_cur_rectangle++
		}
	}

	//fmt.Println(cur_rectangle)
	sort.Ints(cur_rectangle)
	return cur_rectangle[num_cur_rectangle/2]
}

func rectangle(r, c, h, w int, q board) int {
	best := r*c + 1
	result := 0

	for top_row := 0; top_row < r-h+1; top_row++ {
		for left_col := 0; left_col < c-w+1; left_col++ {
			bottom_row := top_row + h - 1
			right_col := left_col + w - 1

			result = median(top_row, left_col, bottom_row, right_col, q)
			fmt.Println("result", result)
			if result < best {
				best = result
			}
		}
	}

	return best
}

func main() {
	q := board{
		{48, 16, 15, 45, 40, 28, 8},
		{20, 11, 36, 19, 24, 6, 33},
		{22, 39, 30, 7, 9, 1, 18},
		{14, 35, 2, 13, 31, 12, 46},
		{32, 37, 21, 3, 41, 23, 29},
		{42, 49, 38, 10, 17, 47, 5},
		{43, 4, 34, 25, 26, 27, 44},
	}

	result := rectangle(7, 7, 5, 3, q)
	fmt.Println("Result", result)
}
*/
