package main

import "goLearn/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exName", "")
	imoocOne.RecieveRouting("routing_two")
}
