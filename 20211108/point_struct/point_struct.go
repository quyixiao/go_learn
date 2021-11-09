package main

import "fmt"

type Address struct {
	Region  string
	Streect string
	No      string
}

type User struct {
	Id   int
	Name string
	Addr *Address
}

type Employee1 struct {
	*User
	Salary float64
	Name   string
}

func main() {
	var me01 User
	fmt.Printf("%#v \n", me01)

	me02 := User{
		Id:   1,
		Name: "kk",
		Addr: &Address{
			"西安市",
			"大山",
			"大模大样",
		},
	}
	fmt.Printf("%#v\n", me02)      //main.User{Id:1, Name:"kk", Addr:(*main.Address)(0xc00006e180)}
	fmt.Println(me02.Addr.Streect) //大山
	me02.Addr.Region = "北京市"
	fmt.Printf("%#v\n", me02.Addr) //&main.Address{Region:"北京市", Streect:"大山", No:"大模大样"}

	me03 := Employee1{
		User: &User{
			Id:   1,
			Name: "kk",
			Addr: &Address{
				"西安",
				"业路",
				"iodis",
			},
		},
	}
	fmt.Println(me03.Addr.Streect) //业路

	user1 := User{}
	user2 := user1
	user2.Name = "张三"
	fmt.Printf("%#v\n", user1) //main.User{Id:0, Name:"", Addr:(*main.Address)(nil)}
	fmt.Printf("%#v\n", user2) //main.User{Id:0, Name:"张三", Addr:(*main.Address)(nil)} 按值类型传递
	change(user1)
	fmt.Printf("%#v\n", user1)		//main.User{Id:0, Name:"", Addr:(*main.Address)(nil)} 按值类型传递
	changePoint(&user1)
	fmt.Printf("%#v\n", user1)		//main.User{Id:0, Name:"yyyyyyy", Addr:(*main.Address)(nil)} 按地址传递

	mex := User{
		Id: 1 ,
		Name: "kk ",
		Addr: &Address{"西安","锦业","001"},
	}
	mex2 := mex
	fmt.Println(mex2.Name)		//
	// fmt.Println(mex2.Addr.Streect ) //如果没有初始化 的话 panic: runtime error: invalid memory address or nil pointer dereference



}

func change(u User) {
	u.Name = "xxxxxx"
}

func changePoint(u * User ){
	u.Name = "yyyyyyy"
}
