package main

import (
	"flag"
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

var (
	topic   = flag.String("topic", "", "NSQ topic")
	message = flag.String("message", "", "Message body")
)

func main() {

	flag.Parse()
	if *topic == "" || *message == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := nsq.NewConfig()
	producer, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := producer.Publish("my_topic", []byte(*message))
	if err != nil {
		log.Panic("Could not connect")
	}

	producer.Stop()
}
