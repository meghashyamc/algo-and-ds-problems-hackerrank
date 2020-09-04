package greedy

import (
	"sort"
)

// Consider an array of integers,arr = [arr[0], arr[1],...,arr[n-1]].
// We define the absolute difference between two elements, a[i] and a[j] (where i != j), to be the absolute value of a[i]-a[j] .

// Given an array of integers, find and print the minimum absolute difference between any two elements in the array. For example, given the array arr = [-2,2,4]
// we can create 3 pairs of numbers: [-2,2],[-2,4] and [2,4]. The absolute differences for these pairs
// are 4,  6 and 2 respectively. The minimum absolute difference is 2 .

// Function Description

// Complete the minimumAbsoluteDifference function in the editor below. It should return an integer that represents the minimum absolute difference between any pair of elements.

// minimumAbsoluteDifference has the following parameter(s):

// n: an integer that represents the length of arr
// arr: an array of integers

func minimumAbsoluteDifference(arr []int32) int32 {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	//max value of int32 in Go
	mindiff := int32(^(uint32(0)) >> 1)

	for i, el := range arr {

		if i > 0 {
			currdiff := el - arr[i-1]
			if currdiff < mindiff {
				mindiff = currdiff
			}
		}
	}

	return mindiff

}
