package main

import "goLearn/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("hello")
	rabbitmq.ConsumeSimple()
}
