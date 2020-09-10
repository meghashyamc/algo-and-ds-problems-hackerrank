package linkedlists

import (
	"reflect"
	"testing"
)

var inputArrSortedInsert = [][]int32{{1, 3, 4, 10, 5}, {5}, {4, -1}}
var answerArrSortedInsert = [][]int32{{1, 3, 4, 5, 10}, {5}, {-1, 4}}

func TestSortedInsert(t *testing.T) {
	for i, arr := range inputArrSortedInsert {

		res := convertToArrDoubleList(sortedInsert(createLinkedListDouble(arr), arr[len(arr)-1]))
		if !reflect.DeepEqual(res, answerArrSortedInsert[i]) {
			t.Error("expected the final linked list for", arr, "to be", answerArrSortedInsert[i], "found", res)
		}
	}
}

func convertToArrDoubleList(llist *DoublyLinkedListNode) []int32 {

	convertedArr := make([]int32, 0)

	for {
		if llist == nil {
			break
		}
		convertedArr = append(convertedArr, llist.data)
		llist = llist.next
	}

	return convertedArr

}

func createLinkedListDouble(arr []int32) *DoublyLinkedListNode {

	if len(arr) == 0 {
		return nil
	}
	var firstnode *DoublyLinkedListNode
	var prevnode *DoublyLinkedListNode
	for _, el := range arr[:len(arr)-1] {

		currnode := &DoublyLinkedListNode{data: el}
		if prevnode != nil {
			prevnode.next = currnode
			currnode.prev = prevnode
		} else {
			firstnode = currnode
		}
		prevnode = currnode
	}
	return firstnode
}
