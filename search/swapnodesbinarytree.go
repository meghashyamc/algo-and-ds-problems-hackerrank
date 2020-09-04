package search

import (
	"errors"
	"fmt"
)

// A binary tree is a tree which is characterized by one of the following properties:

// It can be empty (null).
// It contains a root node only.
// It contains a root node with a left subtree, a right subtree, or both. These subtrees are also binary trees.
// In-order traversal is performed as

// Traverse the left subtree.
// Visit root.
// Traverse the right subtree.
// For this in-order traversal, start from the left child of the root node and keep exploring the left subtree until you reach a leaf. When you reach a leaf, back up to its parent, check for a right child and visit it if there is one. If there is not a child, you've explored its left and right subtrees fully. If there is a right child, traverse its left subtree then its right in the same manner. Keep doing this until you have traversed the entire tree. You will only store the values of a node as you visit when one of the following is true:

// it is the first node visited, the first time visited
// it is a leaf, should only be visited once
// all of its subtrees have been explored, should only be visited once while this is true
// it is the root of the tree, the first time visited
// Swapping: Swapping subtrees of a node means that if initially node has left subtree L and right subtree R, then after swapping, the left subtree will be R and the right subtree, L.

// For example, in the following tree, we swap children of node 1.

//                                 Depth
//     1               1            [1]
//    / \             / \
//   2   3     ->    3   2          [2]
//    \   \           \   \
//     4   5           5   4        [3]
// In-order traversal of left tree is 2 4 1 3 5 and of right tree is 3 5 1 2 4.

// Swap operation:

// We define depth of a node as follows:

// The root node is at depth 1.
// If the depth of the parent node is d, then the depth of current node will be d+1.
// Given a tree and an integer, k, in one operation, we need to swap the subtrees of all the nodes at each depth h, where h ∈ [k, 2k, 3k,...]. In other words, if h is a multiple of k, swap the left and right subtrees of that level.

// You are given a tree of n nodes where nodes are indexed from [1..n] and it is rooted at 1. You have to perform t swap operations on it, and after each swap operation print the in-order traversal of the current state of the tree.

// Function Description

// Complete the swapNodes function in the editor below. It should return a two-dimensional array where each element is an array of integers representing the node indices of an in-order traversal after a swap operation.

// swapNodes has the following parameter(s):
// - indexes: an array of integers representing index values of each node[i], beginning with node[1], the first element, as the root.
// - queries: an array of integers, each representing a k value.
//Note: -1 is used to represent a null node.

//more info: https://www.hackerrank.com/challenges/swap-nodes-algo

type tree struct {
	left          *tree
	right         *tree
	data          int32
	thisnodedepth int32
}

//first in first out queue
type queue []*tree

func (q *queue) push(t *tree) {
	if q == nil {
		q = func() *queue {
			newq := make(queue, 0)
			return &newq
		}()
	}
	*q = append(*q, t)
}

func (q *queue) pop() *tree {

	if q == nil || len(*q) == 0 {
		return nil
	}
	valtopop := (*q)[0]
	*q = (*q)[1:]
	return valtopop

}

func (q *queue) peek() *tree {

	if q == nil || len(*q) == 0 {
		return nil
	}
	return (*q)[0]
}

func (q *queue) print() {

	if q == nil {
		fmt.Println("queue is empty")
		return
	}
	fmt.Print("queue is: ")
	for _, el := range *q {
		fmt.Print(el, " ")
		fmt.Println()
	}
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {

	treeObj := convertToTree(indexes)

	n := len(indexes)

	answerArr := make([][]int32, 0)
	for _, query := range queries {
		if query < 1 || query > int32(n) {
			panic(errors.New("query out of bounds"))
		}
		swapNodesSingleQuery(treeObj, query)
		inorderArr := make([]int32, 0)
		traverseInOrder(treeObj, &inorderArr)
		answerArr = append(answerArr, inorderArr)
	}
	return answerArr

}

func swapNodesSingleQuery(t *tree, query int32) {

	if t == nil {
		return
	}
	var queryFactor int32 = 2
	origQuery := query

	q := &queue{}
	q.push(t)
	for {
		popped := q.pop()

		if popped.left.data != -1 {
			q.push(popped.left)
		}
		if popped.right.data != -1 {
			q.push(popped.right)
		}
		if query == popped.thisnodedepth {
			popped.left, popped.right = popped.right, popped.left
			if nextnode := q.peek(); nextnode != nil {
				if query < nextnode.thisnodedepth {
					query = origQuery * queryFactor
					queryFactor++
				}
			}
		}
		// q.print()
		if len(*q) == 0 {
			return
		}

	}

}

func traverseInOrder(t *tree, inorderArr *[]int32) {

	if t == nil {
		return
	}
	traverseInOrder(t.left, inorderArr)
	if t.data != -1 {
		*inorderArr = append(*inorderArr, t.data)
	}
	traverseInOrder(t.right, inorderArr)

}

func convertToTree(indexes [][]int32) *tree {

	q := &queue{}
	root := &tree{data: 1, thisnodedepth: 1}
	q.push(root)
	for _, thislevelIndexes := range indexes {

		treeObj := q.pop()
		treeObj.left = &tree{data: thislevelIndexes[0], thisnodedepth: treeObj.thisnodedepth + 1}
		treeObj.right = &tree{data: thislevelIndexes[1], thisnodedepth: treeObj.thisnodedepth + 1}
		if thislevelIndexes[0] != -1 {
			q.push(treeObj.left)
		}
		if thislevelIndexes[1] != -1 {
			q.push(treeObj.right)
		}
	}

	return root
}
