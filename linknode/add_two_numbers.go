package linknode

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	prevHead := new(ListNode)
	prev := prevHead

	quota := 0 // 进位

	for l1 != nil || l2 != nil {
		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := quota + v1 + v2

		quota = sum / 10

		prev.Next = &ListNode{Val: sum % 10}

		prev = prev.Next
	}

	if quota > 0 {
		prev.Next = &ListNode{Val: quota}
	}

	return prevHead.Next
}
