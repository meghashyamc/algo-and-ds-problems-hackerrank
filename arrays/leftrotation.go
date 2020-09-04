package arrays

/*

A left rotation operation on an array shifts each of the array's elements 1 unit to the left.
 For example, if 2 left rotations are performed on array [1,2,3,4,5] , then the array would become .
[3,4,5,1,2]
Given an array a of n integers and a number d, perform  left rotations on the array.
 Return the updated array to be printed as a single line of space-separated integers.

 Complete the function rotLeft in the editor below. It should return the resulting array of integers.

rotLeft has the following parameter(s):

An array of integers a.
An integer d, the number of rotations.

Sample Input:
5 4
1 2 3 4 5

Sample Output:
5 1 2 3 4


*/
func rotLeft(a []int32, d int32) []int32 {

	if len(a) < 2 || d < 1 || d >= int32(len(a)) {
		return a
	}

	rotatedSlice := append(a[d:], a[:d]...)

	return rotatedSlice

}
