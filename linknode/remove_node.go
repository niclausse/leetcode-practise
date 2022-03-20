package linknode

// 注： node不是尾节点
// 思路： 4->5->1->2
//       4->1->1->2
//       4->1---->2
func remove(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func removeNode(head, node *ListNode) *ListNode {
	if head == node {
		return head.Next
	}

	var prev, cur = &ListNode{Next: head}, head
	for cur != nil {
		if cur == node {
			prev.Next = cur.Next
			cur = nil
			return head
		}

		prev = cur
		cur = cur.Next
	}

	return head
}
