package graphs

import (
	"testing"
)

var inputNodesAndColor = [][]int64{{5, 1}, {5, 2}, {5, 2}}
var inputEdges = [][][]int32{{{1, 2, 2, 3}, {2, 3, 4, 5}}, {{1, 2, 3, 2}, {2, 3, 4, 5}}, {{1, 2, 3, 2}, {2, 3, 4, 5}}}
var inputColors = [][]int64{{1, 2, 3, 1, 3}, {1, 3, 3, 5, 6}, {1, 2, 3, 5, 6}}
var outputDistances = []int32{2, -1, -1}

func TestFindShortest(t *testing.T) {

	for i := 0; i < len(inputNodesAndColor); i++ {
		res := findShortest(int32(inputNodesAndColor[i][0]), inputEdges[i][0], inputEdges[i][1], inputColors[i], inputNodesAndColor[i][1])
		if res != outputDistances[i] {
			t.Error("expected minimum distance", outputDistances[i], "for input nodes and colour:", inputNodesAndColor[i], "got", res)
		}
	}
}

func TestMinDistanceForColor(t *testing.T) {
	res := getMinDistanceForColor(map[int32][]int32{1: []int32{2}, 2: []int32{3, 4}, 3: []int32{2, 5}, 4: []int32{2}, 5: []int32{3}}, inputColors[0], []int32{1, 4}, inputNodesAndColor[0][1])

	if res != 2 {
		t.Error("expected 2 as the minimum distance, got", res)
	}
}
