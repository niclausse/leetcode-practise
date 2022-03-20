package stack

import (
	"errors"
	"sync"
)

// ArrayStack 数组实现栈-顺序栈
type ArrayStack struct {
	size int // capacity of stack
	data []interface{}
	lock sync.RWMutex
}

func (a *ArrayStack) Push(item interface{}) error {
	if item == nil {
		return errors.New("stack: element in stack should not be nil")
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	if a.size == len(a.data) {
		return errors.New("stack: capacity loaded")
	}

	a.data = append(a.data, item)
	return nil
}

func (a *ArrayStack) Pop() {
	if len(a.data) == 0 {
		return
	}

	a.data = a.data[:len(a.data)-1]
	return
}

func (a *ArrayStack) Top() interface{} {
	return a.data[len(a.data)-1]
}

func NewArrayStack(size int) Stack {
	return &ArrayStack{
		size: size,
		data: make([]interface{}, 0, size),
	}
}
