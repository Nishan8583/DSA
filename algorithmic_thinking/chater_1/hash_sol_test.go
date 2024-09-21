package main

import (
	"fmt"
	"testing"
)

func TestAreIdentical(t *testing.T) {
	testCases := [][6]int{
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1},
	}

	if !helpTest(testCases) {
		t.Fail()
	}

	fmt.Println("test case pass for identical")

	testCases = [][6]int{
		{1, 2, 3, 4, 5, 7},
		{6, 5, 4, 3, 2, 1},
	}
	if helpTest(testCases) {
		t.Fail()
	}

	fmt.Println("")
}
