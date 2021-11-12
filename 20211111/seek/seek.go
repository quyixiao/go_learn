package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("/Users/quyixiao/Desktop/test/test.txt")
	bytes := make([]byte, 100)
	n, _ := file.Read(bytes)
	fmt.Println(n) //3
	fmt.Println(string(bytes))			//123
	//参数有两个，一个是偏移量，一个是相对位置
	// 文件开始 ，0 os.Seeek_SET
	//当前位置 1 ，os.SEEK
	// 文件未尾 2 ，os.SEEK_END
	x, _ := file.Seek(0, 0)

	n, _ = file.Read(bytes)
	fmt.Println(x) //当前所在的位置
	fmt.Println(n) //3
	fmt.Println(string(bytes))		//123
	file.Close()
}
