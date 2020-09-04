package dynamicprogramming

import (
	"os"
	"testing"
)

var s1Arr, s2Arr []string

const yeslimit = 3

func TestMain(m *testing.M) {
	setup()
	code := m.Run()

	os.Exit(code)
}

func setup() {

	s1Arr = []string{"daBcd", "AbcDE", "BFZZVHdQYHQEMNEFFRFJTQmNWHFVXRXlGTFNBqWQmyOWYWSTDSTMJRYHjBNTEWADLgHVgGIRGKFQSeCXNFNaIFAXOiQORUDROaNoJPXWZXIAABZKSZYFTDDTRGZXVZZNWNRHMvSTGEQCYAJSFvbqivjuqvuzafvwwifnrlcxgbjmigkms", "AfPZN", "MBQEVZPBjcbswirgrmkkfvfvcpiukuxlnxkkenqp"}
	s2Arr = []string{"ABC", "ABDE", "BFZZVHQYHQEMNEFFRFJTQNWHFVXRXGTFNBWQOWYWSTDSTMJRYHBNTEWADLHVGIRGKFQSCXNFNIFAXOQORUDRONJPXWZXIAABZKSZYFTDDTRGZXVZZNWNRHMSTGEQCYAJSF", "APZNC", "MBQEVZP"}

}

func TestAbbreviation(t *testing.T) {

	for i := 0; i < yeslimit; i++ {
		res := abbreviation(s1Arr[i], s2Arr[i])
		if res != yes {
			t.Error("in the case of", s1Arr[i], s2Arr[i], "expected YES, got", res)
		}

	}

	for i := yeslimit; i < len(s1Arr); i++ {
		res := abbreviation(s1Arr[i], s2Arr[i])
		if res != no {
			t.Error("in the case of", s1Arr[i], s2Arr[i], "expected NO, got", res)
		}
	}

}

// func TestAbbreviationHelper(t *testing.T) {

// 	for i := 0; i < 2; i++ {
// 		res := abbreviationHelper(s1Arr[i], 0, s2Arr[i], 0)
// 		if !res {
// 			t.Errorf("expected true, got %v", res)
// 		}
// 		memCacheAbbr = make([][]int, 0)
// 	}

// 	for i := 2; i < len(s1Arr); i++ {
// 		res := abbreviationHelper(s1Arr[i], 0, s2Arr[i], 0)
// 		if res {
// 			t.Errorf("expected false, got %v", res)
// 		}
// 		memCacheAbbr = make([][]int, 0)
// 	}

// }

func strToPtr(s string) *string {

	return &s
}
