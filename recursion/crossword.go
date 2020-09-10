package recursion

import (
	"strings"
)

// A 10x10 Crossword grid is provided to you, along with a set of words (or names of places) which need to be filled into the grid. Cells are marked either + or -. Cells marked with a - are to be filled with the word list.

// The following shows an example crossword from the input crossword grid and the list of words to fit, :
// words = [POLAND, LHASA, SPAIN, INDIA]
// Input 	   		Output

// ++++++++++ 		++++++++++
// +------+++ 		+POLAND+++
// +++-++++++ 		+++H++++++
// +++-++++++ 		+++A++++++
// +++-----++ 		+++SPAIN++
// +++-++-+++ 		+++A++N+++
// ++++++-+++ 		++++++D+++
// ++++++-+++ 		++++++I+++
// ++++++-+++ 		++++++A+++
// ++++++++++ 		++++++++++
// POLAND;LHASA;SPAIN;INDIA

// Function Description

// Complete the crosswordPuzzle function in the editor below. It should return an array of strings, each representing a row of the finished puzzle.

// crosswordPuzzle has the following parameter(s):

// crossword: an array of 10 strings of length 10 representing the empty grid
// words: a string consisting of semicolon delimited strings to fit into crossword
// more info: https://www.hackerrank.com/challenges/crossword-puzzle

type blank struct {
	startx int
	starty int
	endx   int
	endy   int
	length int
}

const blankchar = '-'

func crosswordPuzzle(crossword []string, words string) []string {

	if len(crossword) == 0 || len(words) == 0 {
		return nil
	}

	wordsList := strings.Split(words, ";")
	wordsMap := convertToMap(wordsList)
	crosswordBytes := getCrosswordBytes(crossword)
	blanks := getBlanksToFill(crosswordBytes)
	if len(blanks) != len(wordsMap) {
		return nil
	}

	if fillWordsInBlanks(blanks, 0, wordsMap, crosswordBytes) {
		return convertToString(crosswordBytes)
	}

	return nil

}

//for each blank, consider all words, if the word fits, consider all remaining words for remaining blanks and so on
//if there is an issue in doing the above, backtrack and continue from the previous blank (ignore the latest word that did fit, and try other words not tried till now)
func fillWordsInBlanks(blanks []blank, blankIndex int, wordsMap map[string]bool, crosswordBytes [][]byte) bool {
	if blankIndex >= len(blanks) && len(wordsMap) == 0 {
		return true
	}
	if blankIndex >= len(blanks) {
		return false
	}
	//list of all words to be tried out
	for word := range wordsMap {
		crosswordBytesCopy := createNewCopy(crosswordBytes)
		//if word fits in blank
		if addWordToBlank(blanks[blankIndex], word, crosswordBytesCopy) {
			wordsMapCopy := make(map[string]bool)
			copymap(wordsMapCopy, wordsMap)
			//remove word that fit from future recursive calls
			delete(wordsMapCopy, word)
			//recursively try to fill remaining words in remaining blanks
			if fillWordsInBlanks(blanks, blankIndex+1, wordsMapCopy, crosswordBytesCopy) {
				copy(crosswordBytes, crosswordBytesCopy)
				return true
			}
			//if we faced a problem, act as if this word never fit, move on to other words
		}
	}
	// fmt.Println("false returned after loop")
	return false
}

func createNewCopy(arr [][]byte) [][]byte {

	if arr == nil {
		return nil
	}
	copyArr := make([][]byte, 0)
	if len(arr) == 0 {
		return copyArr
	}
	for i := 0; i < len(arr); i++ {
		partArr := make([]byte, 0)
		for j := 0; j < len(arr[i]); j++ {
			partArr = append(partArr, arr[i][j])
		}
		copyArr = append(copyArr, partArr)
	}
	return copyArr
}

