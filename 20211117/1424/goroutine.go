package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	waitGroup := &sync.WaitGroup{}
	var counter int
	incr := func() {
		defer waitGroup.Done()
		for i := 0; i < 100; i++ {
			counter++
			runtime.Gosched()
		}
	}
	desc := func() {
		defer waitGroup.Done()
		for i := 0; i < 100; i++ {
			counter--
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
