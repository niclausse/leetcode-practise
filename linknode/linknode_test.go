package linknode

import (
	"strconv"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	cur := RemoveNthFromEnd(head, 2)
	str := ""
	for cur != nil {
		str += strconv.Itoa(cur.Val)
		cur = cur.Next
	}

	if str != "1235" {
		t.Errorf("remove result not expected")
	}
}
