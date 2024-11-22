package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_NODES    = 50000
	LABEL_LENGTH = 16
)

type segtree_node struct {
	left      int // holds the left index
	right     int // holds the right index
	max_index int // holds the max value between those index
}

type treap_node struct {
	label    string
	priority int
}

// init_segtree will just initialize the segment tree without calculating the actual max_index
// node is the index we want to actually populate right now
// left is its left index
// and right is its right index
func init_segtree(segtree *[MAX_NODES*4 + 1]segtree_node, node int, left, right int) {
	mid := 0
	(*segtree)[node].left = left
	(*segtree)[node].right = right

	// we are in the base case, where left and right are the same
	if left == right {
		return
	}

	mid = (left + right) / 2

	// build the left segment tree
	// its left child will be in node*2
	// since we are now trying to get the left half, the right index will be mid
	init_segtree(segtree, node*2, left, mid)

	// build right segment tree
	// so its left index will be mid+2 and right will still be righ
	init_segtree(segtree, node*2+1, mid+1, right)
}

func fill_segtree(
	segtree *[MAX_NODES*4 + 1]segtree_node,
	node int,
	treap_nodes *[MAX_NODES]treap_node,
) int {
	var left_max, right_max int

	// this is the base case. we reach here at the most bottom of the segtree
	// since both left and right are same, we can not be anywhere else but at the bottom
	// so just use its own value as max_value
	if (*segtree)[node].left == (*segtree)[node].right {
		(*segtree)[node].max_index = (*segtree)[node].left
		return (*segtree)[node].max_index
	}

	// if we are not at the bottom, the max of both left segment tree and right segment tree
	left_max = fill_segtree(segtree, node*2, treap_nodes)
	right_max = fill_segtree(segtree, node*2+1, treap_nodes)

	// we get the max value based on the priority
	if (*treap_nodes)[left_max].priority > (*treap_nodes)[right_max].priority {
		(*segtree)[node].max_index = left_max
	} else {
		(*segtree)[node].max_index = right_max
	}

	return (*segtree)[node].max_index
}

func query_segtree(
	segtree *[MAX_NODES*4 + 1]segtree_node,
	node int,
	treap_nodes *[MAX_NODES]treap_node,
	left, right int,
) int {
	var left_max, right_max int

	// basically check if nodes are out of bound
	if (right < (*segtree)[node].left) || (left > (*segtree)[node].right) {
		return -1
	}

	// check if we are withing the range the query wants us to get the value for
	if (left <= (*segtree)[node].left) && (*segtree)[node].right <= right {
		return (*segtree)[node].max_index
	}

	// we do not update left and right, becaus the caller wants to get the value from left and right
	// and we are not filling it
	left_max = query_segtree(segtree, node*2, treap_nodes, left, right)
	right_max = query_segtree(segtree, node*2+1, treap_nodes, left, right)

	// if left search was out of bound, then we found it in right_max
	if left_max == -1 {
		return right_max
	}

	// if right_max was out of bound
	if right_max == -1 {
		return left_max
	}

	if (*treap_nodes)[left_max].priority > (*treap_nodes)[right_max].priority {
		return left_max
	}

	return right_max
}

func splitTerps(value string) (string, int, error) {
	ss := strings.Split(value, "/")
	if len(ss) != 2 {
		return "", -1, fmt.Errorf(
			"invalud string legnth. expected %d got %d with string %s",
			2,
			len(ss),
			value,
		)
	}

	label := ss[0]
	priority, err := strconv.Atoi(ss[1])
	if err != nil {
		return label, -1, fmt.Errorf(
			"could not convert priority into an integer err=%v, priority=%v ",
			err,
			ss[1],
		)
	}

	return label, priority, nil
}

func main() {
	treap_nodes := [MAX_NODES]treap_node{}
	segtree := [MAX_NODES*4 + 1]segtree_node{}
	var num_nodes int

	fmt.Scanf("%d ", &num_nodes)
	for num_nodes > 0 {
		for i := 0; i < num_nodes; i++ {
			labelPr := ""
			fmt.Scanf("%s", &labelPr)
			fmt.Println("read string", labelPr)
			labe, pr, err := splitTerps(labelPr)
			fmt.Println(labe, pr, err)
			if err != nil {
				fmt.Println("Error while parsing terp", err)
			}
			treap_nodes[i].label = labe
			treap_nodes[i].priority = pr
			// treap_nodes[i].label
		}

		init_segtree(&segtree, 1, 0, num_nodes-1)
		fill_segtree(&segtree, 0, &treap_nodes)
		solve(&treap_nodes, 0, num_nodes-1, &segtree)
		fmt.Printf("\n")
		fmt.Scanf("%d", *&num_nodes)
	}
}

func solve(
	treap_nodes *[MAX_NODES]treap_node,
	left, right int,
	segtree *[MAX_NODES*4 + 1]segtree_node,
) {
	var root_index int
	var root treap_node

	if left > right {
		return
	}

	root_index = query_segtree(segtree, 1, treap_nodes, left, right)
	root = (*treap_nodes)[root_index]
	fmt.Printf("(")
	solve(treap_nodes, left, root_index-1, segtree)
	fmt.Printf("%s/%d", root.label, root.priority)
	solve(treap_nodes, root_index+1, right, segtree)
	fmt.Printf(")")
}
