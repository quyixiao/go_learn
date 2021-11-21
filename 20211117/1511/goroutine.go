package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	waitGroup := &sync.WaitGroup{}
	var counter int32 = 0
	incr := func() {
		defer waitGroup.Done()
		for i := 0; i < 100; i++ {

			atomic.AddInt32(&counter, 1)
			runtime.Gosched()
		}
	}
	desc := func() {
		defer waitGroup.Done()
		for i := 0; i < 100; i++ {
			atomic.AddInt32(&counter, -1)
			runtime.Gosched()
		}
	}
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go incr()
		go desc()
	}
	waitGroup.Wait()
	fmt.Println(counter) // 不是0
}
