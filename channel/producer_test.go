package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var ch = make(chan int, 3)
var wg = sync.WaitGroup{}

func JJproducer() {
	i := 0
	for {
		time.Sleep(time.Second)
		ok := JJtrySend(i)
		if !ok {
			continue
		}
		fmt.Printf("%d is send\n", i)
		i++
	}
}

func JJtrySend(i int) bool {
	select {
	case ch <- i:
		return true
	default:
		return false
	}
}

func JJconsumer() {
	for {
		i, ok := JJtryRecv()
		if !ok {
			continue
		}
		fmt.Printf("%d is received\n", i)
	}
}

func JJtryRecv() (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

func TestNew(t *testing.T) {
	wg.Add(2)
	go func() {
		JJproducer()
		wg.Done()
	}()
	go func() {
		JJconsumer()
		wg.Done()
	}()
	wg.Wait()
}
