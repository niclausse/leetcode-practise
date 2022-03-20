package linknode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代法
func mergeTwoLists1(list1, list2 *ListNode) *ListNode {
	prevHead := new(ListNode)
	prev := prevHead

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			prev.Next = list2
			list2 = list2.Next
			prev = prev.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
			prev = prev.Next
		}
	}

	if list1 == nil {
		prev.Next = list2
	} else if list2 == nil {
		prev.Next = list1
	}

	return prevHead.Next
}

// 递归法
func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var result *ListNode

	if list1.Val < list2.Val {
		result = list1
		result.Next = mergeTwoLists2(list1.Next, list2)
	} else {
		result = list2
		result.Next = mergeTwoLists2(list1, list2.Next)
	}

	return result
}
