package greedy

import (
	"sort"
)

// A group of friends want to buy a bouquet of flowers. The florist wants to maximize his number of new customers and the money he makes.
//To do this, he decides he'll multiply the price of each flower by the number of that customer's previously purchased flowers plus 1. The first flower will be original price, (0+1) x original price, the next will be (1+1) x original price and so on.

// Given the size of the group of friends, the number of flowers they want to purchase and the original prices of the flowers, determine the minimum cost to purchase all of the flowers.

// For example, if there are k = 3 friends that want to buy n = 4 flowers that cost c = [1,2,3,4] each will buy one of the flowers priced [2,3,4] at the original price.
// Having each purchased x = 1 flower, the first flower in the list,c[0], will now cost (1+1) x 1 = 2. The total cost will be 2+3+4+2 = 11.

// Function Description

// Complete the getMinimumCost function in the editor below. It should return the minimum cost to purchase all of the flowers.

// getMinimumCost has the following parameter(s):

// c: an array of integers representing the original price of each flower
// k: an integer, the number of friends

//More details and constraints: https://www.hackerrank.com/challenges/greedy-florist/problem

func getMinimumCost(k int32, c []int32) int32 {

	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})

	var previousFlowers int32
	var transactionsTillK int32
	var mintotalcost int32
	for _, cost := range c {
		transactionsTillK++
		if transactionsTillK == k+1 {
			transactionsTillK = 1
			previousFlowers++
		}
		mintotalcost += (previousFlowers + 1) * cost
		// fmt.Println("transactionsTillK:", transactionsTillK, "cost:", cost, "prevFlowers:", previousFlowers, "mintotalcost:", mintotalcost)
	}

	return mintotalcost

}
