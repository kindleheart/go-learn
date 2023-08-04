package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex()
}

func Map() {
	m := sync.Map{}
	m.Store("cat", "mimi")
	m.Store("name", "jay")

	if v, ok := m.Load("cat"); ok {
		fmt.Println(v.(string))
	}
}

func mutex() {
	mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println("mutex")
	defer mutex.Unlock()
}

func Failed1() {

}
