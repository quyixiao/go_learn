package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("user.log")
	if err == nil {
		defer file.Close()
		reader := bufio.NewReader(file)
		bytes := make([]byte, 5)
		reader.Read(bytes)
		fmt.Println(string(bytes)) //99382 读取5个长度的内容

		for {
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
			} else {
				fmt.Println(isPrefix, string(line))
				//false 322222222222
				//false 9832983
				//false 89i32832
				//false dsiodia
				//false iodsoia
			}

		}
	}

}
