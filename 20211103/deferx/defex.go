package main

import "fmt"

func main() {

	// 当函数退出的时候执行
	defer func() {
		fmt.Println("defer2")
	}()

	defer func() {
		fmt.Println("defer2")
	}()

	fmt.Println("main over ")
	//main over
	//xxx






}
