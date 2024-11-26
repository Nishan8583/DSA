package main

import (
	"fmt"
	"testing"
)

type testCaseStruct struct {
	Op      string
	Person1 int
	Person2 int
}

func TestSol(t *testing.T) {
	parent := [MAX_PEOPLE + 1]int{} // each index is the edge nubmer, and value is index of the parent
	size := [MAX_PEOPLE + 1]int{}   // arraym, the representative index will hold the number of memebers in that community
	// do not look up using non-representative member

	num_people := 7
	num_comunity := 6
	for i := 0; i <= num_people; i++ {
		parent[i] = i // initially it will be a parent of itself
		size[i] = 1   // initially each person will be in its own community, thus the size will be 1
	}

	test_cases := []testCaseStruct{
		{"A", 1, 4},
		{"A", 4, 5},
		{"A", 3, 6},
		{"E", 1, 5},
		{"E", 2, 5},
		{"A", 1, 5},
		{"A", 2, 5},
		{"A", 4, 3},
		{"s", 4, -1},
		{"A", 7, 6},
		{"S", 4, -1},
	}

	for _, test_case := range test_cases {
		op := test_case.Op
		person1 := test_case.Person1
		person2 := test_case.Person2
		// get the operation we want to do
		// if its an Add operation
		if op == "A" {
			fmt.Scanf("%d%d", &person1, &person2)
			union(person1, person2, &parent, &size, num_comunity)
		} else if op == "E" { // Examine opretaion
			fmt.Scanf("%d,%d", &person1, &person2)

			// if both have the same representative then, yes they are in the same community
			if find(person1, &parent) == find(person2, &parent) {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
		} else {
			fmt.Scanf("%d", &person1)
			fmt.Printf("%d\n", size[find(person1, &parent)])
		}
	}
}
