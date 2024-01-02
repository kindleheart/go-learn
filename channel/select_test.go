package channel

import (
	"fmt"
	"testing"
	"time"
)

func work(ch <-chan int) {
	after := time.After(time.Second * 8)
	heartbeat := time.NewTicker(3 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case v := <-ch:
			// do something
			fmt.Println(v)
		case <-heartbeat.C:
			fmt.Println("心跳机制")
		case <-after:
			fmt.Println("超时机制")
			return
		}
	}
}

func spawn(ch <-chan int) chan struct{} {
	c := make(chan struct{})
	go func() {
		work(ch)
		c <- struct{}{}
	}()
	return c
}

func TestSelectTime(t *testing.T) {
	ch := make(chan int, 1)
	c := spawn(ch)
	time.Sleep(time.Second * 5)
	ch <- 1
	<-c
}
