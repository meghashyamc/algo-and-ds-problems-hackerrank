package greedy

// Given a string,A, we define some operations on the string as follows:

// a. reverse(A) denotes the string obtained by reversing string A. Example: reverse("abc") = "cba"

// b.shuffle(A) denotes any string that's a permutation of string A. Example:

//shuffle("god") belongs to ["god", "gdo", "ogd", "odg", "dgo", "dog"]
// c. merge(A) denotes any string that's obtained by interspersing the two strings A1 & A2, maintaining the order of characters in both. For example, A1 = "abc" & A2 = "def", one possible result of merge(A1,A2) could be "abcdef", another could be "abdecf", another could be "adbecf" and so on.

// Given a string s such that s belongs to merge(reverse(A), shuffle(A)) for some string A, find the lexicographically smallest A .

// For example, s = abab. We can split it into two strings of ab. The reverse is ba and we need to find a string to shuffle in to get abab. The middle two characters match our reverse string, leaving the a and b at the ends. Our shuffle string needs to be ab.
// Lexicographically ab < ba, so our answer is ab.

// Function Description

// Complete the reverseShuffleMerge function in the editor below. It must return the lexicographically smallest string fitting the criteria.

// reverseShuffleMerge has the following parameter(s):
// s: a string
//more details here: https://www.hackerrank.com/challenges/reverse-shuffle-merge/

func reverseShuffleMerge(s string) string {

	charFrequencies := make(map[byte]int)

	//get the frequency of each character in s
	for _, ch := range []byte(s) {
		charFrequencies[ch]++
	}

	//get the frequency of each character in A (the answer string)
	for key, val := range charFrequencies {

		charFrequencies[key] = val / 2
	}

	answerA := make([]byte, 0)
	charsDone := make(map[byte]int)
	ignoredChars := make(map[byte]int)

	for i := len(s) - 1; i >= 0; i-- {
		// fmt.Println("looking at:", string(s[i]))
		// fmt.Println("charsDone is", charsDone)

		if charsDone[s[i]] == charFrequencies[s[i]] {
			continue
		}

		// fmt.Println("i is", i, "character is", string(s[i]))
		if isLexicographicallySmallest(charsDone, charFrequencies, s[i]) || !canBeIgnored(charsDone, charFrequencies, ignoredChars, &s, i) {
			// fmt.Println("added", string(s[i]))
			answerA = append(answerA, s[i])
			// fmt.Println("answerA is", string(answerA))
			charsDone[s[i]]++

		} else {
			// fmt.Println("ignored", string(s[i]))
			ignoredChars[s[i]]++

		}

		if len(answerA) == len(s)/2 {
			break
		}

	}
	return string(answerA)

}

//is the character given the lexicographically smallest character we can add?
func isLexicographicallySmallest(charsDone, charFrequencies map[byte]int, b byte) bool {

	for ch, freq := range charFrequencies {

		//ignore lexicographically greater characters

		if ch >= b {
			// fmt.Println(string(ch), ">=", string(rune(b)))
			continue
		}
		//there exists a lexicographically smaller character that has not yet been added
		if charsDone[ch] < freq {
			// fmt.Println(string(ch), "<", string(rune(b)), "and done frequency", charsDone[b], "less than required freq", freq)

			return false
		}
	}
	return true
}

//a character must not be ignored if there aren't enough characters that can be encountered later
//a character must not be ignored if the next critical character is lexicographicaly larger than this character
func canBeIgnored(charsDone, charFrequencies, ignoredChars map[byte]int, s *string, i int) bool {

	if charFrequencies[(*s)[i]] <= ignoredChars[(*s)[i]] {
		// fmt.Println("found in critical list")
		return false
	}

	//this map keeps track of characters we iterate over
	// it is required so that we take into account characters that become critical as we iterate
	tempcharFreqMap := make(map[byte]int)
	for j := i - 1; j >= 0; j-- {

		if charsDone[(*s)[j]] != charFrequencies[(*s)[j]] {
			// fmt.Println("tempfreqmap:", tempcharFreqMap, "charfrequencies:", charFrequencies)
			if (*s)[j] < (*s)[i] {
				// fmt.Println("smaller char seen:", string((*s)[j]))
				return true
			}

			if (*s)[j] == (*s)[i] {
				continue
			}

			// fmt.Println("s[j] is ", (*s)[j], "charFrequencies[(*s)[j]]:", charFrequencies[(*s)[j]], "tempcharFreqMap[(*s)[j]]:", tempcharFreqMap[(*s)[j]])
			if charFrequencies[(*s)[j]]-tempcharFreqMap[(*s)[j]] <= ignoredChars[(*s)[j]] {
				// fmt.Println("found larger critical character:", string((*s)[j]))

				return false

			}
			tempcharFreqMap[(*s)[j]]++
		}

	}
	return true
}
