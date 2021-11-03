package main

import (
	"fmt"
)

func main() {
	add2 := func(n int) int {
		return n + 2
	}
	fmt.Println(add2(2))

	add10 := func(n int) int {
		return n + 10
	}
	fmt.Println(add10(10))

	// 闭包问题
	addbase := func(base int) func(int) int {
		// 返回一个函数
		return func(n int) int {
			return base + n
		}
	}
	fmt.Println(addbase(8)(10)) // 18
	fmt.Printf("%T \n",addbase(8))			// 返回一个函数
	add4 := addbase(4)
	add8 := addbase(8)
	add11 := addbase(11)

	fmt.Println(add4(4))					//8
	fmt.Println(add8(4))						// 12
	fmt.Println(add11(4))//15
}

