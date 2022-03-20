package linknode

// 判断链表是否有环

// 1' 哈希表法
// 时间复杂度O(N)： 最坏情况下，需要遍历整个链表
// 空间复杂度O(N)： 创建了一个hashTable, 最坏情况下要存储整个链表节点
func hasCycleByHashTable(head *SingleNode) bool {
	seen := make(map[*SingleNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}

		seen[head] = struct{}{}
		head = head.Next
	}

	return false
}

// 2' 快慢指针法
// 快慢指针同时出发， 快指针追上慢指针（重合）则说明有环
// 时间复杂度O(N):
//		当不存在环时，消耗时间为快指针到达尾部需要的时间， N/2 * unit_time
// 		当存在环时，
// 空间复杂度O(1): 只使用了两个指针的额外空间
func hasCycleByFastSlowPointer(head *SingleNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}

// 返回入环节点

func detectCycleByHashTable(head *SingleNode) *SingleNode {
	seen := make(map[*SingleNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		head = head.Next
	}
	return nil
}

func detectCycleByFastSlowPointer(head *SingleNode) *SingleNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
}
