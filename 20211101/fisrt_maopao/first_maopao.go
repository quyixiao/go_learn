package main

import "fmt"

func main() {
	numbs := []int{1, 3, 5, 8, 9, 11, 20, 30, 4, 6, 5}
	maxNum := numbs[0]
	secondNum := numbs[0]
	//求数组中的最大值到第二大值
	for _, v := range numbs {
		if v > maxNum {
			secondNum = maxNum
			maxNum = v
		} else if v > secondNum {
			secondNum = v
		}
	}
	fmt.Println(maxNum) //30 得到最大值
	fmt.Println(secondNum)

}
