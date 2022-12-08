package main

import (
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	err = createStream(js)
	if err != nil {
		log.Fatal(err)
	}

	// queue := queue.NewQueue()
	// queueListener := queue_listener.NewQueueListener(queue, 100)
	// queueListener.ListenQueue()
	// defer queueListener.StopListen()

	// // // // // // Use a WaitGroup to wait for 10 messages to arrive
	// // // // // // wg := sync.WaitGroup{}

	// if _, err := js.QueueSubscribe("CHATS.test", "queue", func(m *nats.Msg) {
	// 	queue.Push(m.Data)
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// }

	for i := 1; i <= 800; i++ {
		_, err = js.Publish("CHATS.test", []byte(strconv.Itoa(i)))
	}

}

const (
	streamName     = "CHATS"
	streamSubjects = "CHATS.*"
)

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
