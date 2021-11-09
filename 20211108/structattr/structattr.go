package main

import "fmt"

type MyUser struct {
	ID       int
	userName string
	age      int
	Addr     string
	remark   string
	tel      string
}

func main() {

	me := MyUser{
		userName: "kk",
		ID:       1,
		Addr:     "西安",
	}
	fmt.Println(me)

	me.tel = "182898329839"
	fmt.Printf("%#v\n", me)
	me2 := &MyUser{
		ID:       2,
		userName: "wuha",
	}

	fmt.Printf("%T \n", me2)
	(*me2).tel = "18389389238"

	fmt.Printf("%#v\n", me2)

	me2.Addr = "北京 "
	fmt.Printf("%#v\n", me2)
}
