package main

import "testing"

func TestSol(t *testing.T) {
	insert("aaa")
	assert(getCount("aa"), 1)
	insert("aa")
	assert(getCount("aa"), 2)
	insert("abb")
	assert(getCount("aa"), 2)
}

func assert(expected, got int) bool {
	return expected == got
}
