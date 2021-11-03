package main

import "fmt"

func main() {
	fmt.Println(sumx(100))
	//
	fmt.Println(mulx(9)) // 324

	tower("A","B","C",3)
}

// 递归 100 + ... + 1
func sumx(index int) int {
	if index <= 1 {
		return 1
	}
	return index + sumx(index-1)
}

// n !
func mulx(index int) int {
	if index <= 1 {
		return 1
	}
	return index * sumx(index-1)
}

// 1个盘子 1 -C             1
// 2个盘子 1 -> B ,2 -> C ,1 -> C   3
// 3个盘子 1->C 2->B ,1 -> B , 3 -> C , 1-> A ,2 -> C ,1 -> C   7  ,
// 4个盘子 1 -> B , 2 -> C , 1 -> C , 3 -> B , 1 - > A , 2 -> B ,1 -> B , 4 -> C , 1 -> C , 2 -> A ,1 -> A , 2 -> C , 1 -> B , 2 -> C , 1 -> C  , 15
// 汉诺塔游戏
func tower(a, b, c string, lay int) {
	if lay == 1 {
		fmt.Println(a, "->", c)
		return
	}

	tower(a,c,b,lay-1)
	fmt.Println(a ,"->",c)
	tower(b,a,c,lay -1 )
}
