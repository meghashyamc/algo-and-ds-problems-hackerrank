package linkedlists

import (
	"testing"
)

//the third array indicates the merge point indices
var inputArrMergePoint = [][][]int32{{{1, 2, 3}, {1, 2, 3}, {1, 1}}, {{1, 2, 3}, {1, 3}, {2, 1}}, {{1, 2, 3}, {1, 3}, {100, 100}}, {{1}, {1}, {0, 0}}, {{}, {}, {100, 100}}}
var answerArrMergePoint = []int32{2, 3, -1, 1, -1}

func TestMergePoint(t *testing.T) {
	for i, arrs := range inputArrMergePoint {

		res := findMergeNode(createLinkedLists(arrs[0], arrs[1], arrs[2][0], arrs[2][1]))
		if res != answerArrMergePoint[i] {
			t.Error("expected the merge point for", arrs[0], "and", arrs[1], "to be", answerArrMergePoint[i], "found", res)
		}
	}
}

func createLinkedLists(arr1 []int32, arr2 []int32, mergeIndex1 int32, mergeIndex2 int32) (*SinglyLinkedListNode, *SinglyLinkedListNode) {

	if len(arr1) == 0 || len(arr2) == 0 || int(mergeIndex1) >= len(arr1) || int(mergeIndex2) >= len(arr2) {
		return nil, nil
	}
	var head1, head2, currnode, prevnode, mergenode *SinglyLinkedListNode

	for i, el := range arr1 {
		currnode = &SinglyLinkedListNode{data: el}
		if i == int(mergeIndex1) {
			mergenode = currnode
		}
		if prevnode != nil {
			prevnode.next = currnode
			prevnode = currnode
		} else {
			head1 = currnode
		}
		prevnode = currnode
	}

	prevnode = nil
	for i := 0; i < int(mergeIndex2); i++ {
		currnode = &SinglyLinkedListNode{data: arr2[i]}

		if prevnode != nil {
			prevnode.next = currnode
			prevnode = currnode
		} else {
			head2 = currnode
		}
		prevnode = currnode
	}
	if prevnode != nil {
		prevnode.next = mergenode
	} else {
		head2 = mergenode
	}
	return head1, head2

}
