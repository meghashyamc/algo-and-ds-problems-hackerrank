package hashtables

// You are given a sorted array and you need to find number of tripets of indices (i,j,k)  such that the elements at those indices are in geometric progression for a given common ratio r and i < j < k.

// For example,arr=[1,4,16,64]. If r=4, we have [1,4,16] and[4,16,64]  at indices (0,1,2) and (1,2,3).

// Function Description

// Complete the countTriplets function in the editor below. It should return the number of triplets forming a geometric progression for a given  as an integer.

// countTriplets has the following parameter(s):

// arr: an array of integers
// r: an integer, the common ratio

func countTripletsPreSorted(arr []int64, r int64) int64 {

	var tripletsCount int64
	frequencies := make(map[int64]int64)
	//record frequency of element values in a map
	for _, el := range arr {
		if _, ok := frequencies[el]; !ok {
			frequencies[el] = 1
			continue
		}
		frequencies[el]++
	}
	// fmt.Println("frequencies:", frequencies)

	if r != 1 {

		for _, el := range arr {

			freqSecondTerm, ok := frequencies[el*r]
			if !ok {
				continue
			}
			freqThirdTerm, ok := frequencies[el*r*r]
			if !ok {
				continue
			}
			tripletsCount += freqSecondTerm * freqThirdTerm
		}
	} else {

		for _, freq := range frequencies {

			if freq < 3 {
				continue
			}

			tripletsCount += freq * (freq - 1) * (freq - 2) / 6
			continue

		}
	}
	return tripletsCount
}
