package linkedlists

import (
	"reflect"
	"testing"
)

var inputArrReverse = [][]int32{{1, 2, 3, 4, 0}, {0}, {1, 0}}
var answerArrReverse = [][]int32{{4, 3, 2, 1}, {}, {1}}

func TestReverse(t *testing.T) {
	for i, arr := range inputArrReverse {

		res := convertToArrDoubleList(reverse(createLinkedListDouble(arr)))
		if !reflect.DeepEqual(res, answerArrReverse[i]) {
			t.Error("expected the final linked list for", arr, "to be", answerArrReverse[i], "found", res)
		}
	}
}
