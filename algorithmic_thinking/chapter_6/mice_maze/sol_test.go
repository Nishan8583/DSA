package main

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	edges := [][3]int{
		{1, 2, 12},
		{1, 3, 6},
		{2, 1, 26},
		{1, 4, 45},
		{1, 5, 7},
		{3, 2, 2},
		{2, 4, 9},
		{4, 3, 8},
		{5, 2, 21},
	}

	min_time := 0
	adj_list := [MAX_CELLS + 1]*edge{}
	for i := 0; i < len(edges); i++ {
		from_cell := edges[i][0]
		to_cell := edges[i][1]
		length := edges[i][2]
		fmt.Println("ADDING", edges[i])
		e := edge{
			to_cell: to_cell,
			length:  length,
			next:    adj_list[from_cell], // point to the beginning of the edge so taht we can just append it to the beginning instead. lets say a[1] = e1,
			// new edge e2, will point to e1, and the new a[1] = e2,
			// e1 is now reachable via  e2
		}
		adj_list[from_cell] = &e
	}

	fmt.Println("adjency list", adj_list)

	total := 0
	num_cells := 5
	exit_cell := 4
	time_limit := 12
	for i := 0; i <= num_cells; i++ {
		min_time = find_time(adj_list, num_cells, i, exit_cell)
		if min_time >= 0 && min_time <= time_limit {
			total++
		}
	}
	fmt.Printf("Total mice that got out %d\n", total)
}
