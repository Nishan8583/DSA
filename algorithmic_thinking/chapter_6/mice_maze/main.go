package main

import "fmt"

const (
	MAX_CELLS = 100
)

type edge struct {
	to_cell int
	length  int
	next    *edge
}

func main() {
	adj_list := [MAX_CELLS + 1]*edge{}
	var num_cases int

	var num_cells, exit_cell, time_limit, num_edges int

	fmt.Println("Number of cases: ")
	fmt.Scanf("%d", &num_cases)
	for case_num := 1; case_num <= num_cases; case_num++ {
		fmt.Println("BLANK LINE:")
		fmt.Scanf("%s") // blank line
		fmt.Println("NUM CELLS")
		fmt.Scanf("%d", &num_cells)
		fmt.Println("EXIT CELL")
		fmt.Scanf("%d", &exit_cell)
		fmt.Println("TIME_LIMIT")
		fmt.Scanf("%d", &time_limit)
		fmt.Println("EDGES")
		fmt.Scanf("%d", &num_edges)

		// resetting the arrays to nil
		for i := 1; i < num_cells; i++ {
			adj_list[i] = nil
		}

		for i := 0; i < num_edges; i++ {
			var from_cell, to_cell, length int
			fmt.Println("EDGE", i)
			fmt.Scanf("%d%d%d", &from_cell, &to_cell, &length)
			fmt.Printf(
				"Input was from_cell=%d to_cell=%d legnth=%d\n",
				from_cell,
				to_cell,
				length,
			)
			e := edge{
				to_cell: to_cell,
				length:  length,
				next:    adj_list[from_cell], // point to the beginning of the edge so taht we can just append it to the beginning instead. lets say a[1] = e1,
				// new edge e2, will point to e1, and the new a[1] = e2,
				// e1 is now reachable via  e2
			}
			adj_list[from_cell] = &e
		}

		total := 0
		for i := 0; i <= num_cells; i++ {
			min_time := find_time(adj_list, num_cells, i, exit_cell)
			if min_time >= 0 && min_time <= time_limit {
				total++
			}
		}
		fmt.Printf("%d\n", total)
		if case_num < num_cases {
			fmt.Printf("\n")
		}
	}
}

func find_time(adj_list [MAX_CELLS + 1]*edge, num_cells, from_cell, exit_cell int) int {
	done := [MAX_CELLS + 1]bool{}     // holds information of whether a cell is done or not
	min_times := [MAX_CELLS + 1]int{} // holds information about minimum time to get to that cell

	var min_time, min_time_index, old_time int // min_time_index, the index of cell whose shortest path has been found, and min_time is the time found
	var found bool                             //

	var e *edge
	for i := 1; i <= num_cells; i++ {
		done[i] = false
		min_times[i] = -1
	}
	min_times[from_cell] = 0

	// now we do a loop, we want to check the shortest time it takes to get to all cells from from_cell
	for i := 0; i < num_cells; i++ {
		min_time = -1
		found = false

		for j := 1; j <= num_cells; j++ {
			// The cell has to be not done and its min_times has to be atleast greater than or equals to 0
			if !done[j] && min_times[j] >= 0 {
				// if min_time has not be set yet or we found some new time which is less than  previous min time
				if min_time == -1 || min_times[j] < min_time {
					min_time = min_times[j]
					min_time_index = j
					found = true
				}
			}
		}
		if !found {
			break
		}
		done[min_time_index] = true

		e = adj_list[min_time_index]
		for e != nil {
			old_time = min_times[e.to_cell]
			if old_time == -1 || old_time > min_time+e.length {
				min_times[e.to_cell] = min_time + e.length
			}
			e = e.next
		}
	}
	fmt.Printf("Min time for %d is %d\n", from_cell, min_times[exit_cell])
	fmt.Printf("Final distance calculated using dijkstras algorithm\n")
	fmt.Println("From Node: ", from_cell)
	fmt.Printf("To_node		Cost\n")
	for i := 1; i < num_cells; i++ {
		fmt.Printf("%d		%d\n", i, min_times[i])
	}
	return min_times[exit_cell]
}
