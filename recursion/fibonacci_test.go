package recursion

import (
	"testing"
)

var inputsFibonacci = []int{-1, 0, 1, 5, 10, 25, 30}
var outputsFibonacci = []int{-1, 0, 1, 5, 55, 75025, 832040}

func TestFibonacci(t *testing.T) {

	for i, input := range inputsFibonacci {
		res := fibonacci(input)
		if res != outputsFibonacci[i] {
			t.Error("expected", outputsFibonacci[i], "as fibonacci(", input, "), got", res)
		}
	}
}
