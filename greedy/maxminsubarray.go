package greedy

import "sort"

// You will be given a list of integers,arr , and a single integer k. You must create an array of length k from elements of arr such that its unfairness is minimized. Call that array subarr.
//Unfairness of an array is calculated as
//max(subarr) - min(subarr)

// Where:
// - max denotes the largest integer in subarr
// - min denotes the smallest integer in subarr

// As an example, consider the array [1,4,7,2] with a k of 2. Pick any two elements, test subarr = [4,7].
// unfairness = max(4,7) - min(4,7) = 7-4 = 3
// Testing for all pairs, the solution [1,2] provides the minimum unfairness.

// Note: Integers in arr may not be unique.

// Function Description

// Complete the maxMin function in the editor below. It must return an integer that denotes the minimum possible value of unfairness.

// maxMin has the following parameter(s):

// k: an integer, the number of elements in the array to create
// arr: an array of integers.
// assume 2 <= k <= length of arr

//more details: https://www.hackerrank.com/challenges/angry-children

func maxMin(k int32, arr []int32) int32 {

	if k <= 1 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	maxval := int32(^(uint32(0)) >> 1)
	minUnfairness := maxval

	for i, num := range arr {
		if i+int(k) > len(arr) {
			break
		}
		currUnfairness := num - arr[i+int(k)-1]
		if currUnfairness < minUnfairness {
			minUnfairness = currUnfairness
		}
	}

	if minUnfairness == maxval {
		return 0
	}
	return minUnfairness

}
