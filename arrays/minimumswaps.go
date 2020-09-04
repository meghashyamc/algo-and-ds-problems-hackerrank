package arrays

import (
	"fmt"
)

/*

You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n]
 without any duplicates. You are allowed to swap any two elements.
 You need to find the minimum number of swaps required to sort the
  array in ascending order.

 Complete the function minimumSwaps in the editor below.
 It must return an integer representing the minimum number of swaps to sort the array.

minimumSwaps has the following parameter(s):

arr: an unordered array of integers

For example, given the array arr = [7,1,3,2,4,5,6], the function should return 5.


*/

func minimumSwaps(arr []int32) int32 {

	posMap := make(map[int32]int32)

	for index, num := range arr {

		posMap[num] = int32(index)
	}

	var swaps int32 = 0
	var i int32
	for i = 0; i < int32(len(arr)); i++ {
		fmt.Println(arr)

		if arr[i] == i+1 {
			fmt.Println("arr[i] == i+1")
			continue
		}

		if arr[i] > i+1 {
			fmt.Println("arr[i] > i+1")
			reqValAtThisPos := posMap[i+1]
			posMap[arr[i]] = posMap[i+1]
			posMap[i+1] = i
			swap(&arr, i, reqValAtThisPos)
			swaps++
			continue

		}
	}

	return swaps

}

func swap(sl *([]int32), i int32, j int32) {

	var temp int32
	temp = (*sl)[i]
	(*sl)[i] = (*sl)[j]
	(*sl)[j] = temp

}
