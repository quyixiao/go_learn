package main

import (
	"fmt"
	"os"
)

func main() {
	//名称，权限
	//os.Mkdir("aaa",0644)
	//os.Remove("test02")
	//os.Rename("aaa","bbb")					//重命名
	//os.Remove("bbb")
	err := os.Mkdir("test/bb", 0644)
	if err != nil {
		fmt.Println(err)							//mkdir test/bb: no such file or directory 如果目录不存在
	}
	//如果父目录不存在 ，如何创建一个文件夹呢？
	os.MkdirAll("test1/xxx",0644)
	// 递归移除目录
	//os.RemoveAll("test1")

	// go中判断一个文件是否存在


}
