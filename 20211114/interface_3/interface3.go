package main

import "fmt"

type A interface {
	do()
}

type B interface {
	do()
}

type C struct {
}

func (c C) do() {
	fmt.Println("测试")
}

func main() {
	var a A = C{}
	a.do()
}
