package main

import (
	"fmt"
)

const (
	MAX_PEOPLE = 100000
)

type edge struct {
	to_person int
	next      *edge
}

func main() {
	parent := [MAX_PEOPLE + 1]int{} // each index is the edge nubmer, and value is index of the parent
	size := [MAX_PEOPLE + 1]int{}   // arraym, the representative index will hold the number of memebers in that community
	// do not look up using non-representative member

	op := "" // will hold the operation that the terminal tells us to do
	var person1, person2 int

	var num_people, num_comunity, num_ops int

	fmt.Scanf("%d%d", &num_people, &num_comunity)

	for i := 0; i <= num_people; i++ {
		parent[i] = i // initially it will be a parent of itself
		size[i] = 1   // initially each person will be in its own community, thus the size will be 1
	}

	fmt.Scanf("%d", &num_ops)

	for i := 0; i < num_ops; i++ {

		// get the operation we want to do
		fmt.Scanf("%s", &op)

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

func find(person int, parent *[MAX_PEOPLE + 1]int) int {
	community := person

	// loop till the its parent is not equals itself
	// for root node we set its parent to be itself
	for parent[community] != community {
		community = parent[community]
	}

	return community
}

func union(person1, person2 int, parent, size *[MAX_PEOPLE + 1]int, num_comunity int) {
	var community1, community2 int
	community1 = find(person1, parent)
	community2 = find(person2, parent)

	// if the representative of these guys are not the same
	// and when we unionize the size will still stay within the size limit
	if (community1 != community2) && size[community1]+size[community2] <= num_comunity {
		// Setting the root node of communi1 to the root node of community2
		// aka, the preveious representative of community1, which had no representative at all
		// will have the representative of community2 be set as its representative, thus making all its children
		// be represented by that new representative
		parent[community1] = parent[community2]

		// increas the total community size
		size[community2] = size[community2] + size[community1]
	}
}
