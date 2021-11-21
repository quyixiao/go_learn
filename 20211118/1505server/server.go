package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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
	for {
		conn, err := server.Accept()
		if err == nil {
			//使用例程处理与客户端连接
			go func(conn net.Conn) {
				defer conn.Close() //延迟关闭客户端连接
				fmt.Printf("client is connected :%s\n", conn.RemoteAddr())
				//读取客户端发送过来的消息
				reader := bufio.NewReader(conn)
				cxt, _, _ := reader.ReadLine()
				fmt.Println(string(cxt))
				//向客户端发送消息
				fmt.Fprintf(conn, "Time :%s \n", time.Now().Format("2006-01-02 15:04:05 "))

			}(conn)
		}
	}

}
