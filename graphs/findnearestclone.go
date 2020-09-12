package graphs

import (
	"errors"
)

// In this challenge, there is a connected undirected graph where each of the nodes is a color. Given a color, find the shortest path connecting any two nodes of that color. Each edge has a weight of 1. If there is not a pair or if the color is not found, print -1.

// For example, given graph_nodes =5, and 4 edges  and g_from = [1,2,2,3] and g_to = [2,3,4,5] colors for each node are arr = [1,2,3,1,3] we can draw the following graph:
// https://s3.amazonaws.com/hr-assets/0/1529952915-a96eba7baa-nearestcloneexample.png

// Each of the nodes is labeled [node]/[color] and is colored appropriately. If we want the shortest path between color 3, blue, we see there is a direct path between nodes 3 and 5. For green, color , 1 we see the path length 2 from 1->2->4. There is no pair for node 4 having color 2, red.

// Function Description

// Complete the findShortest function in the editor below. It should return an integer representing the length of the shortest path between two nodes of the same color, or -1 if it is not possible.

// findShortest has the following parameter(s):

// g_nodes: an integer, the number of nodes
// g_from: an array of integers, the start nodes for each edge
// g_to: an array of integers, the end nodes for each edge
// ids: an array of integers, the color id per node
// val: an integer, the id of the color to match

// https://www.hackerrank.com/challenges/find-the-nearest-clone

/*
 * For the unweighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] to <name>To[i].
 *
 */

type vertex struct {
	starterID   int32
	vertexID    int32
	vertexDepth int32
}
type queue []vertex

func (q *queue) rpush(i vertex) {
	if q == nil {
		q = &queue{}
	}
	*q = append(*q, i)
}

func (q *queue) lpop() vertex {
	if q == nil || len(*q) == 0 {
		panic(errors.New("can't pop from empty queue"))
	}
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped

}
func (q *queue) lpeek() vertex {
	if q == nil || len(*q) == 0 {
		panic(errors.New("can't peek into empty queue"))
	}
	return (*q)[0]
}
func findShortest(graphNodes int32, graphFrom []int32, graphTo []int32, colors []int64, reqColor int64) int32 {

	if graphNodes == 0 {
		return 0
	}
	if len(colors) != int(graphNodes) || len(graphFrom) != len(graphTo) {
		panic(errors.New("error in parameters"))
	}

	adjacentsList := make(map[int32][]int32)

	for i := 0; i < len(graphFrom); i++ {
		addToMap(adjacentsList, graphFrom[i], graphTo[i])
		addToMap(adjacentsList, graphTo[i], graphFrom[i])
	}

	var starterIDs []int32

	for i, color := range colors {
		if color == int64(reqColor) {
			starterIDs = append(starterIDs, int32(i)+1)
		}
	}
	//colour not found or there is no pair
	if len(starterIDs) <= 1 {
		return -1
	}

	return getMinDistanceForColor(adjacentsList, colors, starterIDs, reqColor)

}

//breadth first search starting from each vertex of the given colour
func getMinDistanceForColor(adjacentsList map[int32][]int32, colors []int64, starterIDs []int32, reqColor int64) int32 {

	q := queue{}

	//map of vertexID and set of starter ID traversals that have visited this vertexID
	visited := make(map[int32]map[int32]bool)

	for _, starterID := range starterIDs {
		visited[starterID] = map[int32]bool{starterID: true}
		q.rpush(vertex{starterID: starterID, vertexID: starterID, vertexDepth: 0})
	}

	for {

		if len(q) == 0 {
			break
		}
		popped := q.lpop()

		for _, adjacent := range adjacentsList[popped.vertexID] {

			val, ok := visited[adjacent]
			if ok {
				if val[popped.starterID] {
					continue
				}
			}

			if colors[adjacent-1] != reqColor {
				if val != nil {
					visited[adjacent][popped.starterID] = true
				} else {
					visited[adjacent] = map[int32]bool{popped.starterID: true}
				}
				q.rpush(vertex{starterID: popped.starterID, vertexID: adjacent, vertexDepth: popped.vertexDepth + 1})

			} else {
				return popped.vertexDepth + 1
			}
		}
	}
	return -1

}

func addToMap(m map[int32][]int32, key int32, val int32) {

	if _, ok := m[key]; !ok {
		m[key] = []int32{val}
		return
	}
	m[key] = append(m[key], val)
}
