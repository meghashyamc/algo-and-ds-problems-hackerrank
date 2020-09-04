package arrays

import (
	"errors"
	"fmt"
)

/*

Starting with a 1-indexed array of zeros and a list of operations, for each operation
 add a value to each of the array element between two given indices, inclusive.
  Once all operations have been performed, return the maximum value in your array.

For example, the length of your array of zeros, n = 10 . Your list of queries is as follows:

     a b k
    1 5 3
    4 8 7
	6 9 1


	Add the values of k between the indices a and b inclusive:

	index->	 1 2 3  4  5 6 7 8 9 10
	[0,0,0, 0, 0,0,0,0,0, 0]
	[3,3,3, 3, 3,0,0,0,0, 0]
	[3,3,3,10,10,7,7,7,0, 0]
	[3,3,3,10,10,8,8,8,1, 0]


	The largest value is 10 after all operations are performed.

Complete the function arrayManipulation in the editor below. It must return an integer,
the maximum value in the resulting array.

arrayManipulation has the following parameters:

n - the number of elements in your array
queries - a two dimensional array of queries where each queries[i] contains three integers,
 a, b, and k.


*/

//n is the number of elements in the original array of zeroes
func arrayManipulation(n int32, queries [][]int32) int64 {

	sliceToManipulate := make([]int32, n+1)

	for _, queryInterval := range queries {
		// fmt.Println("considering query interval...", queryInterval)

		if len(queryInterval) < 3 {
			panic(errors.New("error in input, length of array can't be less than 3"))
		}
		sliceToManipulate[queryInterval[0]] += queryInterval[2]
		fmt.Println("adding k...sliceToManipulate is", sliceToManipulate)
		if queryInterval[1]+1 <= n {
			sliceToManipulate[queryInterval[1]+1] -= queryInterval[2]
		}
		// fmt.Println("subtracting k...sliceToManipulate is", sliceToManipulate)

	}

	var maxval int64
	var sum int64
	for _, valAtPoint := range sliceToManipulate {

		sum += int64(valAtPoint)
		// fmt.Println("sum is...", sum)
		if sum > maxval {
			maxval = sum
		}
		// fmt.Println("maxval is...", maxval)

	}

	return maxval

}
