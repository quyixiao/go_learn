package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 一个操作3s未执行完成就任务超时
	result := make(chan int)
	timeout := make(chan int)
	// 任务例程
	go func() {
		interval := time.Duration(rand.Int31n(10)) * time.Second
		fmt.Println(interval)
		time.Sleep( interval)
		result <- 0
	}()
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- 0
	}()

	select {
	case <-result:
		fmt.Println("执行完成 ")
	case <-timeout:
		fmt.Println("执行超时")
	}

}
