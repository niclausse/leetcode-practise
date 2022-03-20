package linknode

// 题24： 两两交换链表中的节点

// 递归
func swapPair(head *SingleNode) *SingleNode {
	newHead := head.Next
	head.Next = swapPair(newHead.Next)
	newHead.Next = head

	return newHead
}
