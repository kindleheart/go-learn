package main

import (
	"fmt"
	"goLearn/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	topicOne := RabbitMQ.NewRabbitMQTopic("topic", "jason.topic.one")
	topicTwo := RabbitMQ.NewRabbitMQTopic("topic", "jason.topic.two")
	for i := 0; i <= 10; i++ {
		topicOne.PublishTopic("Hello topic one!" + strconv.Itoa(i))
		topicTwo.PublishTopic("Hello topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
