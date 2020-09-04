package search

import (
	"sort"
)

// We define the following:

// A subarray of array a of length n is a contiguous segment from a[i] through a[j] where 0 <=i<=j<n.
// The sum of an array is the sum of its elements.
// Given an n element array of integers,a , and an integer,m , determine the maximum value of the sum of any of its subarrays modulo m. For example, Assume a = [1,2,3] and m=2. The following table lists all subarrays and their moduli:

// 		sum	%2
// [1]		1	1
// [2]		2	0
// [3]		3	1
// [1,2]		3	1
// [2,3]		5	1
// [1,2,3]		6	0
// The maximum modulus is 1.

// Function Description

// Complete the maximumSum function in the editor below. It should return a long integer that represents the maximum value of .

// maximumSum has the following parameter(s):

// a: an array of long integers, the array to analyze
// m: a long integer, the modulo divisor
//more info: https://www.hackerrank.com/challenges/maximum-subarray-sum

type prefixSumIndex struct {
	prefixSum int64
	index     int
}

func MaximumSum(a []int64, m int64) int64 {

	if len(a) == 0 {
		return 0
	}
	//in this slice, an element at index i
	// contains (sum of elements till index i)%m
	//also contains index i
	//note: (a+b)%m = (a%m + b%m)%m
	prefixSumArr := make([]prefixSumIndex, 0)

	for i, el := range a {
		if i == 0 {
			prefixSumArr = append(prefixSumArr, prefixSumIndex{prefixSum: el % m, index: i})
			continue
		}
		prefixSumArr = append(prefixSumArr, prefixSumIndex{prefixSum: getModSum(el, prefixSumArr[i-1].prefixSum, m), index: i})
	}

	//sort in descending order of prefix sums
	sort.Slice(prefixSumArr, func(i, j int) bool {

		return prefixSumArr[i].prefixSum > prefixSumArr[j].prefixSum
	})

	//get max prefix sum currently present
	var maxSum int64 = getMaxPrefixSum(prefixSumArr)
	// fmt.Println("prefixSumArr:", prefixSumArr)

	//see if any difference of consecutive prefix sums (such that the indices of the sums are in reverse order compared to the prefix sums) is greater than
	//our max prefix sum
	for i := range prefixSumArr {

		if i != len(prefixSumArr)-1 {
			if prefixSumArr[i].index < prefixSumArr[i+1].index {
				// fmt.Println("looking at ", prefixSumArr[i], "and", prefixSumArr[i+1])
				diffInSums := getDifferenceMod(prefixSumArr[i].prefixSum, prefixSumArr[i+1].prefixSum, m)
				if diffInSums > maxSum {
					maxSum = diffInSums
				}
			}
		}
	}
	return maxSum
}

func getModSum(a, b, m int64) int64 {

	return (a%m + b) % m
}

func getDifferenceMod(a, b, m int64) int64 {

	return (b - a + m) % m
}

func getMaxPrefixSum(arr []prefixSumIndex) int64 {

	if len(arr) == 0 {
		return 0
	}
	maxPossInt64 := ^uint(0) >> 1
	var maxSum = int64(-maxPossInt64 - 1)
	for _, el := range arr {

		if el.prefixSum > maxSum {
			maxSum = el.prefixSum
		}

	}
	return maxSum

}
