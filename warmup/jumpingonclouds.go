package warmup

/*
Emma is playing a new mobile game that starts
with consecutively numbered clouds. Some of the
clouds are thunderheads and others are cumulus.
She can jump on any cumulus cloud having a numbe
r that is equal to the number of the current cloud
plus 1 or 2.  She must avoid the thunderheads.
Determine the minimum number of jumps it will take
Emma to jump from her starting postion to the last cloud.
It is always possible to win the game.

Function Description

Write the jumpingOnClouds function below.
 It should return the minimum number of jumps
  required, as an integer.

jumpingOnClouds has the following parameter(s):

c: an array of binary integers

*/

import ()

const safe = 0
const unsafe = 1

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	n := len(c)

	if n <= 1 {
		return 0
	}

	if n == 2 {
		return 1
	}
	currIndex := 0
	steps := 0

	for {

		if currIndex >= n-1 {
			break
		}

		if currIndex < n-2 {

			if c[currIndex+2] == safe {
				currIndex += 2
				steps++
				continue
			}

		}

		if c[currIndex+1] == safe {

			currIndex++
			steps++

		}

	}

	return int32(steps)

}
