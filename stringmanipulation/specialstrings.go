package stringmanipulation

// A string is said to be a special string if either of two conditions is met:

// All of the characters are the same, e.g. aaa.
// All characters except the middle one are the same, e.g. aadaa.
// A special substring is any substring of a string which meets one of those criteria. Given a string, determine how many special substrings can be formed from it.

// For example, given the string s = mnonopoo, we have the following special substrings: .
//{m,n,o,n,o,p,o,o,non,ono,opo,oo}
// Function Description

// Complete the substrCount function in the editor below. It should return an integer representing the number of special substrings that can be formed from the given string.

// substrCount has the following parameter(s):

// n: an integer, the length of string s
// s: a string

type charRepeatsObj struct {
	char    rune
	repeats int
}

func SubstrCount(n int32, s string) int64 {

	charRepeatsArr := make([]charRepeatsObj, 0)
	//populate character and number of repeats array
	charRepeatsArrIndex := -1
	for i, ch := range s {
		if i > 0 && s[i] == s[i-1] {
			charRepeatsArr[charRepeatsArrIndex].repeats++
			continue
		}
		charRepeatsArr = append(charRepeatsArr, charRepeatsObj{char: ch, repeats: 1})
		charRepeatsArrIndex++
	}

	specialSubStrCount := 0
	//find number of special strings that have all characters the same

	for _, charRepeats := range charRepeatsArr {

		specialSubStrCount += charRepeats.repeats * (charRepeats.repeats + 1) / 2
	}

	//find number of special strings that have all characters the same except the middle character
	for i, charRepeats := range charRepeatsArr {

		if i == 0 || i == len(charRepeatsArr)-1 {
			continue
		}

		if charRepeats.repeats == 1 && charRepeatsArr[i-1].char == charRepeatsArr[i+1].char {

			specialSubStrCount += min(charRepeatsArr[i-1].repeats, charRepeatsArr[i+1].repeats)
		}

	}
	return int64(specialSubStrCount)

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
