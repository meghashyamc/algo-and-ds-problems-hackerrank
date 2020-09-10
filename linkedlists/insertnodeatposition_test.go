package linkedlists

import (
	"reflect"
	"testing"
)

//last two elements of input slice represent data and position respectively
var inputArr = [][]int32{{16, 13, 7, 1, 2}, {1, 2, 3, 4, 2}, {100, 100}, {5, 100, 100}, {5, 100, -1}}
var answerArr = [][]int32{{16, 13, 1, 7}, {1, 2, 4, 3}, {}, {5}, {5}}

func TestInsertNodeAtPosition(t *testing.T) {

	for i, arr := range inputArr {

		res := convertToArr(insertNodeAtPosition(createLinkedList(arr, 2), arr[len(arr)-2], arr[len(arr)-1]))
		if !reflect.DeepEqual(res, answerArr[i]) {
			t.Error("expected the final linked list for", arr, "to be", answerArr[i], "found", res)
		}
	}
}

func convertToArr(llist *SinglyLinkedListNode) []int32 {

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

func createLinkedList(arr []int32, indicestoignore int) *SinglyLinkedListNode {

	if len(arr) == 0 {
		return nil
	}
	var firstnode *SinglyLinkedListNode
	var prevnode *SinglyLinkedListNode
	for _, el := range arr[:len(arr)-indicestoignore] {

		currnode := &SinglyLinkedListNode{data: el}
		if prevnode != nil {
			prevnode.next = currnode
		} else {
			firstnode = currnode
		}
		prevnode = currnode
	}
	return firstnode
}
