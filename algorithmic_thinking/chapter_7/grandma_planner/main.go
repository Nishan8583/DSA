package main

import (
	"fmt"
)

const (
	MAX_TOWNS = 700
)

type edge struct {
	to_town int
	length  int
	next    *edge
}

func main() {
	/*
		Input sample:
							4				// number of towns we have
							0 3 8 2 // the cost of reahing from town 1 to other towns, each position gives town number starting town 1 2 3 4, 1->1 costs 0, 1->2 costs 3, 1->3 costs 8 and so on
							3 0 2 1
							8 2 0 5
							2 1 5 0
							1				// number of towns with cookies
							2				// 2 has a store with cookie
	*/
	adj_list := [MAX_TOWNS + 1]*edge{} // holds the adjency list
	var num_towns, num_stores int      // total num_towns provided and num_stores holds the number of stores with cookie
	stores := [MAX_TOWNS + 1]int{}     // array of stores, 1 means it has cookie, 0 means it does not

	fmt.Printf("Please enter number of town: ")
	fmt.Scanf("%d", &num_towns)

	fmt.Printf(
		"From town each town costs (give comma separated cost, for each town represented by index position): \n",
	)
	for from_town := 1; from_town <= num_towns; from_town++ { // each line represents from_town
		fmt.Printf("from_town %d: ", from_town)
		for to_town := 1; to_town <= num_towns; to_town++ { // each index in the line represents to_town
			length := 0
			fmt.Scanf(
				"%d",
				&length,
			) // each value represents length/cost to get from_town to to_town
			e := edge{
				to_town: to_town,
				length:  length,
			}
			e.next = adj_list[from_town]
			adj_list[from_town] = &e
		}
	}

	fmt.Printf("Please enter number of stores with cookies: ")
	fmt.Scanf("%d", &num_stores)
	for i := 1; i <= num_stores; i++ {
		store := 0
		fmt.Printf("store: ")
		fmt.Scanf("%d", &store)
		stores[store] = 1
	}

	solve(adj_list, num_towns, stores)
}

func solve(adj_list [MAX_TOWNS + 1]*edge, num_towns int, stores [MAX_TOWNS + 1]int) {
	// we keep track of shortest path for both state, one without cookie (0) and one with cookie (1)
	done := [MAX_TOWNS + 1][2]bool{} // 1st dimenion holds info about wheter its done or not, second holds info about the state (0,1) whehter we have the cookie or not, each index represents town
	min_distances := [MAX_TOWNS + 1][2]int{}

	num_paths := [MAX_TOWNS + 1][2]int{} // will holds the number of shortest path to the town, for both state, 0 i.e no cookie state, and 1 i.e with cookie state
	var found bool

	var min_distance, min_town_index, min_state_index, old_distance int
	var e *edge

	for state := 0; state <= 1; state++ {
		for i := 1; i <= num_towns; i++ {
			done[i][state] = false
			min_distances[i][state] = -1
			num_paths[i][state] = 0
		}
	}

	min_distances[1][0] = 0 // min cost to reach to itself without cookie (i.e. state 0) is 0
	num_paths[1][0] = 1     // number of shortest path to reach town 1 with no cookie is always 1, since it can always reach to itself with no cookie

	// we loop twice the number of towns because we want to account for both state of having a cookie and not having a cookie
	for i := 0; i <= num_towns*2; i++ {
		min_distance = -1
		found = false

		// we loop twice for the state of 0 i..e with no cookie, and 1 with cookie
		// we keep shortest distance information for both state, one with cookie and one without
		for state := 0; state <= 1; state++ {
			for j := 1; j <= num_towns; j++ {
				if !done[j][state] &&
					min_distances[j][state] >= 0 { // check if the state is not done and minimum distance is gte 0, will always hit true in the first case cause we set min_distances[1][0] to true
					if min_distance == -1 ||
						min_distances[j][state] < min_distance { // check if min distance is -1 (value not found yet) or this new min distance we found is less than previous min distance we found
						min_distance = min_distances[j][state]
						min_state_index = state
						min_town_index = j
						found = true
						fmt.Printf("Setting town = %d state %v to done\n", j, state)
					}
				}
			}

			if !found {
				break
			}

			done[min_town_index][min_state_index] = true

			if min_state_index == 0 &&
				stores[min_town_index] == 1 { // we are in a state where we do not have the cookie, but we just entereed a town with cookie i.e. the current index which has the new lowest cost is the one with cookie
				old_distance = min_distances[min_town_index][1]
				if (old_distance == -1) ||
					(old_distance >= min_distance) { // we do not need cost to jump from a node that has cookie to itself, its just visual, we are changing state from not having cookie to the one where we are having cookie
					min_distances[min_town_index][1] = min_distance

					// if another short path with same value was found, then increase the num_paths count for that path

					old_count := num_paths[min_town_index][1]
					if old_distance == min_distance {
						fmt.Println(
							"Older 0 index value, found equivalent value ",
							num_paths[min_town_index][0],
							old_distance,
							min_distance,
						)
						num_paths[min_town_index][1] += num_paths[min_town_index][0]
						fmt.Printf(
							"Stage change Increaseing older shortest path for %d older count = %d to  %d\n",
							min_town_index,
							old_count,
							num_paths[min_town_index][1],
						)
					} else {
						// if more shorter path is found, then reset the shortest path count index
						fmt.Println("Found new shortest distance", old_distance, min_distance)
						num_paths[min_town_index][1] = num_paths[min_town_index][0]

						fmt.Printf(
							"Stage change reset shortest path for %d older count = %d to  %d\n",
							min_town_index, old_count,
							num_paths[min_town_index][1],
						)
					}
					num_paths[min_town_index][1] %= 1000000 // we want to keep the number of shortest path count below 1000000
				}
			} else {
				e = adj_list[min_town_index]
				for e != nil {
					old_distance = min_distances[e.to_town][min_state_index]
					if (old_distance == -1) ||
						old_distance >= min_distance+e.length {
						min_distances[e.to_town][min_state_index] = min_distance + e.length

						// if same shorter path was found again, increase the count of that path
						older_count := num_paths[e.to_town][min_state_index]
						if old_distance == min_distance+e.length {
							fmt.Println("Found equivalent values", min_distance, min_distance+e.length)
							num_paths[e.to_town][min_state_index] += num_paths[min_town_index][min_state_index]

							fmt.Printf(
								"Stage change Increaseing older shortest path for %d  older_count = %d tto %d\n",
								min_town_index, older_count,
								num_paths[min_town_index][min_state_index],
							)
						} else {

							fmt.Println("Found new shorter values", min_distance, min_distance+e.length)
							// else reset the short path count
							num_paths[e.to_town][min_state_index] = num_paths[min_town_index][min_state_index]

							fmt.Printf(
								"Stage change reset shortest path for %d to older_count=%d to  %d\n",
								min_town_index, older_count,
								num_paths[min_town_index][min_state_index],
							)
						}
						num_paths[e.to_town][min_state_index] %= 1000000
					}
					e = e.next
				}
			}
		}
	}

	fmt.Printf(
		"The shortest distance=%d shortest distance count=%d\n",
		min_distances[num_towns][1],
		num_paths[num_towns][1],
	) // print the shortest distance to the destination town, with state where we have the cookie
}
