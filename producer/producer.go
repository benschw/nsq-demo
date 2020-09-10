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
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish(*topic, []byte(*message))
	if err != nil {
		log.Fatal("Could not connect")
	}

	producer.Stop()
}
