package recursion

// The Fibonacci Sequence

// The Fibonacci sequence appears in nature all around us, in the arrangement of seeds in a sunflower and the spiral of a nautilus for example.

// The Fibonacci sequence begins with fibonacci(0)=0 and fibonacci(1) = 1 and  as its first and second terms. After these first two elements, each subsequent element is equal to the sum of the previous two elements.

// Given n, return the nth number in the sequence.

// As an example,n=5 . The Fibonacci sequence to 6 is [0,1,1,2,3,5,8]. With zero-based indexing, f(5)=5.

// Function Description

// Complete the recursive function fibonacci in the editor below. It must return the nth element in the Fibonacci sequence.

// fibonacci has the following parameter(s):

// n: the integer index of the sequence to return
//more info: https://www.hackerrank.com/challenges/ctci-fibonacci-numbers

func fibonacci(n int) int {

	if n < 0 {
		return -1
	}
	memCache := make([]int, n+1)
	for i := range memCache {
		memCache[i] = -1
	}
	return fibonacciHelper(n, memCache)
}

func fibonacciHelper(n int, memCache []int) int {

	if memCache[n] != -1 {
		return memCache[n]
	}
	if n == 0 {
		memCache[n] = 0
		return memCache[n]
	}
	if n == 1 {
		memCache[n] = 1
		return memCache[n]
	}
	memCache[n] = fibonacciHelper(n-2, memCache) + fibonacciHelper(n-1, memCache)
	return memCache[n]
}
