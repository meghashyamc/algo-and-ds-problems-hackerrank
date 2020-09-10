package stacksandqueues

import (
	"errors"
)

// Given an integer array of size , find the maximum of the minimum(s) of every window size in the array. The window size varies from 1 to n.

// For example, given arr = [6,3,5,1,12], consider window sizes of 1 through 5. Windows of size 1 are (6),(3),(5),(1),(12). The maximum value of the minimum values of these windows is 12. Windows of size 2 are (6,3), (3,5), (5,1), (1,12) and their minima are (3,3,1,1). The maximum of these values is 3.
//Continue this process through window size 5 to finally consider the entire array. All of the answers are 12,3,3,1,1.

// Function Description

// Complete the riddle function in the editor below. It must return an array of integers representing the maximum minimum value for each window size from 1 to n.

// riddle has the following parameter(s):

// arr: an array of integers

//more info: https://www.hackerrank.com/challenges/min-max-riddle

type intStack []int

func (s *intStack) push(i int) {
	if s == nil {
		s = &intStack{}
	}
	*s = append(*s, i)
}

func (s *intStack) pop() int {
	if s == nil || len(*s) == 0 {
		panic(errors.New("can't pop from empty stack"))
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *intStack) peek() int {
	if s == nil || len(*s) == 0 {
		panic(errors.New("can't peek into empty stack"))
	}
	return (*s)[len(*s)-1]

}
func riddle(arr []int64) []int64 {

	nextSmallerArray := make([]int, len(arr))
	//populate next smaller number's index in array
	getNextOrPrevSmallerArray(false, arr, nextSmallerArray)
	prevSmallerArray := make([]int, len(arr))
	getNextOrPrevSmallerArray(true, arr, prevSmallerArray)

	minMaxAnswerArr := make([]int64, len(arr)+1)
	for i := 0; i < len(minMaxAnswerArr); i++ {
		minMaxAnswerArr[i] = -1
	}

	for i := 0; i < len(arr); i++ {

		windowLen := nextSmallerArray[i] - prevSmallerArray[i] - 1
		if minMaxAnswerArr[windowLen] == -1 || (minMaxAnswerArr[windowLen] < arr[i]) {
			minMaxAnswerArr[windowLen] = arr[i]
		}

	}

	for i := len(arr) - 1; i >= 1; i-- {
		minMaxAnswerArr[i] = max(minMaxAnswerArr[i], minMaxAnswerArr[i+1])

	}
	return minMaxAnswerArr[1:]
}

func max(a, b int64) int64 {

	if a > b {
		return a
	}
	return b
}

func getNextOrPrevSmallerArray(prev bool, arr []int64, smallerIndexArr []int) {
	s := intStack{}
	var extremityIndex int
	if !prev {
		extremityIndex = len(arr)
	} else {
		extremityIndex = -1
	}

	if !prev {
		for i := 0; i < len(arr); i++ {

			if i == 0 {
				s.push(i)
				continue
			}

			updateStack(i, &s, arr, smallerIndexArr, extremityIndex)
		}

	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			if i == len(arr)-1 {
				s.push(i)
				continue
			}
			updateStack(i, &s, arr, smallerIndexArr, extremityIndex)
		}
	}
	// fmt.Println("stack after loop:", s)
	if len(s) != 0 {

		for _, remainingIndex := range s {

			smallerIndexArr[remainingIndex] = extremityIndex

		}
	}
	// fmt.Println("final stack:", s)
}

func updateStack(i int, s *intStack, arr []int64, smallerIndexArr []int, extremityIndex int) {
	// fmt.Println("stack:", s)
	if arr[i] < arr[s.peek()] {
		// fmt.Println("comparing a[i] and to stack peek:", arr[i], (*s).peek().val)
		for j := len(*s) - 1; j >= 0; j-- {
			if arr[i] < arr[s.peek()] {
				smallerIndexArr[s.pop()] = i
			} else {
				break
			}
		}
	}
	s.push(i)
	// fmt.Println("stack after updation:", s)
}
