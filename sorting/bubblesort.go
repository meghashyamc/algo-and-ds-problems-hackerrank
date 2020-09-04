package sorting

import "fmt"

// Consider the following version of Bubble Sort:

// for (int i = 0; i < n; i++) {

//     for (int j = 0; j < n - 1; j++) {
//         // Swap adjacent elements if they are in decreasing order
//         if (a[j] > a[j + 1]) {
//             swap(a[j], a[j + 1]);
//         }
//     }

// }
// Given an array of integers, sort the array in ascending order using the Bubble Sort algorithm above. Once sorted, print the following three lines:

// Array is sorted in numSwaps swaps., where numSwaps is the number of swaps that took place.
// First Element: firstElement, where firstElement is the first element in the sorted array.
// Last Element: lastElement, where lastElement is the last element in the sorted array.
// Hint: To complete this challenge, you must add a variable that keeps a running tally of all swaps that occur during execution.

// For example, given a worst-case but small array to sort: a =[6,4,1] we go through the following steps:

// swap    a
// 0       [6,4,1]
// 1       [4,6,1]
// 2       [4,1,6]
// 3       [1,4,6]
// It took 3 swaps to sort the array. Output would be

// Array is sorted in 3 swaps.
// First Element: 1
// Last Element: 6
// Function Description

// Complete the function countSwaps in the editor below. It should print the three lines required, then return.

// countSwaps has the following parameter(s):

// a: an array of integers .

func countSwaps(a []int32) {

	numSwaps := 0
	for i := 0; i < len(a); i++ {
		thisIterationSwaps := 0
		for j := 0; j < len(a)-1; j++ {

			if a[j] > a[j+1] {
				swap(a, j, j+1)
				numSwaps++
				thisIterationSwaps++
			}

		}
		if thisIterationSwaps == 0 {
			break
		}
	}

	fmt.Println("Array is sorted in", numSwaps, "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[(len(a)-1)])

}

func swap(arr []int32, i, j int) {

	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
