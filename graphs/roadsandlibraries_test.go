package graphs

import (
	"testing"
)

var inputLibsAndRoads = [][]int32{{3, 2, 1}, {6, 2, 5}, {7, 3, 2}, {9, 91, 84}, {1, 5, 3}}
var inputCities = [][][]int32{{{1, 2}, {3, 1}, {2, 3}}, {{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}}, {{7, 1}, {1, 3}, {3, 2}, {1, 2}, {4, 6}, {5, 6}}, {{8, 2}, {2, 9}}, {}}
var outputRoadsAndLibraries = []int64{4, 12, 16, 805, 5}

func TestRoadsAndLibraries(t *testing.T) {

	for i := 0; i < len(inputLibsAndRoads); i++ {
		res := roadsAndLibraries(inputLibsAndRoads[i][0], inputLibsAndRoads[i][1], inputLibsAndRoads[i][2], inputCities[i])
		if res != outputRoadsAndLibraries[i] {
			t.Error("expected cost of roads and libraries to be", outputRoadsAndLibraries[i], "got", res, "for input", inputLibsAndRoads[i])
		}
	}
}

func TestUnionFind(t *testing.T) {
	n := 10
	cityRoots = make([]int32, n)
	initialize(cityRoots)
	size = make([]int32, n)

	connect(0, 1)
	connect(1, 2)
	connect(4, 5)
	if !areConnected(0, 2) {
		t.Error("expected 0 and 2 to be connected, but they were not connected")
	}
	if areConnected(0, 4) {
		t.Error("expected 0 and 4 to be not connected, but they were  connected")
	}

	root1 := root(1)
	root2 := root(2)
	if root1 != 0 {
		t.Error("expected 0 to be the root of 1, found", root1)
	}
	if root2 != 0 {
		t.Error("expected 0 to be the root of 2, found", root2)
	}
}
