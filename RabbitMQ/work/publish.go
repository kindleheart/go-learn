package main

import (
	"fmt"
	"goLearn/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"hello")

	for i := 0; i <= 20; i++ {
		rabbitmq.PublishSimple("Hello work!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