func addWordToBlank(currblank blank, word string, crosswordBytes [][]byte) bool {

	if currblank.length != len(word) {
		return false
	}
	var isVerticalBlank bool
	if currblank.startx == currblank.endx {
		isVerticalBlank = true
	}

	if isVerticalBlank {
		startRow := currblank.starty
		endRow := currblank.endy
		fixedCol := currblank.startx
		if verifyVerticalFill(startRow, endRow, fixedCol, word, crosswordBytes) {
			for i, j := 0, startRow; i < len(word) && j < endRow; i, j = i+1, j+1 {
				crosswordBytes[j][fixedCol] = word[i]
			}
			return true
		}
		return false
	}

	startCol := currblank.startx
	endCol := currblank.endx
	fixedRow := currblank.starty
	if verifyHorizontalFill(startCol, endCol, fixedRow, word, crosswordBytes) {
		for i, j := 0, startCol; i < len(word) && j < endCol; i, j = i+1, j+1 {
			crosswordBytes[fixedRow][j] = word[i]
		}
		return true
	}
	return false
}

func verifyHorizontalFill(startCol, endCol, fixedRow int, word string, crosswordBytes [][]byte) bool {
	for i, j := 0, startCol; i < len(word) && j < endCol; i, j = i+1, j+1 {
		if crosswordBytes[fixedRow][j] == blankchar {
			continue
		}
		if crosswordBytes[fixedRow][j] == word[i] {
			continue
		}
		return false
	}
	return true
}
func verifyVerticalFill(startRow, endRow, fixedCol int, word string, crosswordBytes [][]byte) bool {

	for i, j := 0, startRow; i < len(word) && j < endRow; i, j = i+1, j+1 {
		if crosswordBytes[j][fixedCol] == blankchar {
			continue
		}
		if crosswordBytes[j][fixedCol] == word[i] {
			continue
		}
		return false
	}
	return true
}

func convertToString(crosswordBytes [][]byte) []string {

	crossword := make([]string, 0)
	for _, row := range crosswordBytes {
		crossword = append(crossword, string(row))
	}
	return crossword
}

func copymap(newmap, oldmap map[string]bool) {

	for key := range oldmap {
		newmap[key] = true
	}
}

func convertToMap(arr []string) map[string]bool {

	mapToReturn := make(map[string]bool)
	for _, s := range arr {
		mapToReturn[s] = true
	}
	return mapToReturn
}

func getBlanksToFill(crosswordBytes [][]byte) []blank {

	if len(crosswordBytes) == 0 {
		return nil
	}

	blanks := make([]blank, 0)
	//add horizontal blanks
	for i, line := range crosswordBytes {

		currBlank := blank{starty: i, endy: i}
		var blankInProgress bool

		for j, ch := range line {
			if ch == blankchar && !blankInProgress && j != len(line)-1 && line[j+1] == blankchar {
				blankInProgress = true
				currBlank.startx = j
				continue
			}
			if (ch != blankchar || j == len(line)-1) && blankInProgress {
				if j == len(line)-1 && ch == blankchar {
					currBlank.endx = j + 1
				} else {
					currBlank.endx = j
				}
				currBlank.length = currBlank.endx - currBlank.startx
				blanks = append(blanks, currBlank)
				currBlank = blank{starty: i, endy: i}
				blankInProgress = false
			}
		}
	}

	//add vertical blanks
	for col := 0; col < len(crosswordBytes[0]); col++ {

		currBlank := blank{startx: col, endx: col}
		var blankInProgress bool
		for row := 0; row < len(crosswordBytes); row++ {

			if crosswordBytes[row][col] == blankchar && !blankInProgress && row != len(crosswordBytes)-1 && crosswordBytes[row+1][col] == blankchar {
				blankInProgress = true
				currBlank.starty = row
				continue
			}
			if (crosswordBytes[row][col] != blankchar || row == len(crosswordBytes)-1) && blankInProgress {
				if row == len(crosswordBytes)-1 && crosswordBytes[row][col] == blankchar {
					currBlank.endy = row + 1
				} else {
					currBlank.endy = row
				}
				currBlank.length = currBlank.endy - currBlank.starty
				blanks = append(blanks, currBlank)
				currBlank = blank{startx: col, endx: col}
				blankInProgress = false
			}
		}
	}
	return blanks
}

func getCrosswordBytes(crossword []string) [][]byte {

	crosswordBytes := make([][]byte, 0)
	for _, line := range crossword {
		crosswordBytes = append(crosswordBytes, []byte(line))
	}
	return crosswordBytes
}
