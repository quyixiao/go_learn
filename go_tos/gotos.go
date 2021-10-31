package main

import "fmt"

func main() {

	/*	var yes string
			fmt.Println("请输入你的操作：（y 或 Y）")

			fmt.Scan(&yes)

			if yes != "Y" || yes != "y" {
				goto end
			}

			fmt.Println("买一个西瓜")

		end:
	*/

	// 1... 100
	i := 0
	sum := 0
start:
	if i > 100 {
		fmt.Println("执行结束 ")
		goto forward
	}
	sum += i
	i++
	fmt.Println("i=", i)
	goto start
forward:
	fmt.Println("sum is ", sum)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				//break只能跳出本层循环
				break
			}
			fmt.Println(i, j)
		}
	}

	//如果跳出多层循环，可以使用beak  begain
begain:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				//break只能跳出本层循环
				break begain
			}
			fmt.Println(i, j)
		}
	}
}
