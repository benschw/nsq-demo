

NSQ demo repo to try out messaging with NSQ & Go

## Run NSQ with Docker Compose

Start NSQ (the patterns below will integrate with this cluster)

	docker-compose -f docker-compose-nsq.yml up

Browse to the admin UI: http://localhost:4171/


## Examples
    
build a docker container with our consumer app in it
    
    docker build -t nsq-consumer -f Dockerfile-consumer .

### Worker

Start the docker compose example with 2 subscribers to the foo channel

    docker-compose -f docker-compose-worker.yml up

Publish a few messages:

	go run cmd/producer/producer.go -topic email -message "hello world"
	go run cmd/producer/producer.go -topic email -message "hello world"
	go run cmd/producer/producer.go -topic email -message "hello world"
	go run cmd/producer/producer.go -topic email -message "hello world"

Each message gets processed by one or the other consumer, but is only received once.

### Pub Sub

Start the docker compose example with 2 subscribers to the foo channel, and 2 to the bar channel

    docker-compose -f docker-compose-pub-sub.yml up
    
Publish a few messages:

	go run cmd/producer/producer.go -topic orders_placed -message "order 123"
	go run cmd/producer/producer.go -topic orders_placed -message "order 124"
	go run cmd/producer/producer.go -topic orders_placed -message "order 125"
	go run cmd/producer/producer.go -topic orders_placed -message "order 126"


Each channel gets its own copy of each message, but is only delivered to one consumer for each channel
