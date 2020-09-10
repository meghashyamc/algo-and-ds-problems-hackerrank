package trees

import (
	"testing"
)

func TestLCA(t *testing.T) {

	node6 := node{data: 6}
	node5 := node{data: 5, right: &node6}
	node4 := node{data: 3}
	node3 := node{data: 4, left: &node4, right: &node5}
	node2 := node{data: 1}
	node0 := node{data: 2, left: &node2, right: &node3}

	res := lca(&node0, 3, 6)
	if res != &node3 {
		t.Error("expected LCA of 4 and 6 to be", &node3, "got", res)
	}

	node6 = node{data: 6}
	node5 = node{data: 3}
	node4 = node{data: 1}
	node3 = node{data: 7, left: &node6}
	node2 = node{data: 2, left: &node4, right: &node5}
	node0 = node{data: 4, left: &node2, right: &node3}

	res = lca(&node0, 1, 7)
	if res != &node0 {
		t.Error("expected LCA of 1 and 7 to be", &node0, "got", res)
	}

	node2 = node{data: 2}
	node0 = node{data: 1, right: &node2}
	res = lca(&node0, 1, 2)
	if res != &node0 {
		t.Error("expected LCA of 1 and 2 to be", &node0, "got", res)
	}

	node6 = node{data: 8}
	node5 = node{data: 6}
	node4 = node{data: 7, left: &node5, right: &node6}
	node3 = node{data: 4}
	node2 = node{data: 2}
	node1 := node{data: 3, left: &node2, right: &node3}
	node0 = node{data: 5, left: &node1, right: &node4}
	res = lca(&node0, 6, 6)
	if res != &node4 {
		t.Error("expected LCA of 6 and 6 to be", &node4, "got", res)
	}

	res = lca(&node0, 6, 7)
	if res != &node0 {
		t.Error("expected LCA of 6 and 7 to be", &node0, "got", res)
	}

	res = lca(&node0, 2, 6)
	if res != &node0 {
		t.Error("expected LCA of 2 and 6 to be", &node0, "got", res)
	}

	res = lca(&node0, 2, 4)
	if res != &node1 {
		t.Error("expected LCA of 2 and 4 to be", &node1, "got", res)
	}
}
