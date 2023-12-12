package linknode

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur // prev 向前走一步
		cur = next // cur 向前走一步
	}

	return prev
}
