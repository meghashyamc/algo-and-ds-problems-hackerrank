package recursion

import (
	"testing"
)

var inputsSuperDigitN = []string{"148", "9875", "123", "9875"}
var inputsSuperDigitK = []int32{3, 4, 3, 4}
var outputsSuperDigit = []int32{3, 8, 9, 8}

func TestSuperDigit(t *testing.T) {

	for i := 0; i < len(inputsSuperDigitN); i++ {
		res := superDigit(inputsSuperDigitN[i], inputsSuperDigitK[i])
		if res != outputsSuperDigit[i] {
			t.Error("expected", outputsSuperDigit[i], "for input:", inputsSuperDigitN[i], inputsSuperDigitK[i], "got", res)
		}
	}
}
