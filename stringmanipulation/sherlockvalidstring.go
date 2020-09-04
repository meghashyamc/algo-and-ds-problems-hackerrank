package stringmanipulation

// Sherlock considers a string to be valid if all characters of the string appear the same number of times. It is also valid if he can remove just 1 character at 1 index in the string, and the remaining characters will occur the same number of times. Given a string , determine if it is valid. If so, return YES, otherwise return NO.

// For example, if s=abc, it is a valid string because frequencies are {a:1,b:1,c:1}. So is s=abcc because we can remove one c and have 1 of each character in the remaining string. If  however, the string s = abccc is not valid as we can only remove 1 occurrence of c. That would leave character frequencies of {a:1,b:1,c:2}.

// Function Description

// Complete the isValid function in the editor below. It should return either the string YES or the string NO.

// isValid has the following parameter(s):

// s: a string

func IsValid(s string) string {

	if len(s) <= 2 {
		return "YES"
	}

	frequenciesOfChars := make(map[rune]int)

	//find frequencies of each character
	for _, ch := range s {
		frequenciesOfChars[ch]++
	}
	// fmt.Println("frequenciesOfChars", frequenciesOfChars)

	charsPerFrequency := make(map[int]int)

	//find number of characters for each frequency
	for _, freq := range frequenciesOfChars {
		charsPerFrequency[freq]++
	}
	// fmt.Println("charsPerFrequency", charsPerFrequency)

	//there can't be more than two possible frequencies
	if len(charsPerFrequency) > 2 {
		return "NO"
	}

	if len(charsPerFrequency) == 1 {
		return "YES"
	}

	maxNumOfChars := 0
	mostCommonFreq := 0
	for freq, numOfChars := range charsPerFrequency {

		if numOfChars > maxNumOfChars {
			mostCommonFreq = freq
			maxNumOfChars = numOfChars
		}

	}

	// fmt.Println("maxNumOfChars:", maxNumOfChars)
	// fmt.Println("mostCommonFreq:", mostCommonFreq)
	for freq, numOfChars := range charsPerFrequency {

		//number of characters with a different frequency must be 1
		//AND the different frequency should be 1 more than the common frequency or 1
		if freq != mostCommonFreq {
			if numOfChars != 1 {
				return "NO"
			}
			if !(freq == mostCommonFreq+1 || freq == 1) {
				return "NO"
			}
		}

	}

	return "YES"
}
