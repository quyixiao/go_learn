package main

import (
	"fmt"
	"runtime"
	"sync"
)

//定义counter由incr和desc两个函数分别进行递增和递减
func incr_mutex(lock *sync.Mutex, group *sync.WaitGroup) {
	lock.Lock() //加锁
	defer func() {
		lock.Unlock()
		group.Done()
	}()

	for i := 0; i < 100; i++ {
		counter++
		runtime.Gosched()
	}
}

func desc_mutex(lock *sync.Mutex, group *sync.WaitGroup) {
	lock.Lock() //加锁
	defer func() {
		lock.Unlock() //释放锁
		group.Done()
	}()

	for i := 0; i < 100; i++ {
		counter--
		runtime.Gosched()
	}
}
var counter int = 0

func main() {

	waitGroup := &sync.WaitGroup{}
	//定义锁
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go incr_mutex(lock, waitGroup)
		go desc_mutex(lock, waitGroup)
	}
	waitGroup.Wait()
	fmt.Println(counter) // 不是0
}
