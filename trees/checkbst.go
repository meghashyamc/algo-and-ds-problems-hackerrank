package trees

// For the purposes of this challenge, we define a binary search tree to be a binary tree with the following properties:

// The data value of every node in a node's left subtree is less than the data value of that node.
// The data value of every node in a node's right subtree is greater than the data value of that node.
// The data value of every node is distinct.

// Given the root node of a binary tree, determine if it is a binary search tree.

// Function Description

// Complete the function checkBST in the editor below. It must return a boolean denoting whether or not the binary tree is a binary search tree.

// checkBST has the following parameter(s):

// root: a reference to the root node of a tree to test
//more info: https://www.hackerrank.com/challenges/ctci-is-binary-search-tree

func checkBST(root *node) bool {

	if root == nil {
		return true
	}
	if root.left != nil && root.left.data >= root.data {
		return false
	}
	if root.right != nil && root.right.data <= root.data {
		return false
	}
	return checkBST(root.left) && checkBST(root.right)
}
