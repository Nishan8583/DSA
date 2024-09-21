package main

import "fmt"

const SIZE = 100000

// Check if the right are identical
// we take two snoflake, we return false if any one value on the right are not equal
func identical_right(snow1 [6]int, snow2 [6]int, start int) bool {

	for offset := 0; offset < 6; offset++ {
		if snow1[offset] != snow2[(start+offset)%6] {
			return false
		}
	}
	return true
}

// We return false, if any one value on the left are not equal
func identical_left(snow1 [6]int, snow2 [6]int, start int) bool {
	snow2_index := 0

	for offset := 0; offset < 6; offset++ {
		snow2_index = start - offset
		if snow2_index < 0 {
			snow2_index = snow2_index + 6
		}

		if snow1[offset] != snow2[snow2_index] {
			return false
		}
	}

	return true
}

func are_identical(snow1, snow2 [6]int) bool {

	// we check at each start  position from 0 to 6
	// at each position we check for right and left
	// lets say for 0 for snow1, identify_right checks if 0 position right is same
	// lets say for 1 for snow1, identify_right checks if 1 position right is same
	// and so on, if any one position left or right matches, return true
	for start := 0; start < 6; start++ {
		if identical_right(snow1, snow2, start) {
			return true
		}
		if identical_left(snow1, snow2, start) {
			return true
		}
	}
	return false
}

func identify_identical(snowflakes [SIZE][6]int, n int) {
	for i := 0; i < n; i++ { // loop trough each snow flakes
		for j := i + 1; j < n; j++ { // loop through second snowflakes
			// check if the 2 snowflakes are equal
			if are_identical(snowflakes[i], snowflakes[j]) {
				fmt.Println("Two snoflakes found")
				return
			}
		}
	}
	fmt.Println("no two snowflakes are alike")
}

func main() {
	snowflakes := [SIZE][6]int{}
	n := 0
	fmt.Println("Allocation done, time to execute")
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &snowflakes[i][j])
		}
	}
	identify_identical(snowflakes, n)
}
