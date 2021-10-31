package main

import "fmt"

func main() {
	var yes string
	fmt.Scan(&yes)
	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")

	switch yes {
	case "Y", "X": //可以写多个变量，Java中只允许写一个变量
		fmt.Println("买一个西瓜")
		break
	case "y":
		fmt.Println("买两个西瓜")
		break
	default:
		fmt.Println("什么都不买")
		break

	}

	var score int
	fmt.Println("请输入成绩：")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}
}
