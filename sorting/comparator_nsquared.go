package sorting

import (
	"sort"
)

// HackerLand National Bank has a simple policy for warning clients about possible fraudulent account activity. If the amount spent by a client on a particular day is greater than or equal to 2 X the client's median spending for a trailing number of days, they send the client a notification about potential fraud. The bank doesn't send the client any notifications until they have at least that trailing number of prior days' transaction data.

// Given the number of trailing days  and a client's total daily expenditures for a period of  days, find and print the number of times the client will receive a notification over all  days.

// For example, d = 3 and expenditures = [10,20,30,40,50]. On the first three days, they just collect spending data. At day 4, we have trailing expenditures of [10,20,30] . The median is 20 and the day's expenditure is 40. Because 40>=2 x 20, there will be a notice.
//The next day, our trailing expenditures are [20,30,40] and the expenditures are 50. This is less than 2 x 30 = 60 so no notice will be sent. Over the period, there was one notice sent.

// Note: The median of a list of numbers can be found by arranging all the numbers from smallest to greatest. If there is an odd number of numbers, the middle one is picked.
// If there is an even number of numbers, median is then defined to be the average of the two middle values.

// Function Description

// Complete the function activityNotifications in the editor below. It must return an integer representing the number of client notifications.

// activityNotifications has the following parameter(s):

// expenditure: an array of integers representing daily expenditures
// d: an integer, the lookback days for median spending

func activityNotifications(expenditure []int32, d int32) int32 {

	lookback := make([]int32, d)
	copy(lookback, expenditure[0:d])

	sort.Slice(lookback, func(i, j int) bool {
		return lookback[i] < lookback[j]
	})

	median := getMedian(lookback)

	var notificationCount int32
	for _, dayExpense := range expenditure[d:len(expenditure)] {

		if float32(dayExpense) >= median*2 {
			// fmt.Println(float32(dayExpense), ">=", median*2)
			notificationCount++
		}

		if dayExpense != expenditure[len(expenditure)-1] {
			median = getNewMedian(lookback, dayExpense)
			// fmt.Println("next median is...", median)
		}
	}
	return notificationCount

}

func getMedian(arr []int32) float32 {
	// fmt.Println("will get median for...", arr)
	n := len(arr)
	var median float32
	if n%2 == 0 {
		median = (float32(arr[n/2]) + float32(arr[(n/2)-1])) / 2.0
	} else {
		median = float32(arr[n/2])
	}
	// fmt.Println("this is the median:", median)
	return median
}

func getNewMedian(sortedArr []int32, newEl int32) float32 {

	n := len(sortedArr)
	newSortedArr := make([]int32, 0)

	for _, el := range sortedArr {

		if el > newEl {

			newSortedArr = append(newSortedArr, newEl)
			newSortedArr = append(newSortedArr, sortedArr[el:n]...)
			sortedArr = newSortedArr[1:len(newSortedArr)]
			return getMedian(sortedArr)
		}
		newSortedArr = append(newSortedArr, el)
	}
	newSortedArr = append(newSortedArr, newEl)
	sortedArr = newSortedArr[1:len(newSortedArr)]

	return getMedian(sortedArr)
}
