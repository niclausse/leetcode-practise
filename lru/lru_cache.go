package lru

import (
	"container/list"
	"sync"
)

type Cache struct {
	cap      int
	values   *list.List
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

type entry struct {
	key, value interface{}
}

func Constructor(capacity int) *Cache {
	return &Cache{
		cap:      capacity,
		values:   list.New(),
		cacheMap: make(map[interface{}]*list.Element),
		lock:     sync.Mutex{},
	}
}

// ["LRUCache","get","put","get","put","put","get","get"]
// [[2],       [2],  [2,6], [1], [1,5], [1,2],[1],  [2]]

func (l *Cache) Get(key interface{}) interface{} {
	v, ok := l.cacheMap[key]
	if !ok {
		return -1
	}

	l.values.MoveToFront(v)
	return v.Value.(*entry).value
}

func (l *Cache) Put(key, value interface{}) {
	if v, ok := l.cacheMap[key]; ok {
		l.values.MoveToFront(v)
		v.Value.(*entry).value = value
		return
	}

	if l.cap == l.values.Len() {
		tail := l.values.Back()
		l.values.Remove(tail)
		delete(l.cacheMap, tail.Value.(*entry).key)
	}

	l.cacheMap[key] = l.values.PushFront(&entry{key: key, value: value})
}
