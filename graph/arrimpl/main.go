package main

import "errors"

type ArrayImpl struct {
	VertexList []string
	EdgeList   [][]EdgeInfo
}

type EdgeInfo struct {
	Connection string
	Cost       int
}

func New() ArrayImpl {
	return ArrayImpl{}
}

// Insert inserts the node node into the vertext list
// and prepares edgeList accordingly
func (a *ArrayImpl) Insert(node string) (int, error) {

	// if the slice is already empty, then no need to check for values
	if len(a.VertexList) == 0 {
		return a.putValues(node), nil
	}

	// check if the node was already present
	for _, v := range a.VertexList {
		if v == node {
			return 0, errors.New(node + " already on the vertex list")
		}
	}

	return a.putValues(node), nil
}

// putValues puts the values in the node, and returns the position for it
func (a *ArrayImpl) putValues(node string) int {
	newPos := len(a.VertexList)
	a.VertexList = append(a.VertexList, node)
	a.EdgeList[newPos] = []EdgeInfo{}
	return newPos
}

// InsertEdge inserts the node in the edgeInfo list
func (a *ArrayImpl) InsertEdge(srcNode string, dstNode string, cost int) {

	for i, v := range a.VertexList {
		if v == srcNode {
			a.EdgeList[i] = append(a.EdgeList[i], EdgeInfo{Connection: dstNode, Cost: cost})
		}
	}
}
