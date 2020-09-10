package trees

// The height of a binary tree is the number of edges between the tree's root and its furthest leaf. For example, the following binary tree is of height 2:

// https://s3.amazonaws.com/hr-assets/0/1527626183-88c8070977-isitBSTSample0.png

// Function Description

// Complete the getHeight or height function in the editor. It must return the height of a binary tree as an integer.

// getHeight or height has the following parameter(s):

// root: a reference to the root of a binary tree.
// Note -The Height of binary tree with single node is taken as zero.

// more info: https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree

type node struct {
	data  int
	left  *node
	right *node
}

func height(root *node) int {

	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}

	return 1 + max(height(root.left), height(root.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
