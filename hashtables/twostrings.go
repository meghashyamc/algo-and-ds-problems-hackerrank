package hashtables

// Given two strings, determine if they share a common substring. A substring may be as small as one character.

// For example, the words "a", "and", "art" share the common substring . The words "be" and "cat" do not share a substring.

// Function Description

// Complete the function twoStrings in the editor below. It should return a string, either YES or NO based on whether the strings share a common substring.

// twoStrings has the following parameter(s):

// s1, s2: two strings to analyze .

// Constraints
//s1 and s2 consist of characters in the range ascii[a-z].
//1 <= p <= 10
//1 <= |s1|,|s2|<=10^5

// Output Format

// For each pair of strings, return YES or NO.

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	chars := make(map[rune]bool)

	for _, ch := range s1 {
		chars[ch] = true
	}

	for _, ch := range s2 {
		if _, ok := chars[ch]; ok {
			return "YES"
		}
	}

	return "NO"

}
