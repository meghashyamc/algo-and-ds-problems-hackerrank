package greedy

import "sort"

// Lena is preparing for an important coding competition that is preceded by a number of sequential preliminary contests. Initially, her luck balance is 0. She believes in "saving luck", and wants to check her theory. Each contest is described by two integers,  and :

// L[i] is the amount of luck associated with a contest. If Lena wins the contest, her luck balance will decrease by L[i]; if she loses it, her luck balance will increase by L[i].
// T[i] denotes the contest's importance rating. It's equal to 1 if the contest is important, and it's equal to 0 if it's unimportant.
// If Lena loses no more than k important contests, what is the maximum amount of luck she can have after competing in all the preliminary contests? This value may be negative.

// For example, k=2 and:

// Contest		L[i]	T[i]
// 1		5	1
// 2		1	1
// 3		4	0
// If Lena loses all of the contests, her luck will be 5+4+1=10. Since she is allowed to lose 2 important contests, and there are only 2 important contests. She can lose all three contests to maximize her luck at 10. If k=1, she has to win at least 1 of the 2 important contests.
// She would choose to win the lowest value important contest worth 1. Her final luck will be 5+4-1=8.

// Function Description

// Complete the luckBalance function in the editor below. It should return an integer that represents the maximum luck balance achievable.

// luckBalance has the following parameter(s):

// k: the number of important contests Lena can lose
// contests: a 2D array of integers where each contests[i] contains two integers that represent the luck balance and importance of the ith contest.
// Constraints:
// 1 <= n <= 100
// 0 <= k <= N
// 1 <= L[i] <= 10^4
// T[i] belongs to [0,1]

func luckBalance(k int32, contests [][]int32) int32 {

	var maxLuck int32
	luckList := make([]int32, 0)

	//add luck for unimportant contests
	for _, contest := range contests {

		if contest[1] == 0 {
			maxLuck += contest[0]
			continue
		}

		luckList = append(luckList, contest[0])
	}

	//sort only important contests from high luck to low luck
	sort.Slice(luckList, func(i, j int) bool {
		return luckList[i] > luckList[j]
	})

	if len(luckList) == 0 {
		return maxLuck
	}

	var maxcontestsToLose = 0
	if k > int32(len(luckList)) {
		maxcontestsToLose = len(luckList)
	} else {
		maxcontestsToLose = int(k)
	}

	for i := 0; i < maxcontestsToLose; i++ {
		maxLuck += luckList[i]
	}

	//case when k is higher than number of important contests
	// so, we can add luck for all important contests
	if k > int32(maxcontestsToLose) {
		return maxLuck
	}

	//case when k is less than or equal to number of important contests
	//we need to subtract luck for important contests that we were not able to lose
	for i := int(k); i < len(luckList); i++ {
		maxLuck -= luckList[i]
	}

	return maxLuck

}
