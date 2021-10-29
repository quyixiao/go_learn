package main

import "fmt"

func main() {
	// 作用域，定义标识符可以使用的范围
	// 在go中用{} 来定义 作用域的范围
	// 使用原则，子语句块可以使用父语句块，但是父语句块中不能使用子语句块变量
	// 变量在语句块中必需使用
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		{
			inner2 := 3
			//
			inner = 3
			fmt.Println(outer,inner,inner2)
		}
	}

	//fmt.Println(inner)



	fmt.Println("%T,%s,%d","kk","kk",10)
}
