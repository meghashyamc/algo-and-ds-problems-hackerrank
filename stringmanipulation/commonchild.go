package stringmanipulation

// A string is said to be a child of a another string if it can be formed by deleting 0 or more characters from the other string. Given two strings of equal length, what's the longest string that can be constructed such that it is a child of both?

// For example, ABCD and ABDC have two children with maximum length 3, ABC and ABD. They can be formed by eliminating either the D or C from both strings. Note that we will not consider ABCD as a common child because we can't rearrange characters and ABCD  ABDC.

// Function Description

// Complete the commonChild function in the editor below. It should return the longest string which is a common child of the input strings.

// commonChild has the following parameter(s):

// s1, s2: two equal length strings

//cache stores string1, string2 and max size of common child for the two strings
var memCache [][]int

func CommonChild(s1 string, s2 string) int32 {

	if s1 == s2 {
		return int32(len(s1))
	}
	if len(s1) == 1 {
		return 0
	}
	memCache = make([][]int, 0)
	for i := 0; i < len(s1)+1; i++ {
		memCache = append(memCache, make([]int, len(s1)+1))
	}
	populateCache()

	return int32(commonChildHelper(&s1, 0, &s2, 0))
}

func commonChildHelper(s1 *string, i1 int, s2 *string, i2 int) int {

	n := len(*s1)

	if memCache[i1][i2] != -1 {
		return memCache[i1][i2]
	}

	if i1 >= n || i2 >= n {
		return 0
	}

	if (*s1)[i1] == (*s2)[i2] {
		memCache[i1][i2] = 1 + commonChildHelper(s1, i1+1, s2, i2+1)

	} else {
		memCache[i1][i2] = max(commonChildHelper(s1, i1, s2, i2+1), commonChildHelper(s1, i1+1, s2, i2))
	}
	return memCache[i1][i2]
}

func populateCache() {

	for i := 0; i < len(memCache); i++ {
		for j := 0; j < len(memCache[0]); j++ {
			memCache[i][j] = -1
		}
	}
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
