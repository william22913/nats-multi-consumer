package queue

import "sync"

type queue struct {
	sync.Mutex

	queue []interface{}
}

func (q *queue) Push(element interface{}) {
	q.Lock()
	defer q.Unlock()

	q.queue = append(q.queue, element)
}

func (q *queue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	element := q.queue[0]
	q.queue = q.queue[1:]

	return element
}
