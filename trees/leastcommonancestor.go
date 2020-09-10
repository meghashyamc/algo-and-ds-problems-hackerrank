package trees

// You are given the pointer to the root of a binary search tree and two values v1 and v2. You need to return the lowest common ancestor (LCA) of v1 and v2 in the binary search tree.

// https://s3.amazonaws.com/hr-assets/0/1529959649-81b68736f7-lcaexample.png

// In the diagram above, the lowest common ancestor of the nodes 4 and 6 is the node 3. Node 3 is the lowest node which has nodes 4 and 6 as descendants.

// Function Description

// Complete the function lca in the editor below. It should return a pointer to the lowest common ancestor node of the two values given.

// lca has the following parameters:
// - root: a pointer to the root node of a binary search tree
// - v1: a node.data value
// - v2: a node.data value
//more info:https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor

func lca(root *node, v1 int, v2 int) *node {

	if root == nil || (root.data == v1 && root.data == v2) {
		return nil
	}
	ancestors := make(map[*node]bool)

	currNode := root
	//searching for v1, storing all possible ancestors in a map
	for {

		if currNode == nil {
			return nil
		}

		if currNode.data == v1 {
			break
		}
		ancestors[currNode] = true
		if currNode.data < v1 {
			currNode = currNode.right

		} else {
			currNode = currNode.left
		}

	}

	currNode = root

	leastCommonAncestor := root
	//searching for v2, updating least common ancestor
	//based on the just encountered ancestor
	//that is already there in v1's ancestor map
	for {
		if currNode == nil {
			return nil
		}
		if currNode.data == v2 {
			return leastCommonAncestor
		}
		if ancestors[currNode] {
			leastCommonAncestor = currNode
		}
		if currNode.data < v2 {
			currNode = currNode.right

		} else {
			currNode = currNode.left
		}

	}
}
