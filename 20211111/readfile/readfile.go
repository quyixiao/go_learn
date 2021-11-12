package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "/Users/quyixiao/Desktop/test/test.txt"
	file, err := os.Open(path)
	fmt.Println(file)
	fmt.Println(err)
	fmt.Printf("%T \n", file) //*os.File 操作系统的文件类型
	var bytes = make([]byte, 20)
	for {
		n, err := file.Read(bytes)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(bytes[:n]))
	}
	file.Close()

}
