package main

import "fmt"

func main() {

	a, b, c, d := calc1(1, 2)
	fmt.Println(a, b, c, d)

	sum, sub, mul, div := calc2(2, 1) //3 1 2 2
	fmt.Println(sum, sub, mul, div)
}

//函数的返回值有多可，可以用多个变量去接收到
func calc1(a, b int) (int, int, int, int) {
	return 1, 2, 3, 4
}

func calc2(a, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b
	return
}
