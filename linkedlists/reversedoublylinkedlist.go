package linkedlists

// You’re given the pointer to the head node of a doubly linked list. Reverse the order of the nodes in the list. The head node might be NULL to indicate that the list is empty. Change the next and prev pointers of all the nodes so that the direction of the list is reversed. Return a reference to the head node of the reversed list.

// Function Description

// Complete the reverse function in the editor below. It should return a reference to the head of your reversed list.

// reverse has the following parameter(s):

// head: a reference to the head of a DoublyLinkedList
//more info: https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list

func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {

	if head == nil {
		return head
	}
	count := 0
	currnode := head
	for {
		count++
		if currnode.next == nil {
			break
		}
		currnode = currnode.next
	}
	lastnode := currnode
	for i := 0; i < count; i++ {
		currnode.next, currnode.prev = currnode.prev, currnode.next
		currnode = currnode.next
	}
	return lastnode

}
