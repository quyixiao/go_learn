package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// 每次读取就一行
	reader := strings.NewReader("12345\n6789")

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
