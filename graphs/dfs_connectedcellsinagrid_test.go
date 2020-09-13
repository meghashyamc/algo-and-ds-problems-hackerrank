package graphs

import (
	"reflect"
	"testing"
)

var inputMatrices = [][][]int32{{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {1, 0, 0, 0}}, {{1, 1, 0, 0, 0, 1}, {0, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0}}}
var outputMaxRegionVals = []int32{5, 5}

func TestMaxRegion(t *testing.T) {

	for i, matrix := range inputMatrices {

		res := maxRegion(matrix)
		if res != outputMaxRegionVals[i] {
			t.Error("expected output", outputMaxRegionVals[i], "for matrix", matrix, "got", res)
		}
	}

}

func TestGetValForRegion(t *testing.T) {
	visited := make([][]bool, 0)
	grid := inputMatrices[0]
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}

	res := getValForRegion(visited, grid, position{i: 0, j: 0})
	if res != 5 {
		t.Error("expected region value to be 5, got", res)
	}
}

func TestGetAdjacents(t *testing.T) {

	res := getAdjacents(inputMatrices[0], position{i: 0, j: 0})
	expectedRes := []position{position{i: 0, j: 1}, position{i: 1, j: 1}}
	if !reflect.DeepEqual(res, expectedRes) {
		t.Error("expected adjacents to be", expectedRes, "found", res)
	}
	res = getAdjacents(inputMatrices[0], position{i: 1, j: 1})
	expectedRes = []position{position{i: 0, j: 1}, position{i: 1, j: 2}, position{i: 0, j: 0}, position{i: 2, j: 2}}
	if !haveSameElements(res, expectedRes) {
		t.Error("expected adjacents to be", expectedRes, "found", res)
	}
}

func haveSameElements(arr1 []position, arr2 []position) bool {

	if len(arr1) != len(arr2) {
		return false
	}
	listofelements := make(map[position]bool)

	for _, pos := range arr1 {
		listofelements[pos] = true
	}

	for _, pos := range arr2 {
		delete(listofelements, pos)
	}
	return len(listofelements) == 0

}
