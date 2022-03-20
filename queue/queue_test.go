package queue

import "testing"

func TestArrayQueue_Enqueue(t *testing.T) {
	q := NewArrayQueue(10)

	t.Log(q.data)

	for i := 0; i < 10; i++ {
		if err := q.Enqueue(i); err != nil {
			t.Log(q.data)
			t.Error(err)
			return
		}
	}

	t.Log(q.data)
}

func TestArrayQueue_Dequeue(t *testing.T) {
	q := NewArrayQueue(10)

	for i := 0; i < 10; i++ {
		if err := q.Enqueue(i); err != nil {
			t.Error(err)
			return
		}
	}

	t.Log(q.data)

	for i := 0; i < 5; i++ {
		t.Log(q.Dequeue())
	}

	t.Log(q.data)

	for i := 0; i < 5; i++ {
		if err := q.Enqueue('a' + i); err != nil {
			t.Error(err)
			return
		}
	}

	t.Log(q.data)
}
