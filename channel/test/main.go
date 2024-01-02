package main

import (
	"fmt"
	"time"
)

func main() {
	waitc := make(chan struct{})
	go func() {
		fmt.Println("111")
		time.Sleep(2 * time.Second)
		close(waitc)
	}()
	<-waitc
}
