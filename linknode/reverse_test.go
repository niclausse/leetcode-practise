package linknode

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	cur := Reverse(head)
	for cur != nil {
		fmt.Printf("%d\t", cur.Val)

		cur = cur.Next
	}
}
