package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	result := make(chan int)
	//启动工作例程
	go func() {
		interval := time.Duration(rand.Int31n(10)) * time.Second
		fmt.Println("sleep:", interval)
		time.Sleep(interval)
		result <- 0
	}()

	//使用select从两个管道中读取数据
	select {
	case <-result:
		fmt.Println("ok")
	case <-time.After(5 * time.Second): //设置5s超时
		fmt.Println("timeout")
	}
}
