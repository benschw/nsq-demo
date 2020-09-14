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

	flag.Parse()
	if *topic == "" || *message == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(*addr, config)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish(*topic, []byte(*message))
	if err != nil {
		log.Fatalf("Could not connect to %s", *addr)
	}

	producer.Stop()
}
