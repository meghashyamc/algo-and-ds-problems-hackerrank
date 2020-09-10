package trees

import (
	"testing"
)

func TestCheckBST(t *testing.T) {

	if !checkBST(nil) {
		t.Error("expected nil to return true, got false")
	}
	node0 := node{data: 1}
	if !checkBST(&node0) {
		t.Error("expected single node to return true, got false")
	}

	node6 := node{data: 7}
	node5 := node{data: 5}
	node4 := node{data: 3}
	node3 := node{data: 1}
	node2 := node{data: 6, left: &node5, right: &node6}
	node1 := node{data: 2, left: &node3, right: &node4}
	node0 = node{data: 4, left: &node1, right: &node2}

	if !checkBST(&node0) {
		t.Error("expected true, got false")
	}
	node6 = node{data: 7}
	node5 = node{data: 100}
	node4 = node{data: 3}
	node3 = node{data: 3}
	node2 = node{data: 6, left: &node5, right: &node6}
	node1 = node{data: 2, left: &node3, right: &node4}
	node0 = node{data: 4, left: &node1, right: &node2}
	if checkBST(&node0) {
		t.Error("expected false, got true")
	}
}
