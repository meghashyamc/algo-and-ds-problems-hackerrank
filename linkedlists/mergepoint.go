package linkedlists

// Given pointers to the head nodes of 2 linked lists that merge together at some point, find the Node where the two lists merge. It is guaranteed that the two head Nodes will be different, and neither will be NULL.

// In the diagram below, the two lists converge at Node x:

// [List #1] a--->b--->c
//                      \
//                       x--->y--->z--->NULL
//                      /
//      [List #2] p--->q
// Complete the int findMergeNode(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) method so that it finds and returns the data value of the Node where the two lists merge.

func findMergeNode(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) int32 {
	lenOfList1 := 0
	currnode := head1
	for {
		if currnode == nil {
			break
		}
		currnode = currnode.next
		lenOfList1++
	}
	if lenOfList1 == 0 {
		return -1
	}

	lenOfList2 := 0
	currnode = head2
	for {
		if currnode == nil {
			break
		}
		currnode = currnode.next
		lenOfList2++
	}
	if lenOfList2 == 0 {
		return -1
	}

	skipNodesNum := lenOfList1 - lenOfList2
	// fmt.Println("skipNodesNum:", skipNodesNum)
	currnode1 := head1
	currnode2 := head2
	for {
		if skipNodesNum > 0 {
			if currnode1 == nil {
				return -1
			}
			currnode1 = currnode1.next
			skipNodesNum--
			continue
		}
		if skipNodesNum < 0 {
			if currnode2 == nil {
				return -1
			}
			currnode2 = currnode2.next
			skipNodesNum++
			continue
		}
		if skipNodesNum == 0 {
			if currnode1 == nil || currnode2 == nil {
				return -1
			}
			// fmt.Println("currnode1:", currnode1, "currnode2:", currnode2)
			if currnode1 == currnode2 {
				return currnode1.data
			}

			currnode1 = currnode1.next
			currnode2 = currnode2.next

		}
	}
}
