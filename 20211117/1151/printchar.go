package main

import (
	"fmt"
	"runtime"
	"time"
)

func printChars(prefix string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()

//		time.Sleep(time.Second)
	}
	// 让出CPU
	// time.Sleep(1 * time.Microsecond)
}

func main() {
	// 把字符串转换为数字
	go printChars("********")
	go printChars("========")


	// 主例程执行完之后，子
	printChars("--------")
	//
	time.Sleep(time.Second * 3)
}
