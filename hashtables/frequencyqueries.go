package hashtables

// You are given q queries. Each query is of the form two integers described below:
// -1 x: Insert x in your data structure.
// -2 y: Delete one occurence of y from your data structure, if present.
// -3 z: Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.

// The queries are given in the form of a 2-D array queries  of size q where queries[i][0] contains the operation, and queries[i][1] contains the data element.
// For example, you are given the array queries = [(1,1),(2,2), (3,2), (1,1), (1,1), (2,1), (3,2)] . The results of each operation are:

// Operation   Array   Output
// (1,1)       [1]
// (2,2)       [1]
// (3,2)                   0
// (1,1)       [1,1]
// (1,1)       [1,1,1]
// (2,1)       [1,1]
// (3,2)                   1
// Return an array with the output: [0,1]

// Function Description

// Complete the freqQuery function in the editor below. It must return an array of integers where each element is a 1 if there is at least one element value with the queried number of occurrences in the current array, or 0 if there is not.

// freqQuery has the following parameter(s):

// queries: a 2-d array of integers

const (
	addquery        = 1
	removequery     = 2
	searchFreqQuery = 3
)

func freqQuery(queries [][]int32) []int32 {

	//key = number, value = its frequency
	queryFrequencies := make(map[int32]int32)

	//key = frequency, value = number of numbers with that frequency
	frequencyList := make(map[int32]int32)

	freqQueryResult := make([]int32, 0)

	for _, query := range queries {

		if len(query) < 2 {
			panic("error in length of input parameter: queries")
		}
		switch query[0] {
		case addquery:
			oldFreq := queryFrequencies[query[1]]
			if oldFreq != 0 {
				if frequencyList[oldFreq] == 1 {
					delete(frequencyList, oldFreq)
				} else {
					frequencyList[oldFreq]--
				}
			}
			queryFrequencies[query[1]]++
			frequencyList[oldFreq+1]++

		case removequery:
			oldFreq := queryFrequencies[query[1]]
			if oldFreq == 0 {
				break
			}
			if frequencyList[oldFreq] == 1 {
				delete(frequencyList, oldFreq)
			} else {
				frequencyList[oldFreq]--
			}
			if oldFreq == 1 {
				delete(queryFrequencies, query[1])
			} else {
				queryFrequencies[query[1]]--
				frequencyList[oldFreq-1]++
			}

		case searchFreqQuery:

			if _, ok := frequencyList[query[1]]; !ok {
				freqQueryResult = append(freqQueryResult, 0)
			} else {
				freqQueryResult = append(freqQueryResult, 1)
			}
		}
	}

	return freqQueryResult
}
