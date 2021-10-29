package main

import (
	"fmt"
)

func main() {
	// 布尔类型，表示真，假，标识符 bool , 字面量 ： 可以选择true ,false
	// 零值为false
	var zero bool
	isBoy := true
	isGirl := true
	fmt.Println(zero, isGirl, isBoy)
	// 操作
	// 逻辑运算(与，&& ,或 || ,非 ! )
	// aBool ,bBool
	// aBool ,当两个都为true的时候，结果都为True
	//    aBool 的可选 值为true ,false
	//    bBool 的可选 值  true   false
	// true &&true  true
	// true && false false
	// true || false true
	// false || false false

	// 关系运算，（== ,!= ）

	// && 的运算符
	fmt.Println(true && true)
	fmt.Println(true && false)
	// || 只要有一个为true ,结果为true
	fmt.Println(false || true)
	//

	fmt.Println(!true)
	fmt.Println(!false)

	// 等于判断
	fmt.Println(true == false )
	fmt.Println(true == true )

	fmt.Println( zero == true )

	//打印变量的类型，不能用Println ,%t 就是占位bool类型的
	fmt.Printf("%T %t\n",zero,zero)

}
