package main

import "goLearn/RabbitMQ"

func main() {
	r := RabbitMQ.NewRabbitMQSimple("hello")
	r.ConsumeSimple()

}
