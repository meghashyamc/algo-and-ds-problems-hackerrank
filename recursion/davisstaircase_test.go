package recursion

import (
	"testing"
)

var inputsStaircase = []int32{-1, 0, 1, 3, 5, 7}
var outputsStaircase = []int32{0, 0, 1, 4, 13, 44}

func TestStepPerms(t *testing.T) {

	for i, input := range inputsStaircase {
		res := stepPerms(input)
		if res != outputsStaircase[i] {
			t.Error("expected", outputsStaircase[i], "as stepPerms(", input, "), got", res)
		}
	}
}
