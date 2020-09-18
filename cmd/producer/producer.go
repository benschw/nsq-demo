package main

import (
	"flag"
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

var (
	addr    = flag.String("addr", "localhost:4150", "NSQ lookupd addr")
	topic   = flag.String("topic", "", "NSQ topic")
	message = flag.String("message", "", "Message body")
)

func main() {

	// parse the cli options
	flag.Parse()
	if *topic == "" || *message == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// configure a new Producer
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(*addr, config)
	if err != nil {
		log.Fatal(err)
	}

	// publish a nessage to the producer
	err = producer.Publish(*topic, []byte(*message))
	if err != nil {
		log.Fatal(err)
	}

	// disconnect
	producer.Stop()
}
