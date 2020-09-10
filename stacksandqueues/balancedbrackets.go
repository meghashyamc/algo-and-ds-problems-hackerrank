package stacksandqueues

// A bracket is considered to be any one of the following characters: (, ), {, }, [, or ].

// Two brackets are considered to be a matched pair if the an opening bracket (i.e., (, [, or {) occurs to the left of a closing bracket (i.e., ), ], or }) of the exact same type. There are three types of matched pairs of brackets: [], {}, and ().

// A matching pair of brackets is not balanced if the set of brackets it encloses are not matched. For example, {[(])} is not balanced because the contents in between { and } are not balanced. The pair of square brackets encloses a single, unbalanced opening bracket, (, and the pair of parentheses encloses a single, unbalanced closing square bracket, ].

// By this logic, we say a sequence of brackets is balanced if the following conditions are met:

// It contains no unmatched brackets.
// The subset of brackets enclosed within the confines of a matched pair of brackets is also a matched pair of brackets.
// Given n strings of brackets, determine whether each sequence of brackets is balanced. If a string is balanced, return YES. Otherwise, return NO.

// Function Description

// Complete the function isBalanced in the editor below. It must return a string: YES if the sequence is balanced or NO if it is not.

// isBalanced has the following parameter(s):

// s: a string of brackets
//more info: https://www.hackerrank.com/challenges/balanced-brackets

const yes = "YES"
const no = "NO"
const roundBracketOpen = '('
const roundBracketClose = ')'
const curlyBracketOpen = '{'
const curlyBracketClose = '}'
const squareBracketOpen = '['
const squareBracketClose = ']'

var bracketsList = map[byte]bool{roundBracketOpen: true, roundBracketClose: true, curlyBracketOpen: true,
	curlyBracketClose: true, squareBracketOpen: true, squareBracketClose: true}

type stack []byte

func (s *stack) push(ch byte) {

	if s == nil {
		s = &stack{}
	}

	*s = append(*s, ch)
}

func (s *stack) pop() byte {
	if len(*s) == 0 {
		return 0
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *stack) peek() byte {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func isBalanced(s string) string {

	// fmt.Println("looking at", s)
	n := len(s)
	if n%2 != 0 {
		return no
	}
	if n == 0 {
		return yes
	}

	openBracketStack := stack{}

	for i := range s {
		// fmt.Println("looking at character", string(s[i]))
		if _, ok := bracketsList[s[i]]; !ok {
			return no
		}
		if isOpenBracket(s[i]) {
			openBracketStack.push(s[i])
			// fmt.Println("stack:", openBracketStack)
			continue
		}
		if len(openBracketStack) == 0 {
			// fmt.Println("stack is empty, will return NO")
			return no
		}
		if getOpposite(openBracketStack.peek()) == s[i] {
			// fmt.Println("found opposite in stack!")
			openBracketStack.pop()
			// fmt.Println("stack:", openBracketStack)
			continue
		}
		return no
	}

	if len(openBracketStack) == 0 {
		return yes
	}
	return no
}

func isOpenBracket(ch byte) bool {

	switch ch {
	case roundBracketOpen:
		return true
	case curlyBracketOpen:
		return true
	case squareBracketOpen:
		return true

	default:
		return false
	}
}

func getOpposite(ch byte) byte {

	switch ch {
	case roundBracketOpen:
		return roundBracketClose
	case curlyBracketOpen:
		return curlyBracketClose
	case squareBracketOpen:
		return squareBracketClose

	default:
		return 0
	}
}
