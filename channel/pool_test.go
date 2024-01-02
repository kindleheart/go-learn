package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestPool_Schedule(t *testing.T) {
	pool := New(8)
	err := pool.Schedule(func() {
		fmt.Println("任务1开始")
		time.Sleep(4 * time.Second)
		fmt.Println("任务1结束")
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pool.Schedule(func() {
		fmt.Println("任务2开始")
		time.Sleep(2 * time.Second)
		fmt.Println("任务2结束")
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	pool.free()
}
