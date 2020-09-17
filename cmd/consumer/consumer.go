package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var (
	addr    = flag.String("addr", "nsqlookupd:4161", "NSQ lookupd addr")
	topic   = flag.String("topic", "", "NSQ topic")
	channel = flag.String("channel", "", "NSQ channel")
)

// MyHandler handles NSQ messages from the channel being subscribed to
type MyHandler struct {
}

func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	log.Printf("Got a message: %s", string(message.Body))
	return nil
}

func main() {

	// parse the cli options
	flag.Parse()
	if *topic == "" || *channel == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// configure a new Consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(*topic, *channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// register our message handler with the consumer
	consumer.AddHandler(&MyHandler{})

	// connect to NSQ and start receiving messages
	//err = consumer.ConnectToNSQD("nsqd:4150")
	err = consumer.ConnectToNSQLookupd(*addr)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// disconnect
	consumer.Stop()
}
