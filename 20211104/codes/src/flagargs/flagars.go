package main

import (
	"flag"
	"fmt"
)

func main() {

	var port int
	// 绑定命令行参数与变量关系
	var host string
	var help bool
	flag.IntVar(&port, "P", 22, "ssh port")
	flag.StringVar(&host, "H", "127.0.0.1", "请输入ip")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage flag arsg [-H 127.0.0.14 ] [-P 22][-v]")
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()
	fmt.Println(help)
	if help {
		fmt.Println("-------------------")
		flag.Usage()
	} else {
		fmt.Println(host, port)
		fmt.Println(flag.Args())
	}
}
