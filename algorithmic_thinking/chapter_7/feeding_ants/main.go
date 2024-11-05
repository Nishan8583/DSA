package main

import "fmt"

const (
	MAX_NODES = 1000
	HIGHEST   = 2000000000
)

// since edge holds information about the superpipe feature, the author decided to use edge as the structure
type edge struct {
	to_node, percentage int
	superpipe           bool // hold info about wheter the current edge has superpipe feature enabled or not
	next                *edge
}

func main() {
	/*
		adg_list := [MAX_NODES + 1]*edge{}
		liquids_needed := [MAX_NODES + 1]int{}

		num_nodes := 0
		fmt.Printf("Enter number of nodes: ")
		fmt.Scanf("%d", &num_nodes)

		var from_node, to_node, percentage, superpipe int

		for i := 0; i < num_nodes-1; i++ {
			fmt.Printf("Enter for %d :", i)
			fmt.Scanf("%d", &from_node)
			fmt.Scanf("%d", &to_node)
			fmt.Scanf("%d", &percentage)
			fmt.Scanf("%d", &superpipe)
			e := &edge{}
			e.to_node = to_node
			e.percentage = percentage
			if superpipe == 1 {
				e.superpipe = true
			} else {
				e.superpipe = false
			}
			e.next = adg_list[from_node]
			adg_list[from_node] = e
		}

		for i := 1; i <= num_nodes; i++ {
			fmt.Printf("Enter for %d :", i)
			fmt.Scanf("%d", &liquids_needed[i])
		}
	*/
	e_list := [MAX_NODES + 1]*edge{}
	edge_slice := [][]int{
		{1, 2, 20, 0},
		{1, 3, 50, 0},
		{1, 4, 30, 1},
		{4, 5, 50, 0},
		{4, 6, 50, 0},
	}
	for _, v := range edge_slice {
		e := &edge{}
		from_node := v[0]
		to_node := v[1]
		percentage := v[2]
		s := false
		if v[3] == 1 {
			s = true
		}
		e.to_node = to_node
		e.percentage = percentage
		e.superpipe = s
		e.next = e_list[from_node]
		e_list[from_node] = e

	}

	liquids_needed := [MAX_NODES + 1]int{0, -1, 2, 9, -1, 7, 8}

	solve(e_list, liquids_needed)
}

func solve(edge_list [MAX_NODES + 1]*edge, liquids_needed [MAX_NODES + 1]int) {
	var low, mid, high float64
	low = 0.0
	high = HIGHEST
	for (high - low) > 0.00001 {
		mid = (low + high) / 2
		fmt.Println("Trying", mid)
		if can_feed(1, mid, edge_list, liquids_needed) {
			fmt.Println("POssible with ", mid)
			high = mid // if we can feed all ants with the water, it means we are in the feasable range, so set mid as the new high
		} else {
			fmt.Println("Not POssible with", mid)
			low = mid // if not, then we have to increase the number of liter, we do that by increasing the low point
		}
	}
	fmt.Println(high)
}

func can_feed(
	node int,
	liquid float64,
	edge_list [MAX_NODES + 1]*edge,
	liquids_needed [MAX_NODES + 1]int,
) bool {
	// which means it requires some liquid, i.e. its a leaf node
	if liquids_needed[node] != -1 {
		fmt.Printf("liquid present %v liquid required = %v\n", liquid, liquids_needed[node])
		return liquid >= float64(liquids_needed[node])
	}
	fmt.Println(node)
	e := edge_list[node]
	ok := true

	for e != nil && ok {
		down_pipe := (liquid * float64(e.percentage)) / 100
		if e.superpipe {
			down_pipe = down_pipe * down_pipe
		}
		if !can_feed(e.to_node, down_pipe, edge_list, liquids_needed) {
			ok = false
		}
		e = e.next
	}

	return ok
}
