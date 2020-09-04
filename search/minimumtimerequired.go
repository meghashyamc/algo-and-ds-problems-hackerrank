package search

// You are planning production for an order. You have a number of machines that each have a fixed number of days to produce an item. Given that all the machines operate simultaneously, determine the minimum number of days to produce the required order.

// For example, you have to produce goal = 10 items. You have three machines that take machines = [2,3,2] days to produce an item. The following is a schedule of items produced:

// Day Production  Count
// 2   2               2
// 3   1               3
// 4   2               5
// 6   3               8
// 8   2              10
// It takes 8 days to produce 10 items using these machines.

// Function Description

// Complete the minimumTime function in the editor below. It should return an integer representing the minimum number of days required to complete the order.

// minimumTime has the following parameter(s):

// machines: an array of integers representing days to produce one item per machine
// goal: an integer, the number of items required to complete the order
//more info: https://www.hackerrank.com/challenges/minimum-time-required

func minTime(machines []int64, goal int64) int64 {

	if goal == 0 {
		return 0
	}
	n := int64(len(machines))
	if n == 0 {
		return -1
	}

	min, max := getMinMax(machines)
	minTime, maxTime := (goal*min)/n, (goal*max)/n+max

	return getDaysToGoal(machines, minTime, maxTime, goal)
}

func getDaysToGoal(machines []int64, lo int64, hi int64, goal int64) int64 {

	var daysToGoalMin int64 = int64(^uint(0) >> 1)
	var achieved, mid int64
	for {

		mid = lo + (hi-lo)/2
		if lo >= hi {
			return min(mid, daysToGoalMin)
		}

		achieved = getTasksAchieved(machines, mid)

		if achieved > goal {

			hi = mid
		} else if achieved < goal {
			lo = mid + 1
		} else {

			if mid < daysToGoalMin {
				daysToGoalMin = mid
			}
			hi = mid

		}
	}
}

func getTasksAchieved(machines []int64, totalDays int64) int64 {

	var achieved int64
	for _, daysPerTask := range machines {

		achieved += totalDays / daysPerTask

	}
	return achieved

}

func getMinMax(machines []int64) (int64, int64) {
	var min int64 = int64(^uint(0) >> 1)
	var max int64
	for _, daysPerTask := range machines {

		if daysPerTask > max {
			max = daysPerTask
		}
		if daysPerTask < min {
			min = daysPerTask

		}

	}

	return min, max
}

func min(i, j int64) int64 {

	if i <= j {
		return i
	}
	return j
}
