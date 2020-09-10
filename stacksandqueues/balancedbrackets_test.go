package stacksandqueues

import (
	"testing"
)

const yeslimit = 1

var bracketsArr = []string{"{[()]}", "{{[[(())]]}}", "{[(])}", "{[(])}"}

func TestIsBalanced(t *testing.T) {

	for i := 0; i <= yeslimit; i++ {

		if isBalanced(bracketsArr[i]) != yes {
			t.Error("for", bracketsArr[i], "expected YES, found NO")
		}
	}

	for i := yeslimit + 1; i < len(bracketsArr); i++ {

		if isBalanced(bracketsArr[i]) == yes {
			t.Error("for", bracketsArr[i], "expected NO, found YES")
		}
	}
}

func TestIsOpenBracket(t *testing.T) {

	for i := range bracketsArr[0] {

		if i <= 2 && !isOpenBracket(bracketsArr[0][i]) {
			t.Error("for", string(bracketsArr[0][i]), "expected isOpenBracket to return true, got false")
		}
		if i > 2 && isOpenBracket(bracketsArr[0][i]) {
			t.Error("for", string(bracketsArr[0][i]), "expected isOpenBracket to return false, got true")
		}

	}
}

func TestGetOpposite(t *testing.T) {

	for i := range bracketsArr[0] {

		if i >= 3 {
			return
		}
		res := getOpposite(bracketsArr[0][i])
		if i == 0 && res != '}' {
			t.Error("expected } but found", res)
		}
		if i == 1 && res != ']' {
			t.Error("expected ] but found", res)
		}
		if i == 2 && res != ')' {
			t.Error("expected ) but found", res)
		}

	}
}

func TestStack(t *testing.T) {
	s := stack{}

	s.push('a')
	if len(s) != 1 {
		t.Error("after pushing a, expected length of stack to be 1, got", len(s))
	}
	if s.peek() != 'a' {
		t.Error("after pushing a, expected s.peek() to return a, got", s.peek())
	}

	s.pop()

	if len(s) != 0 {
		t.Error("after popping a, expected length of stack to be 0, got", len(s))
	}
}
