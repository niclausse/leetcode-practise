package linknode

import "errors"

type SingleLinklist struct {
	size int
	len  int
	head *SingleNode
	tail *SingleNode
}

func (s *SingleLinklist) Len() int {
	return s.len
}

func (s *SingleLinklist) PushBack(v interface{}) error {
	if s.size <= s.len {
		return errors.New("linklist: capacity over loaded")
	}

	node := &SingleNode{Val: v}

	if s.len == 0 {
		s.head = node
		s.tail = node
		s.len++
		return nil
	}

	s.tail.Next = node
	s.tail = node
	return nil
}

func (s *SingleLinklist) PushFront(v interface{}) error {
	if s.size <= s.len {
		return errors.New("linklist: capacity over loaded")
	}

	node := &SingleNode{Val: v}

	if s.len == 0 {
		s.head = node
		s.tail = node
		s.len++
		return nil
	}

	node.Next = s.head
	s.head = node
	s.len++
	return nil
}

func (s *SingleLinklist) Remove(n *SingleNode) {
	if s.len == 0 || n == nil {
		return
	}

	if s.head.Val == n.Val && s.head.Next == n.Next {
		s.head = s.head.Next
		s.len--
		return
	}

	if s.len == 1 {

	}

	prev := s.head
	cur := prev.Next

	for cur != nil {
		if cur.Next == n.Next && cur.Val == n.Val {
			if s.tail == cur {
				s.tail = prev
			}

			prev.Next = cur.Next
			s.len--
			return
		}

		prev = prev.Next
		cur = cur.Next
	}
}

func NewSingleLinklist(size int) *SingleLinklist {
	return &SingleLinklist{size: size}
}

type SingleNode struct {
	Next *SingleNode
	Val  interface{}
}
