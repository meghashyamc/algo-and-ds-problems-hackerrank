package linkedlists

// You’re given the pointer to the head node of a linked list, an integer to add to the list and the position at which the integer must be inserted. Create a new node with the given integer, insert this node at the desired position and return the head node.

// A position of 0 indicates head, a position of 1 indicates one node away from the head and so on. The head pointer given may be null meaning that the initial list is empty.

// As an example, if your list starts as 1->2->3 and you want to insert a node at position 2 with data 4, your new list should be 1->2->4->3

// Function Description Complete the function insertNodeAtPosition in the editor below. It must return a reference to the head node of your finished list.

// insertNodeAtPosition has the following parameters:

// head: a SinglyLinkedListNode pointer to the head of the list
// data: an integer value to insert as data in your new node
// position: an integer position to insert the new node, zero based indexing
//more info: https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

	var prev *SinglyLinkedListNode
	var i int32
	currnode := llist
	for i = 0; ; i++ {
		if i == position {

			newnode := &SinglyLinkedListNode{data: data, next: currnode}
			if prev != nil {
				prev.next = newnode
				break
			} else {
				return newnode
			}

		}
		prev = currnode
		if currnode != nil {
			currnode = currnode.next
		} else {
			break
		}
	}
	return llist
}
