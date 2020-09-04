package dynamicprogramming

// Alice is a kindergarten teacher. She wants to give some candies to the children in her class.  All the children sit in a line and each of them has a rating score according to his or her performance in the class.  Alice wants to give at least 1 candy to each child. If two children sit next to each other, then the one with the higher rating must get more candies. Alice wants to minimize the total number of candies she must buy.

// For example, assume her students' ratings are [4, 6, 4, 5, 6, 2]. She gives the students candy in the following minimal amounts: [1, 2, 1, 2, 3, 1]. She must buy a minimum of 10 candies.

// Function Description

// Complete the candies function in the editor below. It must return the minimum number of candies Alice must buy.

// candies has the following parameter(s):

// n: an integer, the number of children in the class
// arr: an array of integers representing the ratings of each student
//more info: https://www.hackerrank.com/challenges/candies

func candies(n int32, arr []int32) int64 {

	if n == 0 {
		return 0
	}

	memCache := make([]int64, n+1)
	var sum int64
	for i := range arr {

		sum += getCandiesForIndex(memCache, arr, i)
	}
	return sum
}

func getCandiesForIndex(memCache []int64, arr []int32, index int) int64 {

	if index > len(arr)-1 {
		return 0
	}
	if memCache[index] != 0 {
		return memCache[index]
	}

	if index == 0 {
		if len(arr) == 1 || arr[index+1] >= arr[index] {
			memCache[index] = 1

		} else {
			memCache[index] = 1 + getCandiesForIndex(memCache, arr, index+1)
		}
	} else if index == len(arr)-1 {
		if arr[index-1] >= arr[index] {
			memCache[index] = 1
		} else {
			memCache[index] = 1 + getCandiesForIndex(memCache, arr, index-1)
		}
	} else if arr[index] <= arr[index-1] && arr[index] <= arr[index+1] {
		memCache[index] = 1

	} else if arr[index] > arr[index-1] && arr[index] > arr[index+1] {

		memCache[index] = maxInt64(getCandiesForIndex(memCache, arr, index+1), getCandiesForIndex(memCache, arr, index-1)) + 1
	} else if arr[index] > arr[index-1] && arr[index] <= arr[index+1] {
		memCache[index] = getCandiesForIndex(memCache, arr, index-1) + 1
	} else if arr[index] <= arr[index-1] && arr[index] > arr[index+1] {
		memCache[index] = getCandiesForIndex(memCache, arr, index+1) + 1
	}
	return memCache[index]
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
