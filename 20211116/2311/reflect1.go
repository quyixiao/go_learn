package main

import (
	"fmt"
	"reflect"
)

func main() {
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Score",
			Type: reflect.TypeOf(int64(0)),
		},
	}
	//定义结构体类型
	UserType := reflect.StructOf(fields)
	// 创建结构体对象
	user := reflect.New(UserType).Elem()
	//设置属性值
	user.Field(0).Set(reflect.ValueOf("Silence"))
	user.Field(1).Set(reflect.ValueOf(int64(10000)))

	//打印对象和对象指针
	fmt.Printf("%#v\n", user.Interface()) //struct { Name string; Score int64 }{Name:"Silence", Score:10000}

	fmt.Printf("%#v\n", user.Addr().Interface()) //&struct { Name string; Score int64 }{Name:"Silence", Score:10000}


}
