package search

import (
	"fmt"
	"sort"
)

// You will be given an array of integers and a target value. Determine the number of pairs of array elements that have a difference equal to a target value.

// For example, given an array of [1, 2, 3, 4] and a target value of 1,
// we have three values meeting the condition: 2-1=1,3-1=1, and 4-3=1.

// Function Description

// Complete the pairs function below. It must return an integer representing the number of element pairs having the required difference.

// pairs has the following parameter(s):

// k: an integer, the target difference
// arr: an array of integers

//1 2 3 4 5

func pairs(k int32, arr []int32) int32 {

	n := len(arr)
	if n <= 1 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)

	var paircount int32
	for i, j := 0, 1; i < n-1 && j < n; {
		fmt.Println("i is", i, "and j is", j)
		diff := arr[j] - arr[i]
		if diff < k {
			j++
		} else if diff == k {
			paircount++
			if j+1 < n && arr[j] == arr[j+1] {
				j++
				continue
			}

			if arr[i] == arr[i+1] {
				i++
				continue
			}
			i++
			j++
		} else {
			i++
			if j == i {
				j++
			}
		}
	}
	return paircount
}
