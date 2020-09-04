package dynamicprogramming

import (
	"unicode"
)

// You can perform the following operations on the string, :

// Capitalize zero or more of a's lowercase letters.
// Delete all of the remaining lowercase letters in a.
// Given two strings, a and b, determine if it's possible to make a equal to b as described. If so, print YES on a new line. Otherwise, print NO.

// For example, given a = AbcDE and b = ABDE, in a we can convert b and delete c to match b. If a = AbcDE and b = AFDE, matching is not possible because letters may only be capitalized or discarded, not changed.

// Function Description

// Complete the function abbreviation in the editor below. It must return either YES or NO.

// abbreviation has the following parameter(s):

// a: the string to modify
// b: the string to match

//Note:
//String a consists only of uppercase and lowercase English letters, ascii[A-Za-z].
//String b consists only of uppercase English letters, ascii[A-Z].
//
//more info:https://www.hackerrank.com/challenges/abbr/

const yes = "YES"
const no = "NO"

func abbreviation(a string, b string) string {
	// fmt.Println("a is", a, "b is", b)
	if a == b {
		return yes
	}
	memCacheAbbr := make([][]int, 0)
	for i := 0; i <= len(a); i++ {
		row := make([]int, len(b)+1)
		for j := 0; j <= len(b); j++ {
			row[j] = -1
		}
		memCacheAbbr = append(memCacheAbbr, row)
	}
	//fmt.Println("memCache:", memCacheAbbr)
	if abbreviationHelper(memCacheAbbr, a, 0, b, 0) {
		return yes
	}
	return no

}

func abbreviationHelper(memCacheAbbr [][]int, a string, aindex int, b string, bindex int) bool {
	// fmt.Println("looking at aindex", aindex, "bindex", bindex)
	if memCacheAbbr[aindex][bindex] != -1 {
		if memCacheAbbr[aindex][bindex] == 1 {
			return true
		}

		return false
	}
	alen := len(a)
	blen := len(b)
	if aindex >= alen && bindex >= blen {
		setMemCache(memCacheAbbr, aindex, bindex, true)
		return true
	}
	if aindex >= alen {
		setMemCache(memCacheAbbr, aindex, bindex, false)
		return false
	}

	var possible bool
	if bindex >= blen {
		if unicode.IsLower(rune(a[aindex])) {
			possible = abbreviationHelper(memCacheAbbr, a, aindex+1, b, bindex)
			setMemCache(memCacheAbbr, aindex, bindex, possible)
			return possible
		}
		setMemCache(memCacheAbbr, aindex, bindex, false)
		return false
	}
	// fmt.Println("aindex is:", string(a[aindex]), "bindex is:", string(b[bindex]))

	if a[aindex] == b[bindex] {
		possible = abbreviationHelper(memCacheAbbr, a, aindex+1, b, bindex+1)
		setMemCache(memCacheAbbr, aindex, bindex, possible)
		return possible
	}

	if unicode.ToUpper(rune(a[aindex])) == rune(b[bindex]) {
		possible = abbreviationHelper(memCacheAbbr, a, aindex+1, b, bindex) || abbreviationHelper(memCacheAbbr, a, aindex+1, b, bindex+1)
		setMemCache(memCacheAbbr, aindex, bindex, possible)
		return possible
	}

	if unicode.IsUpper(rune(a[aindex])) && a[aindex] != b[bindex] {
		setMemCache(memCacheAbbr, aindex, bindex, false)
		return false
	}

	possible = abbreviationHelper(memCacheAbbr, a, aindex+1, b, bindex)
	setMemCache(memCacheAbbr, aindex, bindex, possible)
	return possible

}

func setMemCache(memCache [][]int, aindex, bindex int, settrue bool) {
	if settrue {
		memCache[aindex][bindex] = 1
	} else {
		memCache[aindex][bindex] = 0
	}
}
