package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//文件内容是123
	path := "/Users/quyixiao/Desktop/test/test.txt"
	file, err := os.Open(path)
	fmt.Println(file)
	fmt.Println(err)
	fmt.Printf("%T \n", file) //*os.File 操作系统的文件类型

	defer file.Close()					//如果读取文件报错，延迟关闭文件
	bytes, err := ioutil.ReadAll(file)
	fmt.Println(string(bytes)) //123
	bytex, e := ioutil.ReadFile("/Users/quyixiao/Desktop/test/test.txt")
	fmt.Println(string(bytex), e) //123 <nil>

	file.Close()
}
