package stack

type Stack interface {
	Push(item interface{}) error
	Pop()
	Top() interface{}
}
