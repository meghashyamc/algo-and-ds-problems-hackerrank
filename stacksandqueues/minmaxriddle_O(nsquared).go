package stacksandqueues

// Given an integer array of size , find the maximum of the minimum(s) of every window size in the array. The window size varies from 1 to n.

// For example, given arr = [6,3,5,1,12], consider window sizes of 1 through 5. Windows of size 1 are (6),(3),(5),(1),(12). The maximum value of the minimum values of these windows is 12. Windows of size 2 are (6,3), (3,5), (5,1), (1,12) and their minima are (3,3,1,1). The maximum of these values is 3.
//Continue this process through window size 5 to finally consider the entire array. All of the answers are 12,3,3,1,1.

// Function Description

// Complete the riddle function in the editor below. It must return an array of integers representing the maximum minimum value for each window size from 1 to n.

// riddle has the following parameter(s):

// arr: an array of integers

//more info: https://www.hackerrank.com/challenges/min-max-riddle

func riddleNSquared(arr []int64) []int64 {

	maxofminArr := make([]int64, 0)
	maxofminArr = append(maxofminArr, maxElement(arr))
	minsOfPrevLevel := arr
	for {
		if len(minsOfPrevLevel) == 1 {
			break
		}
		minsOfThisLevel := make([]int64, 0)
		var newMaxofMins int64
		for i := 0; i < len(minsOfPrevLevel)-1; i++ {
			newmin := min(minsOfPrevLevel[i], minsOfPrevLevel[i+1])
			if newmin > newMaxofMins {
				newMaxofMins = newmin
			}
			minsOfThisLevel = append(minsOfThisLevel, newmin)
		}
		minsOfPrevLevel = minsOfThisLevel
		maxofminArr = append(maxofminArr, newMaxofMins)

	}
	return maxofminArr
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func maxElement(arr []int64) int64 {

	var max int64

	for _, el := range arr {
		if el > max {
			max = el
		}
	}
	return max
}
