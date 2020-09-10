package stacksandqueues

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// A queue is an abstract data type that maintains the order in which elements were added to it, allowing the oldest elements to be removed from the front and new elements to be added to the rear. This is called a First-In-First-Out (FIFO) data structure because the first element added to the queue (i.e., the one that has been waiting the longest) is always the first one to be removed.

// A basic queue has the following operations:

// Enqueue: add a new element to the end of the queue.
// Dequeue: remove the element from the front of the queue and return it.
// In this challenge, you must first implement a queue using two stacks. Then process q queries, where each query is one of the following 3 types:

// 1 x: Enqueue element x into the end of the queue.
// 2: Dequeue the element at the front of the queue.
// 3: Print the element at the front of the queue.

// Function Description

// Complete the put, pop, and peek methods in the editor below. They must perform the actions as described above.
//more info: https://www.hackerrank.com/challenges/ctci-queue-using-two-stacks
// Input Format

// The first line contains a single integer,q , the number of queries.

// Each of the next q lines contains a single query in the form described in the problem statement above. All queries start with an integer denoting the query type, but only query 1 is followed by an additional space-separated value,x , denoting the value to be enqueued.

// It is guaranteed that a valid answer always exists for each query of types 2 and 3.
// Output Format

// For each query of type 3, return the value of the element at the front of the fifo queue on a new line.

type queue []int

const typePush = 1
const typePop = 2
const typePeek = 3

func (q *queue) push(el int) {

	if q == nil {
		q = &queue{}
	}
	*q = append(*q, el)
}

func (q *queue) pop() (int, error) {
	if q == nil || len(*q) == 0 {
		return 0, errors.New("can't pop from empty queue")
	}
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped, nil
}

func (q *queue) peek() (int, error) {
	if q == nil || len(*q) == 0 {
		return 0, errors.New("can't peek into empty queue")
	}
	return (*q)[0], nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	numOfQueriesStr, err := reader.ReadString('\n')
	checkError(err)
	numOfQueriesStr = strings.TrimRight(numOfQueriesStr, "\n")
	numOfQueries, err := strconv.Atoi(numOfQueriesStr)
	checkError(err)
	q := queue{}
	for i := 0; i < numOfQueries; i++ {
		queryStr, err := reader.ReadString('\n')
		checkError(err)
		queryStr = strings.TrimRight(queryStr, "\n")
		queryParts := strings.Split(queryStr, " ")
		queryType, err := strconv.Atoi(queryParts[0])
		checkError(err)
		switch queryType {
		case typePush:
			elementToPush, err := strconv.Atoi(queryParts[1])
			checkError(err)
			q.push(elementToPush)

		case typePop:
			_, err := q.pop()
			checkError(err)
		case typePeek:
			val, err := q.peek()
			checkError(err)
			fmt.Println(val)

		}

	}
}

func checkError(err error) {
	if err != nil {
		if err != io.EOF {
			panic(err)
		}

	}
}
