package main

import (
	"fmt"
	"goLearn/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rOne := RabbitMQ.NewRabbitMQRouting("exName", "routing_one")
	rTwo := RabbitMQ.NewRabbitMQRouting("exName", "routing_two")
	for i := 0; i <= 10; i++ {
		rOne.PublishRouting("Hello routing one!" + strconv.Itoa(i))
		rTwo.PublishRouting("Hello routing Two!" + strconv.Itoa(i))
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

}
