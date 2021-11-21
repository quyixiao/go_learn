package main

import "fmt"

func main() {
	// 定义两个带缓冲区的管道
	channel01 := make(chan int, 1)
	channel02 := make(chan int, 1)
	// 启用例程像两个管道中写入数据
	go func() {
		channel01 <- 1
	}()

	go func() {
		channel02 <- 2
	}()

	// 使用select从两个管道中读取数据
	select {
	case <-channel01:
		fmt.Println("channel01 ")
	case <-channel02:
		fmt.Println("channel02")
	}

}
