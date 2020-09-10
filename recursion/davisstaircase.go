package recursion

// Davis has a number of staircases in his house and he likes to climb each staircase , , or  steps at a time. Being a very precocious child, he wonders how many ways there are to reach the top of the staircase.

// Given the respective heights for each of the s staircases in his house, find and print the number of ways he can climb each staircase, modulo 10^7 + 7  on a new line.

// For example, there is s=1  staircase in the house that is n=5 steps high. Davis can step on the following sequences of steps:

// 1 1 1 1 1
// 1 1 1 2
// 1 1 2 1
// 1 2 1 1
// 2 1 1 1
// 1 2 2
// 2 2 1
// 2 1 2
// 1 1 3
// 1 3 1
// 3 1 1
// 2 3
// 3 2
// There are 13 possible ways he can take these 5 steps. 13 modulo 10^7+7 = 13.

// Function Description

// Complete the stepPerms function in the editor below. It should recursively calculate and return the integer number of ways Davis can climb the staircase, modulo 10000000007.

// stepPerms has the following parameter(s):

// n: an integer, the number of stairs in the staircase
//more info: https://www.hackerrank.com/challenges/ctci-recursive-staircase

const modConst = 10000000007

func stepPerms(n int32) int32 {

	if n < 0 {
		return 0
	}
	memCache := make([]int32, n+1)
	for i := range memCache {
		memCache[i] = -1
	}
	return stepPermsHelper(n, memCache)

}

func stepPermsHelper(n int32, memCache []int32) int32 {

	if memCache[n] != -1 {
		return memCache[n]
	}

	switch n {

	case 0:
		memCache[n] = 0

	case 1:
		memCache[n] = 1
	case 2:
		memCache[n] = stepPermsHelper(1, memCache) + 1

	case 3:
		memCache[n] = stepPermsHelper(2, memCache) + stepPermsHelper(1, memCache) + 1

	default:
		memCache[n] = int32((int64(stepPermsHelper(n-1, memCache))%modConst + int64(stepPermsHelper(n-2, memCache))%modConst + int64(stepPermsHelper(n-3, memCache))%modConst) % modConst)

	}

	return memCache[n]
}
