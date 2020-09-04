package search

import (
	"fmt"
	"sort"
)

// Each time Sunny and Johnny take a trip to the Ice Cream Parlor, they pool their money to buy ice cream. On any given day, the parlor offers a line of flavors. Each flavor has a cost associated with it.

// Given the value of money and the cost of each flavor for  trips to the Ice Cream Parlor, help Sunny and Johnny choose two distinct flavors such that they spend their entire pool of money during each visit. ID numbers are the 1- based index number associated with a cost. For each trip to the parlor, print the ID numbers for the two types of ice cream that Sunny and Johnny purchase as two space-separated integers on a new line. You must print the smaller ID first and the larger ID second.

// For example, there are n=5 flavors having cost = [2,1,3,5,6]. Together they have money = 5 to spend. They would purchase flavor ID's 1 and 3 for a cost of 2+3=5. Use 1 based indexing for your response.

// Note:

// Two ice creams having unique IDs i and j may have the same cost.
// There will always be a unique solution.
// Function Description

// Complete the function whatFlavors in the editor below. It must determine the two flavors they will purchase and print them as two space-separated integers on a line.

// whatFlavors has the following parameter(s):

// cost: an array of integers representing price for a flavor
// money: an integer representing the amount of money they have to spend

//more info: https://www.hackerrank.com/challenges/ctci-ice-cream-parlor

func whatFlavors(cost []int32, money int32) {

	costindices := make(map[int32][]int)

	for i, singlecost := range cost {

		if _, ok := costindices[singlecost]; !ok {
			costindices[singlecost] = []int{i + 1}
			continue
		}
		costindices[singlecost] = append(costindices[singlecost], i+1)
	}
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] < cost[j]
	})

	n := len(cost)
	for i, j := 0, n-1; i < n && j >= 0 && j > i; {

		if cost[i]+cost[j] == money {
			answerIndices := getIndices(costindices, []int32{cost[i], cost[j]})
			if answerIndices == nil {
				fmt.Println("could not find indices with the required sum")
			}
			fmt.Println(answerIndices[0], answerIndices[1])
			return
		}

		if cost[i]+cost[j] < money {
			i++
			continue
		}
		j--
	}
}

func getIndices(costindices map[int32][]int, costs []int32) []int {

	var samecost bool
	if costs[0] == costs[1] {
		samecost = true
	}
	firstcostindices, ok := costindices[costs[0]]
	if !ok || len(firstcostindices) < 1 {
		return nil
	}

	if samecost {
		return firstcostindices
	}
	secondcostindices, ok := costindices[costs[1]]
	if !ok || len(secondcostindices) < 1 {
		return nil
	}

	answerIndices := []int{firstcostindices[0], secondcostindices[0]}
	sort.Ints(answerIndices)
	return answerIndices

}
