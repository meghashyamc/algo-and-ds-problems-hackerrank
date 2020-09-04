package dynamicprogramming

import "testing"

var arr []int32

func TestCandies(t *testing.T) {
	arr = []int32{1, 2, 2}
	res := candies(int32(len(arr)), arr)

	if res != 4 {
		t.Error("expected the number of candies to be 4 for", arr, "found", res)
	}

	arr = []int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}
	res = candies(int32(len(arr)), arr)

	if res != 19 {
		t.Error("expected the number of candies to be 19 for", arr, "found", res)
	}
}
