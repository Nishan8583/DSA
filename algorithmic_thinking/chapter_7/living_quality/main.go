package main

import (
	"fmt"
)

const (
	MAX_ROWS = 3001
	MAX_COLS = 3001
)

type board [MAX_ROWS][MAX_COLS]int

func can_make_quality(quality, r, c, h, w int, q board) bool {
	zero_one := [MAX_ROWS][MAX_COLS]int{}
	sum := [MAX_ROWS + 1][MAX_COLS + 1]int{}
	var bottom_row, right_col int
	total := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if q[i][j] <= quality {
				zero_one[i][j] = -1
			} else {
				zero_one[i][j] = 1
			}
		}
	}

	for i := 0; i <= c; i++ {
		sum[0][i] = 0
	}
	for i := 0; i <= r; i++ {
		sum[i][0] = 0
	}
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			sum[i][j] = zero_one[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}

	for top_row := 1; top_row < r-h+1; top_row++ {
		for left_col := 1; left_col < c-w+1; left_col++ {
			bottom_row = top_row + h - 1
			right_col = left_col + w - 1
			total = 0
			total = sum[bottom_row][right_col] - sum[top_row-1][right_col] - sum[bottom_row][left_col-1] + sum[top_row-1][left_col-1]
			if total <= 0 {
				return true
			}
		}
	}

	return false
}

func rectangle(r, c, h, w int, q board) int {
	low := 0
	high := r*c + 1
	mid := 0

	for (high - low) > 1 {
		mid = (low + high) / 2
		if can_make_quality(mid, r, c, h, w, q) {
			high = mid
		} else {
			low = mid

		}
	}
	return high
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
