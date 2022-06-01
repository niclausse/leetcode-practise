package lru

type LinkNode struct {
	Val  int
	Next *LinkNode
}

type LinkListCache struct {
	Cap  int
	Len  int
	Data *LinkNode
}

func NewLinkListCache() *LinkListCache {
	return &LinkListCache{
		Cap:  20,
		Len:  0,
		Data: nil,
	}
}

func (l *LinkListCache) Put(e int) {
	if l.Len == 0 {
		l.Data = &LinkNode{Val: e}
		l.Len = 1
		return
	}

	var prev, head *LinkNode = nil, l.Data

	for l.Data != nil {
		if l.Data.Val == e {
			if prev == nil { // 元素存在，且位于头部
				return
			}

			prev.Next = l.Data.Next
			l.Data.Next = head
			return
		}
	}

	l.Data = &LinkNode{Val: e, Next: l.Data}
}
