package main

import "fmt"

func main() {

	//索引 => 记录已经加到的n
	sum := 0
	// 初始化子语句 ，条件子语句 ，后置子语句
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0

	i := 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)

	for { //表示一个死循环
		i++
		if i > 3000 {

			break
		}
	}
	desc := "我爱中国"

	// 字符串，数组，切片，映射，管道 ,如果遍历一个字符串，只能用for range去遍历
	for i, ch := range desc {
		fmt.Printf("%T %q\n", i, ch)
	}

	for i := 0; i < 10; i++ {
		if i == 5{
			continue
		}
	}




}
