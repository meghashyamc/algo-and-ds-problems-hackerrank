package stacksandqueues

import (
	"reflect"
	"testing"
)

var testArraysRiddle = [][]int64{{2, 6, 1, 12}, {1, 2, 3, 5, 1, 13, 3}, {3, 5, 4, 7, 6, 2}, {10, 3, 10, 10, 10, 1, 3, 10}}
var answerArraysRiddle = [][]int64{{12, 2, 1, 1}, {13, 3, 2, 1, 1, 1, 1}, {7, 6, 4, 4, 3, 2}, {10, 10, 10, 3, 3, 1, 1, 1}}

func TestRiddle(t *testing.T) {

	for i, arr := range testArraysRiddle {

		res := riddle(arr)
		if !reflect.DeepEqual(res, answerArraysRiddle[i]) {

			t.Error("expected answer", answerArraysRiddle[i], "for", arr, "got", res)
		}
	}
}

func TestGetNextOrPrevSmallerArray(t *testing.T) {

	arr := []int64{10, 20, 30, 50, 10, 70, 30}
	nextSmallerArray := []int{7, 4, 4, 4, 7, 6, 7}
	prevSmallerArray := []int{-1, 0, 1, 2, -1, 4, 4}
	res := make([]int, len(arr))
	getNextOrPrevSmallerArray(false, arr, res)
	if !reflect.DeepEqual(res, nextSmallerArray) {
		t.Error("expected next smaller array", nextSmallerArray, "got", res)
	}
	res = make([]int, len(arr))

	getNextOrPrevSmallerArray(true, arr, res)
	if !reflect.DeepEqual(res, prevSmallerArray) {
		t.Error("expected next smaller array", prevSmallerArray, "got", res)
	}

}
