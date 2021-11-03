package main

import "fmt"







func main() {
	fmt.Printf("%T\n", fsum) // func(int, int) int
	var f func(int, int) int = fsum
	fmt.Println(f(1, 2))
	printx(list, "1", "2", "3") // 1       2       3


	// 定义一个匿名函数
	sayHello1 := func (name string) {
		fmt.Println("hello ", name)
	}
	sayHello1("kka")

	//匿名函数只调用一次的情况下使用
	func (name string){					// hahah
		fmt.Println(name)
	}("hahah")

	values := func(args ...string) {
		for _,v := range  args {
			fmt.Println(v)
		}
	}

	printx(values,"A","B","C")
}


func fsum(a, b int) int {
	return a + b
}

// callback格式 ，将传递的数据按照每行打印 还是按照一行一行打印
func printx(callback func(...string), args ...string) {
	fmt.Println("print函数的输出 ：")
	callback(args...)



}


func list(args ...string) {
	for _, v := range args {
		fmt.Printf("%s \t", v)
	}
}
