package linknode

// RemoveNthFromEnd 删除倒数第n个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 双指针法
	// 第一个指针先走n步, 当第一个指针走完时, 第二个指针刚好处于倒数第n个节点
	dummy := &ListNode{Next: head}
	first := head // 第一个指针起始位置比second指针多一个,  方便删除第n个节点
	second := dummy

	// first先走n步
	for i := 0; i < n; i++ {
		// 输入n不合理, 直接返回, 不做删除操作
		if first == nil {
			return head
		}
		first = first.Next
	}

	// first second同时往后走
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 删除倒数第n个节点
	second.Next = second.Next.Next
	return dummy.Next
}
