package linkedlists

// Given a reference to the head of a doubly-linked list and an integer,data , create a new DoublyLinkedListNode object having data value data and insert it into a sorted linked list while maintaining the sort.

// Function Description

// Complete the sortedInsert function in the editor below. It must return a reference to the head of your modified DoublyLinkedList.

// sortedInsert has two parameters:

// head: A reference to the head of a doubly-linked list of DoublyLinkedListNode objects.
// data: An integer denoting the value of the data field for the DoublyLinkedListNode you must insert into the list.
// Note: Recall that an empty list (i.e., where head == nil) and a list with one element are sorted lists.
//more info: https://www.hackerrank.com/challenges/insert-a-node-into-a-sorted-doubly-linked-list

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {

	prevnode, nextnode := getPrevAndNextNodesForInsert(head, data)

	newnode := &DoublyLinkedListNode{data: data, prev: prevnode, next: nextnode}
	if prevnode != nil {
		prevnode.next = newnode
	} else {
		return newnode
	}
	return head
}

func getPrevAndNextNodesForInsert(head *DoublyLinkedListNode, data int32) (*DoublyLinkedListNode, *DoublyLinkedListNode) {

	if head == nil {
		return nil, nil
	}
	currnode := head

	for {
		if currnode.data > data {
			return currnode.prev, currnode
		}
		if currnode.next != nil {
			currnode = currnode.next
		} else {
			return currnode, nil
		}
	}

}
