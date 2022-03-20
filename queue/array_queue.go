package queue

import "errors"

type ArrayQueue struct {
	size int // capacity
	data []interface{}
	head int
	tail int
}

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{size: size, data: make([]interface{}, 0, size)}
}

func (a *ArrayQueue) Enqueue(e interface{}) error {
	if a.tail == a.size {
		if a.head == 0 {
			return errors.New("queue: capacity loaded")
		}

		// 数据搬移
		for i := a.head; i < a.tail; i++ {
			a.data[i-a.head] = a.data[i]
		}

		a.tail -= a.head
		a.head = 0
	}

	a.data = append(a.data, e)
	a.tail++
	return nil
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.head == a.tail {
		return nil
	}

	res := a.data[a.head]
	a.head++
	return res
}
