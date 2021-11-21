package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//延迟关闭连接
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	input := bufio.NewScanner(os.Stdin)
	for {

		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			fmt.Println(err)
			break
		}

		input.Scan()
		n, err := writer.WriteString(input.Text() + "\n")
		writer.Flush()
		if err != nil {
			fmt.Println(n, err)
			break
		}

	}

	if err != nil {
		os.Exit(-1)
	}

}
