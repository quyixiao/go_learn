package main

import (
	"fmt"
)

func main() {

	const  Name  = "kk"

	// 省略类型
	const PI = 3.1
	// 定义多个常量（类型相同 ）
	const  C1,C2 = 1,2
	// 定义多个常量 （类型不同）
	const
	(
		C3 string = "xxx"
		C4 int = 1
	)

	const c5 ,c6 = "xx",1

	// 定义多个常量 省略类型
	fmt.Println(Name)

	const a,b = "cc","dd"

	const (
		c = "xx"
		d = "f"
	)
	fmt.Println(a,b,c,d)
	fmt.Println(C1,C2,C3,C4)
	// 常量定义的是一些不可变的值 ，  要大写编写
	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12
		C13 string = "kk"
		//如果赋值和上面的类型一样的话，就可以省略掉
	)
	fmt.Println(C7,C8,C9,C10,C11,C12,C13)
	// 1 1 1 3.14 3.14 3.14 kk
	// 枚举 ，const + iota
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1,E2,E3)
	// 0 1 2
	const (
		E4 int = iota
		E5
		E6
	)
	fmt.Println(E4,E5,E6)
	// 0 1 2


	const (
		E7 int = (iota+1)*100
		E8
		E9
	)
	fmt.Println(E7,E8,E9)
	// 100 200 300

}

