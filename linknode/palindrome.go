package linknode

// 判断是否为回文链表
// 思路：反转后半部分列表后与前半部分比较
// 1‘ 快慢指针找到中间节点/前半段最后一个节点
// 2’ 反转后半部分链表
// 3‘ 比较
// 4’ [可选，还原链表]
// 5‘ 返回
// 时间复杂度： O(N)
// 空间复杂度： O(1)
// 实际场景中的潜在问题： 并发场景需要加锁，有额外的开销
func isPalindrome(head *SingleNode) bool {
	if head == nil {
		return false
	}

	firstHalfEnd := endOfFirstHalf(head)
	// 反转后半部分链表
	secondHalfStart := reverseLinkedList(firstHalfEnd.Next)

	// 比较前半部分和后半部分是否相同
	p1, p2 := head, secondHalfStart
	result := true

	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	reverseLinkedList(secondHalfStart)

	return result
}

func reverseLinkedList(head *SingleNode) *SingleNode {
	var prev, cur *SingleNode = nil, head

	for cur != nil {
		nextTmp := cur.Next // 暂存当前节点的后继节点
		cur.Next = prev     // 将当前节点的前驱节点修改为后继节点

		prev = cur
		cur = nextTmp // 前进一步
	}

	return prev
}

// 快慢指针找到链表前半段最后一个节点
func endOfFirstHalf(head *SingleNode) *SingleNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
