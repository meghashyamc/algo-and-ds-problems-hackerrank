package hashtables

// You are given an array and you need to find number of tripets of indices (i,j,k)  such that the elements at those indices are in geometric progression for a given common ratio r and i < j < k.

// For example,arr=[1,4,16,64]. If r=4, we have [1,4,16] and[4,16,64]  at indices (0,1,2) and (1,2,3).

// Function Description

// Complete the countTriplets function in the editor below. It should return the number of triplets forming a geometric progression for a given  as an integer.

// countTriplets has the following parameter(s):

// arr: an array of integers
// r: an integer, the common ratio

func CountTriplets1(arr []int64, r int64) int64 {

	var tripletcount int64
	dividends := make(map[int]([]int))

	for i := 0; i <= len(arr)-2; i++ {

		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] == 0 {
				continue
			}
			// fmt.Println("checking arr[j]/arr[i]:", arr[j]/arr[i], ", j is", j, "and i is", i)
			if arr[j]/arr[i] == r {

				val, ok := dividends[i]
				if !ok {
					val = make([]int, 0)
					val = append(val, j)
					dividends[i] = val
					continue
				}
				val = append(val, j)
				dividends[i] = val
			}

		}
	}
	// fmt.Println("dividends is...", dividends)
	for _, val := range dividends {

		for _, subkey := range val {

			subval, ok := dividends[subkey]

			if !ok {
				continue
			}
			tripletcount += int64(len(subval))
		}
	}

	return tripletcount
}
