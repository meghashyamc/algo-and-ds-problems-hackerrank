package search

import (
	"sort"
)

// Given 3 arrays a,b,c of different sizes, find the number of distinct triplets (p,q,r) where p is an element of a, written as p belongs to q, q belongs to b, and r belongs to c, satisfying the criteria: .
// p <=q and q >= r
// For example, given a = [3,5,7], b = [3,6] and c = [4,6,9], we find four distinct triplets: .
// (3,6,4), (3,6,6), (5,6,4), (5,6,6)
// Function Description

// Complete the triplets function in the editor below. It must return the number of distinct triplets that can be formed from the given arrays.

// triplets has the following parameter(s):

// a, b, c: three arrays of integers .
//more info: https://www.hackerrank.com/challenges/triple-sum

type array []int32

func triplets(a []int32, b []int32, c []int32) int64 {

	//sort a,b,c in descending order
	sort.Slice(a, array(a).less)
	sort.Slice(b, array(b).less)
	sort.Slice(c, array(c).less)

	na := len(a)
	nb := len(b)
	nc := len(c)

	//get how many distinct numbers are there after every index of a and c
	numsAfterIndexA := getNumsAfterIndex(a)
	numsAfterIndexC := getNumsAfterIndex(c)

	var tripletcount int64

	for i, j, k := 0, 0, 0; i < na && j < nb && k < nc; {

		//triplet found
		if a[i] <= b[j] && b[j] >= c[k] {
			//as a and c are reverse sorted, a[i] <= b[j] means, a[i+1], a[i+2] etc. are all <= b[j]
			//similar logic for c
			tripletcount += int64(numsAfterIndexA[i]) * int64(numsAfterIndexC[k])
			for {
				//consider same numbers like 3,3,3...once
				//if we find a triplet, we need to only increment b's index because all other indices of
				//a and c already fulfill our condition for this index of b
				if j == nb-1 || b[j] != b[j+1] {
					j++
					break
				}
				j++
			}
			continue
		}
		//a is the problem, so increment a's index
		if a[i] > b[j] && b[j] >= c[k] {
			i++
			//b is the problem, so increment b's index
		} else if a[i] <= b[j] && b[j] < c[k] {
			k++
			//both a and b are not fulfilling the condition,
			//increment indices for both
		} else {
			i++
			k++
		}

	}
	return tripletcount
}

func (a array) less(i, j int) bool {

	return a[i] > a[j]
}

//takes in a reverse sorted slice
//returns a slice that shows how many distinct numbers come in arr after index i (including i)
func getNumsAfterIndex(arr []int32) []int32 {

	if arr == nil {
		return nil
	}
	numsAfterIndex := make([]int32, 0)
	var distinctCount int32
	for i, num := range arr {

		if i < len(arr)-1 {
			if num != arr[i+1] {
				distinctCount++
			}
		} else {

			if len(arr) == 1 || num != arr[i-1] {
				distinctCount++
			}
		}
	}

	for i, num := range arr {

		if i != 0 && num != arr[i-1] {
			distinctCount--
		}
		numsAfterIndex = append(numsAfterIndex, distinctCount)

	}
	return numsAfterIndex
}
