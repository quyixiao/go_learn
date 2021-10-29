package main

import "fmt"

// 在函数内定义的局部变量一定需要使用，但是在函数外定义的全局变量不一定需要使用
var version string = "1.0"

func main() {
	// 定义变量me为string类型的变量
	//变量名需要满足标识符命名规则
	// 1.必需是非空的unicode字符串组成，数字 ，
	// 2.不能以数字开头
	// 3.不能为go的关键字（25个）
	// 4.避免和go的预定义标识符冲突，true,false
	// 5.驼峰
	// 6. 标识符区别大小写
	// 7. go定义的变量必需是需要使用的，不然会报错
	var me  = "张三"
	fmt.Println(me)
	// 如果一个变量没有赋值的话，默认为零值，每一种数据类型的零值是不一样的
	fmt.Println(version)

	var name,user string = "zhangsan","lissi"
	fmt.Printf(name)
	fmt.Printf(user)
	//定义多个变量，但是类型是不一样的
	var (
		age int = 11
		higxx int = 6
	)
	fmt.Println(age,higxx)
	age ,higxx = higxx ,age
	fmt.Println(age,higxx)
	var (
		s = "kk"
		a = 32
	)
	s,a = "bb",32
	fmt.Println(a,s)

	//这是一个简单的声明，只能在函数内部使用
	isBoy := true
	fmt.Println(isBoy)

}
