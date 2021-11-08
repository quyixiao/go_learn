package main

import (
	"flag"
	"fmt"
)

func main() {

	port := flag.Int("P", 22, "ssh port")
	host := flag.String("H", "127.0.0.1", "请输入ip")
	help := flag.Bool("h", false, "help")

	fmt.Printf("%T %T %T \n", *port, *host, *help)						// int string bool
	flag.Usage = func() {
		fmt.Println("usage flag arsg [-H 127.0.0.14 ] [-P 22][-v]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()
	fmt.Println(help)
	if *help {
		fmt.Println("-------------------")
		flag.Usage()
	} else {
		fmt.Println(host, port)
		fmt.Println(flag.Args())
	}
}
