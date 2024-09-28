package main

import (
	"fmt"
	"testing"
)

var zara = "Zara"

func buildTreeForTest() {

	// make sure that map is empty
	nodes = make(map[string]descs)

	inputs := []string{
		//"1",            // Number of test cases, here 1
		//"7 2",          // Number of lines is 7 and distance is 2
		"Lucas 1 Enzo", // Lucas has 1 child, which is Enzo
		"Zara 1 Amber",
		"Sana 2 Gabriel Lucas",
		"Enzo 2 Min Becky",
		"Kevin 2 Jan Cassie",
		"Amber 4 Vlad Sana Ashley Kevin",
		"Vlad 1 Omar",
	}

	for _, v := range inputs {
		createInitNodes(v)
	}
}

func TestBuildTreeSolution(t *testing.T) {
	buildTreeForTest()
	fmt.Println("checking total nodes")
	assert[int](len(nodes), 7, t)

	sanityCheck := func(name string, childrenLen int, childrens []string) {
		fmt.Println("checking children for ", name)
		assert[int](len(nodes[name].children), childrenLen, t)
		copy := make(map[string]int)
		for i, v := range childrens {
			copy[v] = i
		}

		for _, n := range childrens {
			if _, ok := copy[n]; ok {
				delete(copy, n)
			}
		}

		fmt.Println("total children left should be 0")
		assert[int](len(copy), 0, t)

	}

	sanityCheck("Zara", 1, []string{"Amber"})
	sanityCheck("Lucas", 1, []string{"Enzo"})
	sanityCheck("Sana", 2, []string{"Gabriel", "Lucas"})
	sanityCheck("Enzo", 2, []string{"Min", "Becky"})
	sanityCheck("Kevin", 2, []string{"Jan", "Cassie"})
	sanityCheck("Amber", 4, []string{"Vlad", "Sana", "Ashley", "Kevin"})
	sanityCheck("Vlad", 1, []string{"Omar"})

}

func TestScore(t *testing.T) {
	buildTreeForTest()
	// Now test the score
	//assert[int](5, scoreOne(zara, 2), t)
	assert[int](10, scoreOne(zara, 3), t)
}

func assert[K comparable, V int | string | bool](expected, got V, t *testing.T) {
	if expected == got {
		return
	}
	t.Fatalf("Not equal %+v %+v\n", expected, got)
}
