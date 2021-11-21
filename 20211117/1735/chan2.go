package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan int = make(chan int, 5)
	var reChannel <-chan int = channel //只读
	var wcChannel chan<- int = channel // 只写

	//只读管道
	go func() {
		wcChannel <- 1
	}()
	// 只写管道
	go func() {
		fmt.Println(<-reChannel)
	}()


	wcChannel <- 3
	fmt.Println(<-reChannel)
	time.Sleep(time.Second * 3 )


}
