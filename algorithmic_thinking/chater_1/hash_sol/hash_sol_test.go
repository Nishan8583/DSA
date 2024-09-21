package main

import (
	"fmt"
	"testing"
)

func helpTest(test_cases [][6]int) bool {
	snowflakes := snowflake_hash{}
	fmt.Println("Allocation done, time to execute")
	for _, v := range test_cases {
		snow := snowflake_node{}

		snow_code := code(v)              // get the code
		snow.next = snowflakes[snow_code] // make the new current code, point to the first node of the index, this way we do not have to loop through the entire linked list and put it in the end
		snow.value = v
		snowflakes[snow_code] = &snow // put it in the index
	}

	return identify_identical(snowflakes, len(test_cases))
}

func TestAreIdentical(t *testing.T) {
	testCases := [][6]int{
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1},
	}

	if !helpTest(testCases) {
		t.Fail()
	}

	testCases = [][6]int{
		{1, 2, 3, 4, 5, 7},
		{6, 5, 4, 3, 2, 1},
	}
	if helpTest(testCases) {
		t.Fail()
	}
}
