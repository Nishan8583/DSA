/*
Adjancy matrix is implemented here, For info on adjancy matrix
https://www.youtube.com/watch?v=9C2cpQZVRBA&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P&index=41
NOTE: Adjancy matarix is not the most efficient in case of space complexity
NOTE: Also this is implemented for unidirected graph
*/

package main

import (
	"fmt"
)

// the structure for adjMatrix
type adjMatrix struct {
	vertexList []string
	matrix     [][]int
}

// a function to return a adjMatrix type, i guess not that useful in simple file at all
func new() adjMatrix {
	return adjMatrix{}
}

// InserVertex adds the new vertex, and expands the matrix accordingly
// since the connection to the new vertex should also be needed to added
func (adj *adjMatrix) InsertVertex(node string) error {

	// if the vertex list is zero
	if len(adj.vertexList) == 0 {
		adj.vertexList = append(adj.vertexList, node)
		adj.matrix = append(adj.matrix, make([]int, 1))
		return nil
	}

	// checking if already present in the vertex list
	for i, v := range adj.vertexList {
		if v == node {
			return fmt.Errorf("Node %s, already present in index %d\n", node, i)
		}
	}

	adj.vertexList = append(adj.vertexList, node)
	adj.matrix = append(adj.matrix, make([]int, len(adj.vertexList)))
	adj.FixMatrix()
	return nil
}

// FixMatrix will get all the values of the matrix, hold it temporarily in a variable
// create a new array as required, and then put all value from previous matrix there
// ooff not so efficient
func (adj *adjMatrix) FixMatrix() {
	for i, v := range adj.matrix {
		temp := v // v will be the slice of int containg the connection with each vertex
		adj.matrix[i] = make([]int, len(adj.vertexList))
		for oldRowsIndex, _ := range temp {
			// updating the new slice with old values
			adj.matrix[i][oldRowsIndex] = temp[oldRowsIndex]
		}
	}
}

// getVertexPosition, gets the index postion for the vertex passed,
// in the vertex list
func (adj *adjMatrix) getVertexPosition(vertex1, vertex2 string) (int, int) {
	vertex1Pos, vertex2Pos := -1, -1
	for i, v := range adj.vertexList {
		if v == vertex1 {
			vertex1Pos = i
		} else if v == vertex2 {
			vertex2Pos = i
		}
	}

	return vertex1Pos, vertex2Pos
}

// AddConnections adds the connection cost to the two vertices
func (adj *adjMatrix) AddConnections(vertex1, vertex2 string, cost int) {
	vertex1Pos, vertex2Pos := adj.getVertexPosition(vertex1, vertex2)
	if vertex1Pos == -1 || vertex2Pos == -1 {
		fmt.Println("Vertex not present")
		return
	}
	adj.matrix[vertex1Pos][vertex2Pos] = cost
	fmt.Println("Vertex Position successfully added")
}

// checkIfConnected checks if two vertices are connected, if they are connected
// their cost is returned, else 0 returned
func (adj *adjMatrix) checkIfConnected(vertex1, vertex2 string) int {
	v1, v2 := adj.getVertexPosition(vertex1, vertex2)
	if v1 == -1 || v2 == -1 {
		return 0
	}
	if adj.matrix[v1][v2] > 0 {
		fmt.Printf("vertex %s and vertex %s are connected \n", vertex1, vertex2)
	}
	return adj.matrix[v1][v2]
}

// FindAllConnectedVertices gets all of the directly connected virtices
func (adj *adjMatrix) findAllConnectedVertices(vertex1 string) []string {

	result := []string{}
	v1, _ := adj.getVertexPosition(vertex1, "")
	if v1 == -1 {
		return result
	}

	// looping through the matrix to find the connection for each
	for i, v := range adj.matrix[v1] {
		if v != 0 {
			result = append(result, adj.vertexList[i])
		}
	}
	return result
}

func (adj *adjMatrix) findShortestPath(vertex1, vertex2 string) {}

func (adj *adjMatrix) findAllPossiblePaths(vertex1, vertex2 string) ([]string, error) {
	v1, v2 := adj.getVertexPosition(vertex1, vertex2)
	if v1 == -1 || v2 == -1 {
		return []string{}, fmt.Errorf("NO connection found")
	}

	possibleConnections := []string{}
	if adj.matrix[v1][v2] > 0 {
		possibleConnections = append(possibleConnections, vertex1)
		possibleConnections = append(possibleConnections, vertex2)
		possibleConnections = append(possibleConnections, "-")
		fmt.Println(possibleConnections)
	}

	for i, v := range adj.matrix[v1] {
		if v != 0 {
			fmt.Println("checking for ", adj.vertexList[i])
			r, err := adj.findAllPossiblePaths(adj.vertexList[i], vertex2)
			if err == nil {
				possibleConnections = append(possibleConnections, r[:len(r)-1]...)
				possibleConnections = append(possibleConnections, "-")
			}
		}
	}
	fmt.Println("Before returning for ", vertex1, vertex2, possibleConnections)
	return possibleConnections, nil
}
func testingInternal() {
	adj := new()
	adj.InsertVertex("A")
	adj.InsertVertex("B")
	adj.InsertVertex("C")
	adj.InsertVertex("D")
	adj.InsertVertex("E")
	adj.InsertVertex("F")
	adj.InsertVertex("G")
	//fmt.Println(adj)
	adj.AddConnections("A", "B", 2)
	adj.AddConnections("A", "D", 2)
	adj.AddConnections("B", "C", 2)
	//adj.AddConnections("D", "B", 2)
	adj.AddConnections("D", "E", 2)
	adj.AddConnections("E", "B", 2)
	fmt.Println(adj)
	fmt.Println(adj.InsertVertex("D"))
	//	fmt.Println(adj.findAllConnectedVertices("A"))
	//	fmt.Println(adj.findAllConnectedVertices("B"))
	fmt.Println(adj.findAllPossiblePaths("A", "B"))
}

func main() {
	testingInternal()
}
