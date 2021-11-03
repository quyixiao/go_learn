package main

import "fmt"

func main() {
	sayHelloWorld()
	fmt.Printf("%T \n", sayHelloWorld) // func() ,是一个函数

	sayHi("张三") // 你好 张三 ,实参

	fmt.Println(add(1, 5)) // 6

}

// 定义一个函数
// 参数
// 返回值
func bublo_sort(heights []int) []int {
	//先把最高的人排到最后
	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
			}
		}
	}
	return heights
}

//定义函数 ，定义一个无参，无返回值的函数
func sayHelloWorld() {
	fmt.Println("hello world!")
}

func sayHi(name string) { //形参
	fmt.Println("你好 " + name)
}

func add(a int, b int) int {
	return a + b
}
