package stacksandqueues

import (
	"testing"
)

var testArrays = [][]int32{{8, 4, 6, 5, 7}, {3, 2, 1}, {1, 2, 3, 4, 5}, {3, 2, 3}, {8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116}}
var answerArray = []int64{20, 4, 9, 6, 26152}

func TestLargestRectangle(t *testing.T) {

	for i, arr := range testArrays {

		if res := largestRectangle(arr); res != answerArray[i] {
			t.Error("for array", arr, "expected answer", answerArray[i], "got", res)
		}
	}

}
