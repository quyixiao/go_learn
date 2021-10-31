package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入名字 ：")
	fmt.Scan(&name)

	fmt.Println("你输入的名字是：" + name)

	var age int
	fmt.Println("请输入的年龄 ")
	fmt.Println(&age)
	fmt.Println("你输入的年龄是", age)

	var height float64
	fmt.Println("请输入身高：")
	fmt.Scan(&height)
	fmt.Println("身高", height)

}
