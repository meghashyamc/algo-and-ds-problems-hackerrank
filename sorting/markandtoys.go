package sorting

import "sort"

// Mark and Jane are very happy after having their first child. Their son loves toys, so Mark wants to buy some. There are a number of different toys lying in front of him, tagged with their prices.
// Mark has only a certain amount to spend, and he wants to maximize the number of toys he buys with this money.

// Given a list of prices and an amount to spend,
//what is the maximum number of toys Mark can buy?
//For example, if  prices = [1,2,3] and Mark has  to spend k = 7,
//he can buy items [1,2,3] for 6 , or [3,4] for 7 units of currency.
//He would choose the first group of 7 items.

// Function Description

// Complete the function maximumToys in the editor below. It should return an integer representing the maximum number of toys Mark can purchase.

// maximumToys has the following parameter(s):

// prices: an array of integers representing toy prices
// k: an integer, Mark's budget
func maximumToys(prices []int32, k int32) int32 {

	sort.Slice(prices, func(i, j int) bool {

		return prices[i] < prices[j]
	})

	var numOfToys int32
	var costTillNow int32
	for _, price := range prices {

		costTillNow += price
		if costTillNow > k {
			return numOfToys
		}
		numOfToys++
	}

	return numOfToys

}
