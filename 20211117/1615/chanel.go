package main

import "fmt"

func main() {
	channel := make(chan string ,2 )

	channel <- "1"

	channel <- "2"

	fmt.Println(len(channel))// 查看管道中的元素 2
	fmt.Println(<- channel)
	fmt.Println(len(channel))// 查看管道中的元素  1
	fmt.Println(<- channel)
	fmt.Println(len(channel))// 查看管道中的元素 0


	channel <- "z"
	channel <- "a"

	close(channel)

	// 需要有某个例程能够关闭管道，否则会发生死锁
	for ch := range  channel{
		fmt.Println(ch)
	}

	//管道关闭是可以再读取管理中的数据，是不允许
	//channel <- "2"

	//goroutine 1 [running]:
	//main.main()
	//        /Users/quyixiao/go/src/go_learn/20211117/1615/chanel.go:20 +0x2a7




}
