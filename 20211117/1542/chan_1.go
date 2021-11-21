package main

import (
	"fmt"
	"sync"
)

func main() {
	//初始化管道
	var channel chan string
	fmt.Printf("%T %v \n", channel, channel) //chan string <nil>
	channel = make(chan string)

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			//写的时候，如果没有读取，将会阻塞住
			channel <- fmt.Sprintf("%d", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
