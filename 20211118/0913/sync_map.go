package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	ID   int
}

func main() {
	var users sync.Map

	users.Store("20", "小明")
	users.Store("30", "张三")
	if value, ok := users.Load("20"); ok == true {
		fmt.Println(value.(string), ok)
	}
	users.Store("userName", User{Name: "张三", ID: 1})
	if v, ok := users.Load("userName"); ok {
		u := v.(User)							//强制类型转化
		fmt.Println(u.ID, u.Name)
	}

}
