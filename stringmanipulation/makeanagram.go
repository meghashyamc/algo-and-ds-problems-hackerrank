package stringmanipulation

// Alice is taking a cryptography class and finding anagrams to be very useful. We consider two strings to be anagrams of each other if the first string's letters can be rearranged to form the second string. In other words, both strings must contain the same exact letters in the same exact frequency For example, bacdc and dcbac are anagrams, but bacdc and dcbad are not.

// Alice decides on an encryption scheme involving two large strings where encryption is dependent on the minimum number of character deletions required to make the two strings anagrams. Can you help her find this number?

// Given two strings, a and b, that may or may not be of the same length, determine the minimum number of character deletions required to make a and b  anagrams. Any characters can be deleted from either of the strings.

// For example, if a = cde and b = dcf, we can delete e from string a and f from string b so that both remaining strings are cd and dc which are anagrams.

// Function Description

// Complete the makeAnagram function in the editor below. It must return an integer representing the minimum total characters that must be deleted to make the strings anagrams.

// makeAnagram has the following parameter(s):

// a: a string
// b: a string

func makeAnagram(a string, b string) int32 {

	characterFreqA := make(map[rune]int32)
	characterFreqB := make(map[rune]int32)

	for _, ch := range a {
		characterFreqA[ch]++
	}
	// fmt.Println("characterFreqA", characterFreqA)

	for _, ch := range b {
		characterFreqB[ch]++
	}
	// fmt.Println("characterFreqB", characterFreqB)

	var numberOfDeletions int32
	for ch, freq := range characterFreqA {
		if val, _ := characterFreqB[ch]; val != freq {
			// fmt.Println("number of deletions changed by ", mod(val-freq), "ch:", ch, "val:", val, "freq:", freq)
			numberOfDeletions += mod(val - freq)
		}
		delete(characterFreqB, ch)
	}

	for _, freq := range characterFreqB {
		numberOfDeletions += freq
	}

	return numberOfDeletions

}

func mod(i int32) int32 {

	if i >= 0 {
		return i
	}

	return -i
}
