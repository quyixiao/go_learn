package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"go_learn/20211104/codes/src/goka"
	_ "go_learn/20211104/codes/src/goka"
	"go_learn/20211104/codes/src/gopkg"
)

func main() {
	fmt.Println("koiwioeiow")
	fmt.Println(gopkg.VERSION)
	goka.PrintName()

	fmt.Println("请输入密码:")
	if bytes, err := gopass.GetPasswd(); err == nil {
		fmt.Println(string(bytes))
	}
}
