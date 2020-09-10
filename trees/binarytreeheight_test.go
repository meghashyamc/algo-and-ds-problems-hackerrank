package trees

import (
	"testing"
)

func TestHeight(t *testing.T) {

	node6 := node{data: 7}
	node5 := node{data: 6, right: &node6}
	node4 := node{data: 4}
	node3 := node{data: 1}
	node2 := node{data: 5, left: &node4, right: &node5}
	node1 := node{data: 2, left: &node3}
	node0 := node{data: 3, left: &node1, right: &node2}

	res := height(&node0)
	if res != 3 {
		t.Error("expected height 3 got", res)
	}

	node6 = node{data: 7}
	node5 = node{data: 5}
	node4 = node{data: 3}
	node3 = node{data: 1}
	node2 = node{data: 6, left: &node5, right: &node6}
	node1 = node{data: 2, left: &node3, right: &node4}
	node0 = node{data: 4, left: &node1, right: &node2}
	res = height(&node0)
	if res != 2 {
		t.Error("expected height 2 got", res)
	}

	res = height(nil)
	if res != 0 {
		t.Error("expected height 0 got", res)
	}
	node0 = node{data: 4}
	res = height(nil)
	if res != 0 {
		t.Error("expected height 0 got", res)
	}

}
