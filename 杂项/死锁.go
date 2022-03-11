package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

type value struct {
	mutex sync.Mutex
	id int
}

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	var v1, v2 value
	v1.id = 1
	v2.id = 2
	
	run := func(v1, v2 *value) {
		defer wg.Done()
		v1.mutex.Lock()
		time.Sleep(time.Second * 2)
		v2.mutex.Lock()
		fmt.Printf("id: %d",v1.id)
		v2.mutex.Unlock()
		v1.mutex.Unlock()
	}

	wg.Add(2)
	go run(&v1, &v2)
	go run(&v2, &v1)
	wg.Wait()

}