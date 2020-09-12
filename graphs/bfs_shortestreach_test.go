package graphs

import (
	"reflect"
	"testing"
)

var inputNodesAndStartBFS = [][]int{{6, 1}, {4, 1}, {3, 2}}
var inputEdgesBFS = [][][]int{{{1, 2, 3, 1}, {2, 3, 4, 5}}, {{1, 1}, {2, 3}}, {{2}, {3}}}
var outputDistancesBFS = [][]int{{6, 12, 18, 6, -1}, {6, 6, -1}, {-1, 6}}

func TestFindDistancesFromStart(t *testing.T) {

	for i := 0; i < len(inputNodesAndStartBFS); i++ {
		res := findDistancesFromStart(inputNodesAndStartBFS[i][0], inputEdgesBFS[i][0], inputEdgesBFS[i][1], inputNodesAndStartBFS[i][1])
		if !reflect.DeepEqual(res, outputDistancesBFS[i]) {
			t.Error("expected distances list to be", outputDistancesBFS[i], "for input nodes", inputNodesAndStartBFS, "got", res)
		}
	}
}
