package main

import "fmt"

func printSpiralOrder(a [][]int, r, c int) {
	top, left := 0, 0
	bottom := r - 1
	right := c - 1

	// the direction to go to
	// 0 -> left to right
	// 1 -> top to bottom
	// 2 -> right to left
	// 3 -> bottom to up
	direction := 0
	for top <= bottom && left <= right {
		switch direction {
		case 0:
			fmt.Println("left to right")
			for i := left; i <= right; i++ {
				fmt.Println(a[top][i])
			}
			top = top + 1
			direction = 1
		case 1:
			fmt.Println("top to bottom")
			for i := top; i <= bottom; i++ {
				fmt.Println(a[i][right])
			}
			direction = 2
			right = right - 1
		case 2:
			fmt.Println("right to left")
			for i := right; i >= left; i-- {
				fmt.Println(a[bottom][i])
			}
			direction = 3
			bottom = bottom - 1
		case 3:
			fmt.Println("bottom to up")
			for i := bottom; i >= top; i-- {
				fmt.Println(a[i][left])
			}
			direction = 0
			left = left + 1
		}

	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	printSpiralOrder(matrix, 4, 4)
}
