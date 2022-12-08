package queue

func NewQueue() Queue {
	return &queue{}
}

type Queue interface {
	Push(interface{})

	Pop() interface{}
}
