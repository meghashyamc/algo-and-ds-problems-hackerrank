package hashtables

// You are given an array and you need to find number of tripets of indices (i,j,k)  such that the elements at those indices are in geometric progression for a given common ratio r and i < j < k.

// For example,arr=[1,4,16,64]. If r=4, we have [1,4,16] and[4,16,64]  at indices (0,1,2) and (1,2,3).

// Function Description

// Complete the countTriplets function in the editor below. It should return the number of triplets forming a geometric progression for a given  as an integer.

// countTriplets has the following parameter(s):

// arr: an array of integers
// r: an integer, the common ratio

func CountTriplets(arr []int64, r int64) int64 {

	var tripletsCount int64

	//key = index element of form a
	//val = number of elements of form ar to a's right
	elementsOfFormAROnRight := make(map[int]int64)

	//key = element
	//value = number of occurences of the element
	//encountered so far in our iteration
	elementsCount := make(map[int64]int64)

	//iterate over arr from right to left

	for i := len(arr) - 1; i >= 0; i-- {

		elementsOfFormAROnRight[i] = elementsCount[arr[i]*r]
		elementsCount[arr[i]]++
	}

	//key = index element of form a
	//val = number of elements of form a/r to a's left
	elementsOfFormAByROnLeft := make(map[int]int64)

	//key = element
	//value = number of occurences of the element
	//encountered so far in our iteration
	elementsCount = make(map[int64]int64)

	for i := 0; i < len(arr); i++ {

		if arr[i]%r == 0 {
			elementsOfFormAByROnLeft[i] = elementsCount[arr[i]/r]
		} else {
			elementsOfFormAByROnLeft[i] = 0
		}
		elementsCount[arr[i]]++
	}
	for i := 0; i < len(arr); i++ {
		tripletsCount += elementsOfFormAByROnLeft[i] * elementsOfFormAROnRight[i]
	}
	return tripletsCount
}
