package arrays

import (
	"fmt"
)

/*
It's New Year's Day and everyone's in line for the Wonderland rollercoaster ride! There are a
number of people queued up, and each person wears a sticker indicating their initial position
 in the queue. Initial positions increment by  from  at the front of the line to  at the back.

Any person in the queue can bribe the person directly in front of them to swap positions. If
two people swap positions, they still wear the same sticker denoting their original places in
line. One person can bribe at most two others. For example, if n = 8  and person 5  bribes person 4
the queue will look like this: 1,2,3,5,4,6,7,8

Fascinated by this chaotic queue, you decide you must know the minimum number of bribes that
 took place to get the queue into its current state!

Function Description

Complete the function minimumBribes below. It must print an integer representing the minimum number of bribes necessary, or Too chaotic if the line configuration is not possible.

minimumBribes has the following parameter(s):

q: an array of integers

Sample input:
2
5
2 1 5 3 4
5
2 5 1 3 4
Sample output:
3
Too chaotic

*/

func minimumBribes(q []int32) {

	var numOfPpl int32 = int32(len(q))
	var origArr = make([]int32, 0)

	var i int32
	for i = 0; i < numOfPpl; i++ {
		origArr = append(origArr, i+1)
	}

	finalPosMap := make(map[int32]int32)
	posToGoMap := make(map[int32]int32)
	bribesMap := make(map[int32]uint8)

	// sets final and original position mapping like so: orign pos -> final pos
	fillFinalPosMap(finalPosMap, q)

	// sets initial bribers for each person to 0
	fillBribesMap(bribesMap, numOfPpl)

	// sets number of positions to go to move from original to final array like so: orig pos -> positions to move
	// (with move to right being positive, left being negative)
	fillPosToGoMap(posToGoMap, finalPosMap, origArr)

	var j int32 = 0
	for i = 0; i < numOfPpl; {
		// fmt.Println("i is", i)
		// fmt.Println("origArr is", origArr)
		// fmt.Println("posToGoMap is", posToGoMap)
		// fmt.Println("bribesMap is", bribesMap)

		if posToGoMap[origArr[i]] == int32(0) {
			i++
			continue
		}

		if i < numOfPpl-1 && posToGoMap[origArr[i]] <= posToGoMap[origArr[i+1]] {

			j = i
			i++
			for i < numOfPpl-1 && posToGoMap[origArr[i]] <= posToGoMap[origArr[i+1]] {

				i++

			}
			continue
		}

		if i < numOfPpl-1 && posToGoMap[origArr[i]] > posToGoMap[origArr[i+1]] {

			posToGoMap[origArr[i]]--
			posToGoMap[origArr[i+1]]++

			bribesMap[origArr[i+1]]++
			if bribesMap[origArr[i+1]] > 2 {
				fmt.Println("Too chaotic")
				return
			}

			swap(&origArr, i, i+1)

			i = j

			continue

		}
	}

	minBribes := getMinBribes(bribesMap)
	fmt.Println(minBribes)
}

func fillFinalPosMap(finalPosMap map[int32]int32, q []int32) {

	for j, pos := range q {
		finalPosMap[pos] = int32(j + 1)
	}
}

func fillBribesMap(bribesMap map[int32]uint8, numOfPpl int32) {

	var i int32
	for i = 0; i < numOfPpl; i++ {
		bribesMap[i+1] = uint8(0)
	}
}

func fillPosToGoMap(posToGoMap map[int32]int32, finalPosMap map[int32]int32, origArr []int32) {

	for _, index := range origArr {
		posToGoMap[index] = finalPosMap[index] - index
	}
}

// func swap(sl *([]int32), i int32, j int32) {

// 	fmt.Println("swap function called with i", i, " and j", j)
// 	var temp int32
// 	temp = (*sl)[i]
// 	fmt.Println("temp is", temp)
// 	(*sl)[i] = (*sl)[j]
// 	fmt.Println("sl[i] is", (*sl)[i])
// 	(*sl)[j] = temp
// 	fmt.Println("sl[j] is", (*sl)[j])

// 	fmt.Println("inside the swapping function, sl is", *sl)
// }

func getMinBribes(bribesMap map[int32]uint8) int32 {

	var minBribes int32 = 0

	for _, bribes := range bribesMap {

		minBribes += int32(bribes)
	}
	return minBribes

}
