

## Run NSQ with Docker Compose

Start NSQ

	docker-compose -f docker-compose-nsq.yml up

Browse to the admin UI: http://localhost:4171/


## Examples

### Worker

Run two consumers (open two terminals and run this in each):

	go run consumer/consumer.go -topic my_topic -channel foo


Publish a few messages:

	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"

Each message gets processed by one or the other consumer, but is only received once.

### Pub Sub

Run a first consumer, creating a "foo" channel:

	go run consumer/consumer.go -topic my_topic -channel foo

Run a second consumer, creating a "bar" channel:

	go run consumer/consumer.go -topic my_topic -channel bar

Publish a few messages:

	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"
	go run producer/producer.go -topic my_topic -message "hello world"

Each consumer (subscriber) gets its own copy of each message.
