package queue_listener

import (
	"fmt"
	"nats-example/queue"
	"time"
)

type queueListener struct {
	queue              queue.Queue
	number_of_listener int
	state              []chan struct{}
}

func (q *queueListener) ListenQueue() {
	for i := 0; i < q.number_of_listener; i++ {
		q.state = append(q.state, q.startListen(i))
	}
}

func (q *queueListener) startListen(x int) chan struct{} {
	listener := make(chan interface{})
	state := make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				//TODO: Process your job here
				msg := q.queue.Pop()
				if msg == nil {
					continue
				}

				fmt.Println("[", x, "] ", "processing message", "->"+string(msg.([]byte))+"<-", "at", time.Now())
				time.Sleep(2 * time.Second)
			case <-state:
				close(listener)
				close(state)
				break
			}
		}
	}()

	return state

}

func (q *queueListener) StopListen() {
	for i := 0; i < len(q.state); i++ {
		q.state[i] <- struct{}{}
	}
}
