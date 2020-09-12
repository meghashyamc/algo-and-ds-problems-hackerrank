package graphs

import (
	"errors"
	"sort"
)

// Consider an undirected graph consisting of n nodes where each node is labeled from 1 to n and the edge between any two nodes is always of length 6. We define node s to be the starting position for a BFS. Given a graph, determine the distances from the start node to each of its descendants and return the list in node number order, ascending. If a node is disconnected, it's distance should be -1.

// For example, there are n=6 nodes in the graph with a starting node s=1. The list of edges = [[1,2],[2,3],[3,4],[1,5]], and each has a weight of 6.
// https://s3.amazonaws.com/hr-assets/0/1528143002-2e9a521ad9-bfs_shortestExample.png

// Starting from node 1 and creating a list of distances, for nodes 2 through 6 we have distances = [6,2,18,6,-1].
// Function Description

// Define a Graph class with the required methods to return a list of distances.

// more info: https://www.hackerrank.com/challenges/ctci-bfs-shortest-reach

const noconnect = -1
const edgeweight = 6

func findDistancesFromStart(graphNodes int, graphFrom []int, graphTo []int, start int) []int {

	if graphNodes <= 1 {
		return nil
	}

	if len(graphFrom) != len(graphTo) {
		panic("error in input parameters, graphFrom and graphTo must have equal lengths")
	}

	adjacents := make(map[int][]int)
	for i := 0; i < len(graphFrom); i++ {
		addToMapInt(adjacents, graphFrom[i], graphTo[i])
		addToMapInt(adjacents, graphTo[i], graphFrom[i])
	}
	verticesWithDistance := getVerticesWithDistance(adjacents, start)

	//sort slice by vertex ID
	sort.Slice(verticesWithDistance, func(i, j int) bool {
		return verticesWithDistance[i].vertexID < verticesWithDistance[j].vertexID
	})

	return getOnlyDistances(start, graphNodes, verticesWithDistance)
}

func getOnlyDistances(start int, n int, verticesWithDistance []vertex) []int {

	distances := make([]int, n-1)

	i := 0
	var seenstart int
	for vertexID := 0; vertexID < n; vertexID++ {

		if i < len(verticesWithDistance) && vertexID+1 == int(verticesWithDistance[i].vertexID) {
			distances[vertexID-seenstart] = int(verticesWithDistance[i].vertexDepth) * edgeweight
			i++
			continue
		}
		if vertexID+1 == start {
			seenstart = 1
			continue
		}
		distances[vertexID-seenstart] = noconnect
	}
	return distances
}

//breadth first search, storing vertices at each level in a slice
func getVerticesWithDistance(adjacents map[int][]int, start int) []vertex {

	q := queue{}
	verticesWithDistance := make([]vertex, 0)
	visited := make(map[int]bool)
	q.rpush(vertex{vertexID: int32(start), vertexDepth: 0})
	visited[start] = true

	for {
		if len(q) == 0 {
			break
		}

		popped := q.lpop()

		for _, adjacent := range adjacents[int(popped.vertexID)] {
			if !visited[adjacent] {
				visited[adjacent] = true
				currVertex := vertex{vertexID: int32(adjacent), vertexDepth: popped.vertexDepth + 1}
				verticesWithDistance = append(verticesWithDistance, currVertex)
				q.rpush(currVertex)
			}
		}
	}
	return verticesWithDistance
}

func addToMapInt(m map[int][]int, key int, val int) {

	if m == nil {
		panic(errors.New("can't add keys to nil map"))
	}
	if _, ok := m[key]; ok {
		m[key] = append(m[key], val)
		return
	}
	m[key] = []int{val}

}
