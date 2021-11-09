package main

import (
	"fmt"
)

type Address struct {
	Region  string
	Streect string
	No      string
}

type User struct {
	Id   int
	Name string
	Addr Address
}

func NewUser(id int, name string, region string) User {
	return User{
		Id:   id,
		Name: name,
		Addr: Address{
			Region: region,
		},
	}
}

func main() {
	me := User{}
	me2 := me
	me2.Name = "kk"
	fmt.Printf("%#v\n", me) //值类型
	fmt.Printf("%#v\n", me2)
	user := NewUser(1,"iodsoi","83ew98")
	fmt.Printf("%#v\n",user)		//main.User{Id:1, Name:"iodsoi", Addr:main.Address{Region:"83ew98", Streect:"", No:""}}

}
