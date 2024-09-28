package main

import (
	"fmt"
	"strings"
)

type descs struct {
	children map[string]descs
}

var nodes = make(map[string]descs)

func createInitNodes(v string) {
	names := strings.Split(v, " ")
	if len(names) < 3 {
		panic("Invalid string for names")
	}

	if _, ok := nodes[names[0]]; ok {
		for _, v := range names[2:] {
			nodes[names[0]].children[v] = descs{}
		}
	} else {
		nodes[names[0]] = descs{
			children: map[string]descs{},
		}
		for _, v := range names[2:] {
			nodes[names[0]].children[v] = descs{}
		}
	}

	return
}

func scoreOne(name string, d int) int {

	fmt.Printf("Depth for %s is %d \n", name, d)
	if d == 1 {
		c := len(nodes[name].children)
		fmt.Printf("number of children for %s is %d\n", name, c)
		return c
	}

	total := len(nodes[name].children)

	// if it has no child right now, just check the main node and see if it has some there

	for nameChild, _ := range nodes[name].children {
		fmt.Printf("Searching inside for Parent: %s child: %s\n", name, nameChild)
		total += scoreOne(nameChild, d-1)
	}

	fmt.Printf("Retruning %d for %s\n", total, name)
	return total
}

func score_all(nodes map[string]descs, depth int) []int {
	scores := make([]int, 0, len(nodes))
	for name, _ := range nodes {
		sc := scoreOne(name, depth)
		scores = append(scores, sc)
	}
	return scores
}
func checkIfHasParent(name string, node *descs) (*descs, bool) {

	// if no childern,then checking is pointless
	if len(node.children) == 0 {
		return nil, false
	}

	// if the name of the boi is found in children, then return this current node as parent
	if _, ok := node.children[name]; ok {
		return node, ok
	}

	// loop through all the chidren and check if any children is parent of the given node
	for _, child := range node.children {
		v, ok := checkIfHasParent(name, &child)
		if ok {
			return v, ok
		}
	}

	return nil, false

}
