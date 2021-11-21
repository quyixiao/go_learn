package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 监听 9999商品

func main() {
	addr := "0.0.0.0:9999"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("lister :", addr)
	defer server.Close()

	input := bufio.NewScanner(os.Stdin)

	for {
		conn, err := server.Accept()
		if err == nil {
			reader := bufio.NewReader(conn)
			writer := bufio.NewWriter(conn)
			fmt.Println("客户端%s连接成功：", conn.RemoteAddr())
			for {

				input.Scan()

				_, err := writer.WriteString(input.Text() + "\n")
				writer.Flush()
				if err != nil {
					fmt.Println(err)
					break
				}
				input, err := reader.ReadString('\n')
				fmt.Println(input)
				if err != nil {
					fmt.Println(err)
					break
				}

			}
			conn.Close()
		}
	}

}
