package main

import "fmt"

func main() {
	/**
	1 * 1 = 1
	1 * 2 = 2 2 * 2 = 4
	1* 3 = 3 2 * 3 = 6   3 * 3 = 9
	.......
	1*9 = 9  ............9 * 9 = 81
	*/
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			// - 号表示左对齐   ， 2 表示占位2 个字符
			fmt.Printf("%d * %d = %-2d\t",j,i, i*j)
		}
		fmt.Println("")

	}
}
