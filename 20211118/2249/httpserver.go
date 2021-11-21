package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	addr := ":9999"
	lister, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		func(){
			defer func(){
				conn.Close()
				fmt.Println("Client Closed:",conn.RemoteAddr())
			}()
		}()
		fmt.Println("Client connected :", conn.RemoteAddr())
	}

	lister.Close()
}
