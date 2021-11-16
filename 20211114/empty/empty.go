package main

import (
	"fmt"
	"go_learn/20211111/codes/manager"
)

type EStruct struct {
}

type Empty interface {
}

func fargs(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
		switch v := arg.(type) {
		case int:
			fmt.Printf("int 类型%T %v \n", v, v)
			break
		case string:
			fmt.Printf("string类型%T %v \n", v, v)
			break
		default:
			fmt.Printf("other %T %v \n", v, v)
			break
		}
	}
}

func main() {
	es := EStruct{}
	var e Empty
	fmt.Printf("%T %T \n", es, e)

	var i interface{} = 1
	fmt.Println(i)
	i = manager.User{} //空接口可以接收任意类型的参数

	fargs("oiewio", 1, 2) //可以接收任意类型的参数

}
