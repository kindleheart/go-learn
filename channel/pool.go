package channel

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

type Pool struct {
	capacity int
	quit     chan struct{}
	active   chan struct{}
	tasks    chan Task
	wg       sync.WaitGroup
}

const (
	defaultCapacity = 10
	maxCapacity     = 1000
)

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}
	p := &Pool{
		capacity: capacity,
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
		tasks:    make(chan Task),
		wg:       sync.WaitGroup{},
	}
	fmt.Println("worker pool start!!!")
	go p.run()
	return p
}

// 向tasks中获取任务执行
func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(i, err)
				<-p.active
			}
			p.wg.Done()
		}()
		fmt.Printf("worker %d start\n", i)
		select {
		case <-p.quit:
			fmt.Printf("worker %d end\n", i)
			<-p.active
		case t := <-p.tasks:
			t()
		}
	}()
}

// 创建capacity个worker
func (p *Pool) run() {
	idx := 0
	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			idx++
			p.newWorker(idx)
		}
	}
}

// Schedule 把t推送到tasks中
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return errors.New("pool is free")
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Println("worker pool freed")
}
