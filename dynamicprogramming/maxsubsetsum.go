package dynamicprogramming

// Given an array of integers, find the subset of non-adjacent elements with the maximum sum. Calculate the sum of that subset.

// For example, given an array arr = [-2,1,3,-4,5]  we have the following possible subsets:

// Subset      Sum
// [-2, 3, 5]   6
// [-2, 3]      1
// [-2, -4]    -6
// [-2, 5]      3
// [1, -4]     -3
// [1, 5]       6
// [3, 5]       8
// Our maximum subset sum is 8.

// Function Description

// Complete the maxSubsetSum function in the editor below. It should return an integer representing the maximum subset sum for the given array.

// maxSubsetSum has the following parameter(s):

// arr: an array of integers
//more info: https://www.hackerrank.com/challenges/max-array-sum/

var memCacheSubsetSum = make(map[int]int32)

func maxSubsetSum(arr []int32) int32 {

	if len(arr) == 0 {
		return 0
	}

	return maxSubsetSumHelper(arr, 0)

}

//returns max subset sum from index = index to index = len(arr) -1
func maxSubsetSumHelper(arr []int32, index int) int32 {

	if index > len(arr)-1 {
		return 0
	}

	cachedVal, ok := memCacheSubsetSum[index]
	if ok {
		return cachedVal
	}

	if index == len(arr)-1 {
		memCacheSubsetSum[index] = arr[len(arr)-1]
		return memCacheSubsetSum[index]
	}

	memCacheSubsetSum[index] = max(arr[index], arr[index]+maxSubsetSumHelper(arr, index+2), maxSubsetSumHelper(arr, index+1))
	return memCacheSubsetSum[index]
}

func max(arr ...int32) int32 {

	maxInt32 := int32(^uint32(0) >> 1)
	var max int32 = -maxInt32 - 1
	for _, el := range arr {
		if el > max {
			max = el
		}
	}

	return max
}
