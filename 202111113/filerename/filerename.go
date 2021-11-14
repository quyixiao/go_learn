package main

import "os"

func main() {
	os.Rename("user.log","user.v2.log")// 重命名
	os.Remove("user.v2.log")						//删除文件

}
