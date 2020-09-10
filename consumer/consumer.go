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
	topic   = flag.String("topic", "", "NSQ topic")
	channel = flag.String("channel", "", "NSQ channel")
)

type MyHandler struct {
}

func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	log.Printf("Got a message: %s", string(message.Body))
	return nil
}

func main() {

	flag.Parse()
	if *topic == "" || *channel == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer(*topic, *channel, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&MyHandler{})

	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatal("Could not connect")
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	consumer.Stop()
}
