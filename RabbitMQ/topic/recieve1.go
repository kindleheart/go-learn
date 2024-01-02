package main

import "goLearn/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQTopic("topic", "#")
	imoocOne.RecieveTopic()
}
