package sorting

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
// constraints:
// 1 <= n <= 2 x 10^5
// 1 <= d <= n
// 0 <= expenditure[i] <= 200

const maxexpense = 200

func ActivityNotifications(expenditure []int32, d int32) int32 {

	countSortArr := make([]int32, maxexpense+1)

	for _, dayExpense := range expenditure[0:d] {

		countSortArr[dayExpense]++
	}

	var notificationCount int32
	for i, dayExpense := range expenditure[d:len(expenditure)] {

		if float32(dayExpense) >= getMedianCountSort(countSortArr, d)*2 {
			notificationCount++
		}

		countSortArr[dayExpense]++
		countSortArr[expenditure[i]]--
	}
	return notificationCount

}

//getMedian receives a counting sort array
//n is the total number of elements for which we want a median
func getMedianCountSort(arr []int32, n int32) float32 {

	var countTillNow int32
	index := 0
	var count int32
	var odd bool
	if n%2 != 0 {
		odd = true
	}
	for index, count = range arr {

		countTillNow += count
		if odd && countTillNow >= n/2+1 {
			break
		}
		if !odd && countTillNow >= n/2 {
			break
		}
	}

	if odd {
		return float32(index)
	}

	//reaching here means n is even
	if countTillNow >= n/2+1 {
		return float32(index)
	}

	newindex := 0
	for newindex, count = range arr[index+1 : len(arr)] {
		countTillNow += count
		if countTillNow >= n/2+1 {
			break
		}
	}

	return (float32(index) + float32(index+newindex+1)) / 2.0

}
