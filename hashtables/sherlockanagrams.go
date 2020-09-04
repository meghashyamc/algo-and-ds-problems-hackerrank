package hashtables

// Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string. Given a string, find the number of pairs of substrings of the string that are anagrams of each other.

// For example, if s = mom, the list of all anagrammatic pairs is[m,m], [mo, om]  at positions [[0], [2], [0,1], [1,2]] respectively.

// Function Description

// Complete the function sherlockAndAnagrams in the editor below. It must return an integer that represents the number of anagrammatic pairs of substrings in .

// sherlockAndAnagrams has the following parameter(s):

// s: a string .

func sherlockAndAnagrams(s string) int32 {

	var numberofanagrams int32

	for substrlen := 1; substrlen < len(s); substrlen++ {

		thislensubstrings := make(map[string]int32)
		for i := 0; i < len(s)+1-substrlen; i++ {
			substr := s[i : i+substrlen]
			// fmt.Println("will check substring:", substr)
			val, ok := thislensubstrings[substr]
			if ok {
				thislensubstrings[substr] = val + 1
				continue
			}
			// fmt.Println("findAnagramInList was sent: thislensubstrings:", thislensubstrings, "substr:", substr)
			anagramSubStr := findAnagramInList(thislensubstrings, substr)
			// fmt.Println("anagramSubStr returned was:", anagramSubStr)
			if anagramSubStr != "" {
				thislensubstrings[anagramSubStr]++
			} else {
				thislensubstrings[substr] = 0
			}
		}
		for _, val := range thislensubstrings {

			numberofanagrams += (val * (val + 1)) / 2
		}
		// fmt.Println("substrings of length", substrlen, "are------------------------------------>", thislensubstrings)

	}
	return numberofanagrams
}

func findAnagramInList(thislensubstrings map[string]int32, substr string) string {

	for keystring := range thislensubstrings {

		if isAnagram(substr, keystring) {
			return keystring
		}

	}
	return ""

}

func isAnagram(s1, s2 string) bool {

	chars := make(map[rune]int32)

	for _, ch := range s1 {

		freq, ok := chars[ch]
		if !ok {
			chars[ch] = 1
		}
		chars[ch] = freq + 1
	}

	for _, ch := range s2 {
		freq, ok := chars[ch]
		if !ok {
			return false
		}
		if freq == 1 {
			delete(chars, ch)
		} else {
			chars[ch] = freq - 1
		}
	}
	if len(chars) != 0 {
		return false
	}
	return true
}
