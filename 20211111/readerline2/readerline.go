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
			line, err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
			} else {
				fmt.Println(string(line))

			}

		}
	}

}
