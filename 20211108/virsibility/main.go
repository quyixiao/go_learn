package main

import (
	"fmt"
	"virsibility/users"
)

func main() {
	v1 := users.User{ID: 11,
		Name: "iodsio",
		//小写的字段在包外不能访问
		//属性小写在包外也是不能访问的
	}
	fmt.Printf("%#v\n",v1)		//users.User{ID:11, Name:"iodsio", birthday:"", addr:users.address{Region:"", streect:"", no:""}}


}
