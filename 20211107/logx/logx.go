package main

import "log"

func main() {
	log.Printf("我是一个人")
	log.Printf("我是一个人")
	log.SetFlags(log.Flags() | log.Lshortfile)			//prefix2021/11/07 20:30:46 logx.go:15: 我是fatalf日志z
	//log.SetFlags(log.Flags() | log.Llongfile)
	log.SetPrefix("prefix") //在日志前面加上前缀 prefix2021/11/07 20:29:59 /Users/quyixiao/go/src/go_learn/20211107/logx/logx.go:14: 我是fatalf日志z


	//打印一个日志，并且触发一个panic
	//log.Panicf("我是panic日志 ：%s","y")
	log.Fatalf("我是fatalf日志%s", "z")
	log.Fatalf("我是fatalf日志%s", "z") //打印之后直接退出


}
