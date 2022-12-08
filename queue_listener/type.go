package queue_listener

import "nats-example/queue"

func NewQueueListener(
	queue queue.Queue,
	number_of_listener int,
) QueueListener {
	return &queueListener{
		queue:              queue,
		number_of_listener: number_of_listener,
	}
}

type QueueListener interface {
	ListenQueue()

	StopListen()
}
