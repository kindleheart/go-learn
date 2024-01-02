package main

import (
	"fmt"
	"goLearn/RabbitMQ"
)

func main() {
	r := RabbitMQ.NewRabbitMQSimple("hello")
	r.PublishSimple("hello world!!!")
	fmt.Println("send success!!!")
}
