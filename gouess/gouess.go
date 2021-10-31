package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置随机数种子
	rand.Seed(time.Now().Unix())

	// 生成随机数 0-100
	guess := rand.Intn(100)
	const maxGuessNum = 5
	fmt.Println("生成的随机数：" ,guess)
start:
	for i := 0; i < maxGuessNum; i++ {
		var input int
		fmt.Println("请输入你猜中的数字")
		fmt.Scan(&input)

		if guess == input {
			fmt.Println("太聪明了 ，你第%d次就猜对了", i+1)
			break
		} else if guess > input {
			fmt.Println("你猜的数字太小了，你还有%d次猜的机会", maxGuessNum-i-1)
		} else {
			fmt.Println("你猜的数字太了，你还有%d次猜的机会", maxGuessNum-i-1)
		}
		if i == maxGuessNum-1 {
			fmt.Println("太笨了，你还要来一次吗？如果需要输入Y,否则输入Other")
			var flag string
			fmt.Scan(&flag)
			if flag == "Y" {
				goto start
			} else {
				goto end
			}
		}
	}
end:
	fmt.Println("游戏结束 ")

}
