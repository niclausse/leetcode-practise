package queue

type Queue interface {
	Enqueue(e interface{}) error
	Dequeue() interface{}
}
