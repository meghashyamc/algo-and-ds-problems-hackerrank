package arrays

import (
	"errors"
	"log"
	"math"
)

/*
Given a 6 x 6 2D array arr, We define an hourglass in A to be a subset of values with indices falling in
this pattern in arr's graphical representation:

a b c
  d
e f g

There are 16 hourglasses in arr, and an hourglass sum is the sum of an hourglass' values. Calculate
 the hourglass sum for every hourglass in arr, then print the maximum hourglass sum.

For example, given the 2D array:

-9 -9 -9  1 1 1
 0 -9  0  4 3 2
-9 -9 -9  1 2 3
 0  0  8  6 6 0
 0  0  0 -2 0 0
 0  0  1  2 4 0

 Our highest hourglass value is 28 from the hourglass:

 0 4 3
   1
 8 6 6

 Complete the function hourglassSum in the editor below. It should return an integer, the maximum hourglass sum in the array.

hourglassSum has the following parameter(s):

arr: an array of integers

Constraints:
 -9 <= arr[i][j] <= 9
0 <= i,j <= 5

*/
const hourGlassWidth = 3
const hourGlassHeight = 3

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {

	// no hour glass possible if there are less than three rows
	if len(arr) < hourGlassHeight {
		return 0
	}

	if len(arr[0]) < hourGlassWidth {
		return 0
	}
	var maxSum int32 = math.MinInt32
	for rowIndex, row := range arr {

		// we are done, as an hour glass from this row onwards would not be possible
		if rowIndex > len(arr)-hourGlassHeight {
			break
		}
		for colindex := 0; colindex <= len(row)-hourGlassWidth; colindex++ {

			currSum, err := getHourGlassSum(&arr, rowIndex, colindex)
			if err != nil {
				log.Fatal("error in getting hour glass sum for one of the hour glasses" + err.Error())
			}

			if currSum > maxSum {

				maxSum = currSum
			}
		}

	}

	return maxSum

}

func getHourGlassSum(arr *([][]int32), rowIndex int, colIndex int) (int32, error) {

	if arr == nil {
		return 0, errors.New("nil array received")
	}
	var topSum int32 = 0
	var middleLineSum int32 = 0
	var botSum int32 = 0

	// get topsum
	for i := colIndex; i < colIndex+hourGlassWidth; i++ {

		topSum += (*arr)[rowIndex][i]
	}

	// get middle line sum
	for i := rowIndex + 1; i < rowIndex+hourGlassHeight-1; i++ {

		middleLineSum += (*arr)[i][colIndex+hourGlassWidth/2]
	}

	// get bottom sum
	for i := colIndex; i < colIndex+hourGlassWidth; i++ {

		botSum += (*arr)[rowIndex+hourGlassHeight-1][i]
	}

	return topSum + middleLineSum + botSum, nil
}
