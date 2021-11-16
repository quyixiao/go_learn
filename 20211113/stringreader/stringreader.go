package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 像操作流一样的操作字符串
	reader := strings.NewReader("123456789")
	bytes := make([]byte, 3)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(n, string(bytes[:n]))

	}

}
