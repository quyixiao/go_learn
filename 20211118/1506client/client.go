package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
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

	//向服务器发送消息
	fmt.Fprintf(conn, "UnixTime:%d\n", time.Now().Unix())
	//读取服务端发送的消息
	reader := bufio.NewReader(conn)
	cxt, _, _ := reader.ReadLine()
	fmt.Println(string(cxt))
}
