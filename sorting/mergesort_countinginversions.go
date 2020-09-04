package sorting

// In an array,arr , the elements at indices i and j (where i < j) form an inversion if arr[i] > arr[j] . In other words, inverted elements arr[i] and arr[j] are considered to be "out of order". To correct an inversion, we can swap adjacent elements.

// For example, consider the dataset arr = [2,4,1]. It has two inversions: (4,1) and (2,1). To sort the array, we must perform the following two swaps to correct the inversions:
//arr = [2,4,1] -> swap(arr[1], arr[2]) -> swap (arr[0], arr[1]) -> [1,2,4]

// Given d datasets, print the number of inversions that must be swapped to sort each dataset on a new line.

// Function Description

// Complete the function countInversions in the editor below. It must return an integer representing the number of inversions required to sort the array.

// countInversions has the following parameter(s):

// arr: an array of integers to sort .

func CountInversions(arr []int32) int64 {
	if len(arr) <= 1 {
		return 0
	}
	auxArray := make([]int32, len(arr))
	copy(auxArray, arr)

	return sortArray(arr, 0, len(arr), auxArray)
}

func sortArray(arr []int32, st, end int, auxArray []int32) int64 {

	// fmt.Println("inside sortArray...arr = ", arr, "st=", st, "end=", end)
	if end <= st+1 {
		return 0
	}

	mid := st + (end-st)/2

	//fmt.Println("will now call sortArray(arr, st, mid):", st, mid)
	inversionCount := sortArray(auxArray, st, mid, arr)
	inversionCount += sortArray(auxArray, mid, end, arr)
	inversionCount += merge(auxArray, st, mid, end, arr)
	return inversionCount

}

func merge(arr []int32, st, mid, end int, auxArray []int32) int64 {

	var inversionCount int64 = 0
	for i, j, k := st, mid, st; ; k++ {

		if i >= mid && j >= end {
			break
		}
		if i >= mid {
			auxArray[k] = arr[j]
			j++
			continue
		}

		if j >= end {
			auxArray[k] = arr[i]
			i++
			continue
		}
		if arr[i] <= arr[j] {

			auxArray[k] = arr[i]
			i++
			continue
		}

		// fmt.Println("inversion count increase for inverting", arr[i], arr[j-1])
		auxArray[k] = arr[j]
		j++
		inversionCount += int64(mid) - int64(i)

		continue

	}

	return inversionCount

}
