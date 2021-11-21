package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("new ")
			return 1
		},
	}

	x := pool.Get()
	fmt.Println(x)
	pool.Put(x)
	x = pool.Get() //第二次没有创建

	x1 := pool.Get() //第二次没有创建
	fmt.Println(x1)

	//new
	//1
	//new
	//1

}
