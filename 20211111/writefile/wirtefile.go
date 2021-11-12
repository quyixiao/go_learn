package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/Users/quyixiao/Desktop/test/test.txt"
	file, err := os.Create(path)
	fmt.Println(file)
	fmt.Println(err)
	fmt.Printf("%T \n", file) //*os.File 操作系统的文件类型

	file.Write([]byte ("abcdef"))
	file.Close()
}
