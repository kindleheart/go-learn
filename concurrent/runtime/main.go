package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// gosched让出时间片
func gosched() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
		wg.Done()
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
	wg.Wait()
}

// goexit 终止当前的Goroutine，在终止之前会执行延迟调用函数。
func goexit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}

func a() {
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go b()
	go a()
	time.Sleep(time.Second)
}
