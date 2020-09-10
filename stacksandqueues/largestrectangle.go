package stacksandqueues

// Skyline Real Estate Developers is planning to demolish a number of old, unoccupied buildings and construct a shopping mall in their place. Your task is to find the largest solid area in which the mall can be constructed.

// There are a number of buildings in a certain two-dimensional landscape. Each building has a height h[i] where i belongs to [i,n]. If you join k adjacent buildings, they will form a solid rectangle of area k x min(h[i], h[i+1],..., h[i+k-1]).

// For example, the heights array h = [3,2,3]. A rectangle of height h and length k = 3 can be constructed within the boundaries. The area formed is h*k = 2*3 = 6.

// Function Description

// Complete the function largestRectangle int the editor below. It should return an integer representing the largest rectangle that can be formed within the bounds of consecutive buildings.

// largestRectangle has the following parameter(s):

// h: an array of integers representing building heights
//more info: https://www.hackerrank.com/challenges/largest-rectangle

func largestRectangle(h []int32) int64 {

	var maxRectangleArea int64
	for i, height := range h {
		var thisRectangleArea = int64(height)
		//add to this rectangle's area on the left
		for l := 0; l < i; l++ {
			if h[i-l-1] < height {
				break
			}
			thisRectangleArea += int64(height)
		}

		// add to this rectangle's area on the right
		for r := 0; r < len(h)-i-1; r++ {
			if h[i+r+1] < height {
				break
			}
			thisRectangleArea += int64(height)
		}
		if thisRectangleArea > maxRectangleArea {
			maxRectangleArea = thisRectangleArea
		}
	}

	return maxRectangleArea

}
