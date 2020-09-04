package warmup

/*
Lilah has a string,s, of lowercase English
 letters that she repeated infinitely many times.
 Given an integer,n, find and print the number of letter a's
 in the first  letters of Lilah's infinite string.

 Function Description

Write the repeatedString function below.
It should return an integer representing the
number of occurrences of a in the prefix of length
  in the infinitely repeating string.

repeatedString has the following parameter(s):

s: a string to repeat
n: the number of characters to consider

*/
import ()

const charToFind = 'a'

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {

	if n <= 0 {
		return 0
	}
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 && s[0] != charToFind {
		return 0
	}

	timesInSingleStr := getTimesInSingleStr(s, charToFind)

	timesStrContainedInN := n / int64(len(s))

	remainder := n % (int64(len(s)))

	timesInRemainder := getTimesInSingleStr(s[0:remainder], charToFind)

	timesInSingleStrInt64 := int64(timesInSingleStr)
	return timesInSingleStrInt64*timesStrContainedInN + int64(timesInRemainder)

}

func getTimesInSingleStr(s string, charToFind rune) int {
	timesInSingleStr := 0
	for _, ch := range s {

		if ch == charToFind {
			timesInSingleStr++
		}
	}

	return timesInSingleStr

}
