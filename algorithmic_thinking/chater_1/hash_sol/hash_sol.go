package main

import "fmt"

const SIZE = 100000

// this will be our value in each index of hashnode
type snowflake_node struct {
	value [6]int // snowflake in current node
	next  *snowflake_node
}

// Hashmap implementation, each index in the array will hold the snowflakes that have common hash code
// THis will help us reduce the space needed to search for similar snowflakes, since we will be searching
// only the ones with same code
type snowflake_hash [SIZE]*snowflake_node

// code gets index for the particular snowflake
// Adding all the numbers in the snowflake and doing module by SIZE, gives number less than SIZE
// This helps reduce collision as compared to just doing SUM code
func code(snowflake [6]int) int {
	sum := 0
	for _, v := range snowflake {
		sum += v
	}
	return sum % SIZE
}

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

func identify_identical(snowflakes snowflake_hash, n int) bool {
	for i := 0; i < n; i++ { // loop trough each snow flakes
		node1 := snowflakes[i]
		for node1 != nil {
			node2 := node1.next
			for node2.next != nil {
				if are_identical(node1.value, node2.value) {
					return true
				}
				node2 = node2.next
			}
			node1 = node1.next
		}

	}

	return false
}

func main() {
	snowflakes := snowflake_hash{}
	n := 0
	fmt.Println("Allocation done, time to execute")
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		snow := snowflake_node{}
		temp_snow := [6]int{}
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &temp_snow[j])
		}
		snow_code := code(temp_snow)      // get the code
		snow.next = snowflakes[snow_code] // make the new current code, point to the first node of the index, this way we do not have to loop through the entire linked list and put it in the end
		snow.value = temp_snow
		snowflakes[snow_code] = &snow // put it in the index
	}
	if identify_identical(snowflakes, n) {
		fmt.Println("Twin snowflakes found.")
	} else {
		fmt.Println("No two snowflakes are alike.")
	}
}
